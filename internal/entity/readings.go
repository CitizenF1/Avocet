package entity

import (
	"avocet/internal/db"
	"avocet/models"
	"avocet/odoo"
	"database/sql"
	"fmt"
	"log"
	"time"
)

// Показания датчиков на нагнетательных скважинах
func InjectReadingsQuery(well models.LastUpdates) error {
	rows, err := db.Conn.Query(`SELECT ITEM_NAME
	,START_DATETIME,CASING_PRESS1,CASING_PRESS2,CASING_PRESS3
,CASING_PRESS4,PUMP_RPM,PUMP_CURRENT,PUMP_FREQUENCY,FLOW_PRESS,PUMP_EFFICIENCY,WH_PRESS,WH_TEMP
FROM VT_WELL_READ_ru_RU
WHERE START_DATETIME > ? AND ITEM_NAME = ?`, well.LastUpdateDate, well.WellNumber)
	if err != nil {
		log.Printf("ERROR SQL Query: %v", err)
		return err
	}
	defer rows.Close()
	wellTests := []models.WellReadings{}
	for rows.Next() {
		var Name, Date string
		var CassingPressure1, CassingPressure2, CassingPressure3, CassingPressure4, PumpRPM,
			PumpCurrent, PumpFrequency, FlowPress, PumpEfficiency, WHPerss, WHTemp sql.NullFloat64

		err := rows.Scan(&Name, &Date, &CassingPressure1, &CassingPressure2, &CassingPressure3,
			&CassingPressure4, &PumpRPM, &PumpCurrent, &PumpFrequency, &FlowPress, &PumpEfficiency, &WHPerss, &WHTemp)
		if err != nil {
			log.Printf("ERROR Scan Rows: %v", err)
			return err
		}
		wellTests = append(wellTests, models.WellReadings{
			AssetID:          well.AssetID,
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
	defer setGauges(wellTests)
	return nil
}

// Запись xlmrpc odoo
func setGauges(welltest []models.WellReadings) {
	if len(welltest) != 0 {
		for _, well := range welltest {
			date, err := time.Parse("2006-01-02T15:04:05Z", well.Date)
			if err != nil {
				log.Println(err)
			}
			Update := make(map[string]interface{})
			Update["asset_id"] = well.AssetID
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
			id, err := odoo.ClientOdoo.ExecuteKw("create", "asset.gauge.new", args, nil)
			if err != nil {
				log.Println(err)
			}
			fmt.Println(id, "Updated")
		}
	} else {
		fmt.Println("+================READINGS=====================+")
		fmt.Println("+=======Nothin to update", time.Now().Format("2006-01-02 03:04:05"), "=========+")
		fmt.Println("+=====================================+")
	}
}
