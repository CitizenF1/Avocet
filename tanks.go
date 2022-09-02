package main

import (
	"database/sql"
	"fmt"
	"rest/db"
	"strings"

	go_odoo "github.com/skilld-labs/go-odoo"
)

type Tank struct {
	ID         int64
	Key        string
	LastUpdate string
}

func FindTanks() []Tank {
	var tanks []Tank
	assets, err := ClientOdoo.FindAssetAssets(go_odoo.NewCriteria().Add("asset_class_choice", "=", "rvs"), go_odoo.NewOptions().FetchFields("name", "key", "fact_indicators_lines"))
	if err != nil {
		fmt.Println(err)
	}
	for _, asset := range *assets {
		if len(asset.FactIndicatorsLines.Get()) != 0 {
			indicator, err := ClientOdoo.GetAssetFactIndicators(asset.FactIndicatorsLines.Get()[len(asset.FactIndicatorsLines.Get())-1])
			if err != nil {
				fmt.Println(err)
			}
			if indicator != nil {
				tanks = append(tanks, Tank{
					ID:         asset.Id.Get(),
					Key:        asset.Key.Get(),
					LastUpdate: indicator.StartDatetime.Get(),
				})
			}
		} else {
			tanks = append(tanks, Tank{
				ID:         asset.Id.Get(),
				Key:        asset.Key.Get(),
				LastUpdate: "2012-01-01 00:00:00",
			})
		}
	}
	return tanks
}

func getTanksReadingsQuery(itemName, date string, assetID int64) {
	itemName = strings.Replace(itemName, " №", "-", 1)
	conn := db.DatabaseConnection()
	defer conn.Close()
	rows, err := conn.Query(`SELECT ITEM_NAME, START_DATETIME, PRESSURE, TEMPERATURE, LEVEL, LEVEL_PH, FLW_OIL_DENS
	FROM VT_INVENTORY_DAY_ru_RU
	WHERE ITEM_NAME LIKE ? AND START_DATETIME > ?`, "%"+itemName, date)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	pstns := []Pstn{}
	for rows.Next() {
		var ITEM_NAME, START_DATETIME, PRESSURE, TEMPERATURE, LEVEL, LEVEL_PH, FLW_OIL_DENS sql.NullString
		err := rows.Scan(&ITEM_NAME, &START_DATETIME, &PRESSURE, &TEMPERATURE, &LEVEL, &LEVEL_PH, &FLW_OIL_DENS)
		if err != nil {
			fmt.Println(err)
		}
		// fmt.Printf("ItemName: %s, Date: %s, Pressure: %s, Tempeture: %s, Level: %s, Level_PH: %s, Oil_dens: %s \n", ITEM_NAME.String, START_DATETIME.String, PRESSURE.String, TEMPERATURE.String, LEVEL, LEVEL_PH.String, FLW_OIL_DENS.String)
		pstns = append(pstns, Pstn{
			AssetID:        assetID,
			Item_name:      ITEM_NAME.String,
			Start_datetime: START_DATETIME.String,
			Pressure:       PRESSURE.String,
			Temperature:    TEMPERATURE.String,
			Level:          LEVEL.String,
			Level_ph:       LEVEL_PH.String,
			Oil_dens:       FLW_OIL_DENS.String,
		})
	}
}

//fact_indicators_uom: 20
func TanksReadings(pstn []Pstn) {

}
