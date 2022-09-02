package main

import (
	"database/sql"
	"fmt"
	"log"
	"rest/db"
	"strings"
	"time"

	go_odoo "github.com/skilld-labs/go-odoo"
)

type WellsUpdate struct {
	ID             int64
	WellType       string
	WellNumber     string
	LastUpdateDate string
}

// wells injection with gauges
func FindInjectionWells() []WellsUpdate {
	var wellsNumbers []WellsUpdate
	wells, err := ClientOdoo.FindAssetAssets(go_odoo.NewCriteria().Add("asset_class_choice", "=", "well"), go_odoo.NewOptions().FetchFields("name", "asset_type_equipment", "gauge_new_lines"))
	if err != nil {
		log.Println(err)
	}
	for _, well := range *wells {
		if well.AssetTypeEquipment.Name == "Injection" {
			wellnumber := strings.TrimPrefix(well.Name.Get(), "Скважина №")
			if len(well.GaugeNewLine.Get()) != 0 {
				gauges, err := ClientOdoo.GetAssetGaugeNews(well.GaugeNewLine.Get())
				if err != nil {
					log.Println(err)
				}
				if gauges != nil {
					last := *gauges
					wellsNumbers = append(wellsNumbers, WellsUpdate{
						ID:             well.Id.Get(),
						WellType:       "Injection",
						WellNumber:     wellnumber,
						LastUpdateDate: last[len(last)-1].StartDatetime.Get().Format("2006-01-02 00:00:00"),
					})
				}
			} else {
				wellsNumbers = append(wellsNumbers, WellsUpdate{
					ID:             well.Id.Get(),
					WellType:       "Injection",
					WellNumber:     wellnumber,
					LastUpdateDate: "2012-01-01 00:00:00",
				})
			}
		}
	}
	return wellsNumbers
}

// Показания датчиков на нагнетательных скважинах
func GetWellInjectReadings(wellNumber string, lastDate string, assetID int64) {
	conn := db.DatabaseConnection()
	defer conn.Close()
	rows, err := conn.Query(`SELECT ITEM_NAME
	,START_DATETIME,CASING_PRESS1,CASING_PRESS2,CASING_PRESS3
	,CASING_PRESS4,PUMP_RPM,PUMP_CURRENT,PUMP_FREQUENCY,FLOW_PRESS,PUMP_EFFICIENCY,WH_PRESS,WH_TEMP
FROM VT_WELL_READ_ru_RU
WHERE START_DATETIME > ? AND ITEM_NAME = ?`, lastDate, wellNumber)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	wellTests := []WellTest{}
	for rows.Next() {
		var Name, Date string
		var CassingPressure1, CassingPressure2, CassingPressure3, CassingPressure4, PumpRPM, PumpCurrent, PumpFrequency, FlowPress, PumpEfficiency, WHPerss, WHTemp sql.NullFloat64
		err := rows.Scan(&Name, &Date, &CassingPressure1, &CassingPressure2, &CassingPressure3, &CassingPressure4, &PumpRPM, &PumpCurrent, &PumpFrequency, &FlowPress, &PumpEfficiency, &WHPerss, &WHTemp)
		if err != nil {
			log.Println(err, "ERROR Scan Rows")
		}
		fmt.Println(Name, CassingPressure1, CassingPressure2, CassingPressure3, CassingPressure4, PumpRPM, PumpCurrent, PumpFrequency, FlowPress, PumpEfficiency, WHPerss, WHTemp)
		wellTests = append(wellTests, WellTest{
			ID:               assetID,
			Date:             Date,
			CassingPressure1: CassingPressure1.Float64,
			CassingPressure2: CassingPressure2.Float64,
			CassingPressure3: CassingPressure3.Float64,
			CassingPressure4: CassingPressure4.Float64,
			PumpRPM:          PumpRPM.Float64,
			PumpCurrent:      PumpCurrent.Float64,
			PumpFrequency:    PumpFrequency.Float64,
			FlowPress:        FlowPress.Float64,
			PumpEfficiency:   PumpEfficiency.Float64,
			WHPerss:          WHPerss.Float64,
			WHTemp:           WHTemp.Float64,
		})
	}
	defer setWellTest(wellTests)
}

func setWellTest(welltest []WellTest) {
	if len(welltest) != 0 {
		for _, well := range welltest {
			date, err := time.Parse("2006-01-02T15:04:05Z", well.Date)
			if err != nil {
				log.Println(err)
			}
			Update := make(map[string]interface{})
			Update["asset_id"] = well.ID
			Update["start_datetime"] = date.Format("2006-01-02 03:04:05")
			Update["cassing_press1"] = well.CassingPressure1
			Update["cassing_press2"] = well.CassingPressure2
			Update["cassing_press3"] = well.CassingPressure3
			Update["cassing_press4"] = well.CassingPressure4
			Update["flow_press"] = well.FlowPress
			Update["pump_current"] = well.PumpCurrent
			Update["pump_efficiency"] = well.PumpEfficiency
			Update["pump_frequency"] = well.PumpFrequency
			Update["pump_rpm"] = well.PumpRPM
			Update["wh_press"] = well.WHPerss
			Update["wh_temp"] = well.WHTemp
			var args []interface{}
			args = append(args, Update)
			id, err := ClientOdoo.ExecuteKw("create", "asset.gauge.new", args, nil)
			if err != nil {
				log.Println(err)
			}
			fmt.Println(id, "Updated")
		}
	} else {
		fmt.Println("+================READINGS=====================+")
		fmt.Println("+=======Nothin to update", time.Now().Format("2006-01-02 03:04:05"), "=======+")
		fmt.Println("+=====================================+")
	}
}

// CHECK WATER VOL LINES
func FindWaterVolsWells() []WellsUpdate {
	var wellsNumbers []WellsUpdate
	wells, err := ClientOdoo.FindAssetAssets(go_odoo.NewCriteria().Add("asset_class_choice", "=", "well"), go_odoo.NewOptions().FetchFields("name", "asset_type_equipment", "water_vol_lines"))
	if err != nil {
		log.Println(err)
	}
	for _, well := range *wells {
		if well.AssetTypeEquipment.Name == "Injection" {
			wellnumber := strings.TrimPrefix(well.Name.Get(), "Скважина №")
			if len(well.WaterVolLines.Get()) != 0 {
				gauges, err := ClientOdoo.GetAssetWaterVol(well.WaterVolLines.Get()[len(well.WaterVolLines.Get())-1])
				if err != nil {
					log.Println(err)
				}
				if gauges != nil {
					wellsNumbers = append(wellsNumbers, WellsUpdate{
						ID:             well.Id.Get(),
						WellType:       "Injection",
						WellNumber:     wellnumber,
						LastUpdateDate: gauges.StartDatetime.Get() + " 00:00:00", // Format("2006-01-02 00:00:00")
					})
				}
			} else {
				wellsNumbers = append(wellsNumbers, WellsUpdate{
					ID:             well.Id.Get(),
					WellType:       "Injection",
					WellNumber:     wellnumber,
					LastUpdateDate: "2012-01-01 00:00:00",
				})
			}
		}
	}
	return wellsNumbers
}

// Закачка воды
func GetWaterVolQuery(wellNumber string, lastUpdate string, assetID int64) {
	conn := db.DatabaseConnection()
	defer conn.Close()
	rows, err := conn.Query(`SELECT ITEM_NAME, START_DATETIME, WATER_VOL
	FROM VT_WELL_TEST_ru_RU
	WHERE ITEM_NAME = ? AND START_DATETIME > ?`, wellNumber, lastUpdate)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	waterVols := []WaterVol{}
	for rows.Next() {
		water := WaterVol{}
		var WATER_VOL sql.NullString
		err := rows.Scan(&water.Name, &water.Date, &WATER_VOL)
		if err != nil {
			log.Println(err)
		}

		Date, err := time.Parse("2006-01-02T15:04:05Z", water.Date)
		if err != nil {
			log.Println(err)
		}
		if WATER_VOL.String == "" {
			WATER_VOL.String = "0"
		}
		waterVols = append(waterVols, WaterVol{
			AssetID:  assetID,
			Name:     water.Name,
			Date:     Date.Format("2006-01-02"),
			WaterVol: WATER_VOL.String,
		})
	}
	SetWaterVol(waterVols)
}

func SetWaterVol(water []WaterVol) {
	if len(water) != 0 {
		for _, v := range water {
			Update := make(map[string]interface{})
			Update["asset_water_vol"] = v.AssetID
			Update["water_vol"] = v.WaterVol
			Update["start_datetime"] = v.Date
			Update["water_vol_uom"] = 20
			var args []interface{}
			args = append(args, Update)
			id, err := ClientOdoo.ExecuteKw("create", "asset.water.vol", args, nil)
			if err != nil {
				log.Println(err)
			}
			fmt.Println(id)
		}
	} else {
		fmt.Println("+==============WATERVOL=======================+")
		fmt.Println("+=======Nothin to update", time.Now().Format("2006-01-02 03:04:05"), "=======+")
		fmt.Println("+=====================================+")
	}
}
