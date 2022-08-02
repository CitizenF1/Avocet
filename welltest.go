package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

// Показания датчиков на нагнетательных скважинах
func GetWellReadings() {
	conn := databaseConnection()
	defer conn.Close()
	tsql := fmt.Sprintf(`SELECT ITEM_NAME
	,START_DATETIME
	,CASING_PRESS1
	,CASING_PRESS2
	,CASING_PRESS3
	,CASING_PRESS4
	,PUMP_RPM
	,PUMP_CURRENT
	,PUMP_FREQUENCY
	,FLOW_PRESS
	,PUMP_EFFICIENCY
	,WH_PRESS
	,WH_TEMP
FROM VT_WELL_READ_ru_RU
WHERE ITEM_NAME = '151'`)
	rows, err := conn.Query(tsql)
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
			// fmt.Println(err)
		}
		// fmt.Printf("Item_name: %s, date: %s, cassing1: %v, cassing2: %v, cassing3: %v, cassing4: %v, PumpRPM: %v, PumpCurrent: %v, PumpFrequency: %v,  FlowPress: %v, PumpEfficiency: %v, WHPerss: %v, WHTemp: %v \n", w.Name, w.Date, w.CassingPressure1, w.CassingPressure2, w.CassingPressure3, w.CassingPressure4, w.PumpRPM, w.PumpCurrent, w.PumpFrequency, w.FlowPress, w.PumpEfficiency, w.WHPerss, w.WHTemp)
		wellTests = append(wellTests, WellTest{
			Name:             Name, // CHANGE TO ID
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
	// b, err := json.Marshal(wellTests)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// err = ioutil.WriteFile("./Test.json", b, 0666)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	defer setWellTest(wellTests)
}

// TODO: Change asset_id pass it as an argument to a function or CHANGE Well.Name TO Well.ID
func setWellTest(welltest []WellTest) {
	for _, well := range welltest {
		date, err := time.Parse("2006-01-02T15:04:05Z", well.Date)
		if err != nil {
			fmt.Println(err)
		}
		Update := make(map[string]interface{})
		Update["asset_id"] = 16 //HERE CHANGE TO ID
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
			fmt.Println(err)
		}
		fmt.Println(id)
	}
}

// TODO: add variable to set in query ITEM_NAME
// суточная закачка воды
func GetWaterVolQuery() {
	conn := databaseConnection()
	defer conn.Close()
	tsql := fmt.Sprintf(`SELECT ITEM_NAME
	,START_DATETIME
	,WATER_VOL
FROM VT_WELL_TEST_ru_RU
WHERE ITEM_NAME = '101'`) // HERE: replace '188'
	rows, err := conn.Query(tsql)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	waterVols := []WaterVol{}
	for rows.Next() {
		water := WaterVol{}
		err := rows.Scan(&water.Name, &water.Date, &water.WaterVol)
		if err != nil {
			fmt.Println(err)
		}

		Date, err := time.Parse("2006-01-02T15:04:05Z", water.Date)
		if err != nil {
			fmt.Println(err)
		}

		waterVols = append(waterVols, WaterVol{
			Name:     water.Name,
			Date:     Date.Format("2006-01-02"),
			WaterVol: water.WaterVol,
		})
	}
	b, err := json.Marshal(waterVols)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile("./Test.json", b, 0666)
	if err != nil {
		fmt.Println(err)
	}
	// SetWaterVol(waterVols)
}

func SetWaterVol(water []WaterVol) {
	for _, v := range water {
		Update := make(map[string]interface{})
		Update["asset_water_vol"] = 2 // CHANGE ID
		Update["water_vol"] = v.WaterVol
		Update["start_datetime"] = v.Date
		Update["water_vol_uom"] = 20
		var args []interface{}
		args = append(args, Update)
		id, err := ClientOdoo.ExecuteKw("create", "asset.water.vol", args, nil)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(id)
	}
}
