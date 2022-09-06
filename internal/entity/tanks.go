package entity

import (
	"avocet/internal/db"
	"avocet/models"
	"avocet/odoo"
	"database/sql"
	"fmt"
	"log"
	"strings"

	go_odoo "github.com/skilld-labs/go-odoo"
)

func FindTanks() []models.LastUpdates {
	var tanks []models.LastUpdates
	assets, err := odoo.ClientOdoo.FindAssetAssets(go_odoo.NewCriteria().Add("asset_class_choice", "=", "rvs"),
		go_odoo.NewOptions().FetchFields("name", "key", "fact_indicators_lines"))
	if err != nil {
		log.Println(err)
	}
	for _, asset := range *assets {
		if len(asset.FactIndicatorsLines.Get()) != 0 {
			indicator, err := odoo.ClientOdoo.GetAssetFactIndicators(asset.FactIndicatorsLines.Get()[len(asset.FactIndicatorsLines.Get())-1])
			if err != nil {
				fmt.Println(err)
			}
			if indicator != nil {
				tanks = append(tanks, models.LastUpdates{
					AssetID:        asset.Id.Get(),
					Key:            asset.Key.Get(),
					LastUpdateDate: indicator.StartDatetime.Get(),
				})
			}
		} else {
			tanks = append(tanks, models.LastUpdates{
				AssetID:        asset.Id.Get(),
				Key:            asset.Key.Get(),
				LastUpdateDate: "2012-01-01 00:00:00",
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
	pstns := []models.Tanks{}
	for rows.Next() {
		var ITEM_NAME, START_DATETIME, PRESSURE, TEMPERATURE, LEVEL, LEVEL_PH, FLW_OIL_DENS sql.NullString
		err := rows.Scan(&ITEM_NAME, &START_DATETIME, &PRESSURE, &TEMPERATURE, &LEVEL, &LEVEL_PH, &FLW_OIL_DENS)
		if err != nil {
			fmt.Println(err)
		}
		pstns = append(pstns, models.Tanks{
			AssetID:       assetID,
			ItemName:      ITEM_NAME.String,
			StartDatetime: START_DATETIME.String,
			Pressure:      PRESSURE.String,
			Temperature:   TEMPERATURE.String,
			Level:         LEVEL.String,
			LevelPh:       LEVEL_PH.String,
			OilDens:       FLW_OIL_DENS.String,
		})
	}
}
