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

// Запрос в базу данных avocet VT_WELL_TEST_ru_RU Закачка воды по номеру скважины и последней дате обновления
func WaterVolQuery(well models.LastUpdates) error {
	rows, err := db.Conn.Query(`SELECT ITEM_NAME, START_DATETIME, WATER_VOL
	FROM VT_WELL_TEST_ru_RU
	WHERE ITEM_NAME = ? AND START_DATETIME > ?`, well.WellNumber, well.LastUpdateDate)
	if err != nil {
		log.Println(err)
		return err
	}
	defer rows.Close()
	waterVols := []models.WaterVol{}
	for rows.Next() {
		water := models.WaterVol{}
		var WATER_VOL sql.NullString
		err := rows.Scan(&water.Name, &water.Date, &WATER_VOL)
		if err != nil {
			log.Println(err)
			return err
		}
		Date, err := time.Parse("2006-01-02T15:04:05Z", water.Date)
		if err != nil {
			log.Println(err)
			return err
		}
		if WATER_VOL.String == "" {
			WATER_VOL.String = "0"
		}
		waterVols = append(waterVols, models.WaterVol{
			AssetID:  well.AssetID,
			Name:     water.Name,
			Date:     Date.Format("2006-01-02"),
			WaterVol: WATER_VOL.String,
		})
	}
	writeWaterVolOdoo(waterVols)
	return nil
}

func writeWaterVolOdoo(waters []models.WaterVol) {
	// ID единицы измерения закачки воды "m3"
	waterVolUom := 20
	if len(waters) != 0 {
		for _, water := range waters {
			Update := make(map[string]interface{})
			Update["asset_water_vol"] = water.AssetID
			Update["water_vol"] = water.WaterVol
			Update["start_datetime"] = water.Date
			Update["water_vol_uom"] = waterVolUom
			var args []interface{}
			args = append(args, Update)
			id, err := odoo.ClientOdoo.ExecuteKw("create", "asset.water.vol", args, nil)
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
