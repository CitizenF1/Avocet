package app

import (
	"avocet/internal/db"
	"avocet/internal/entity"
	"avocet/models"
	"avocet/odoo"
	"fmt"
	"log"
	"sync"
)

func Run(wg *sync.WaitGroup) {
	odoo.ClientOdoo = *odoo.NewOdooClient()
	db.Conn = db.DatabaseConnection()
	defer db.Conn.Close()
	// Простои Скважин
	wellDowtimes(wg)
	// Устьевые показания
	wellReadings(wg)
	// Закачка воды
	wellWaterVols(wg)
}

func wellReadings(wg *sync.WaitGroup) {
	// функция возвращяет массив структуры LastUpdate по заданным кретериям из odoo
	lastUpdates := entity.GetUpdateByAssetLine("well", "gauge_new_lines")
	for _, well := range lastUpdates {
		wg.Add(1)
		fmt.Println("+=====================================+")
		fmt.Println("|  Checking for update.........  |")
		fmt.Printf("LastUpdate: %s, ID: %v, WellType: %s, Number: %s --Устьевые показания--\n", well.LastUpdateDate, well.AssetID, well.WellType, well.WellNumber)
		fmt.Println("+=====================================+")
		// анонимная функция запрос в базу данных
		go func(well models.LastUpdates) {
			defer wg.Done()
			err := entity.InjectReadingsQuery(well)
			if err != nil {
				log.Println(err)
			}
		}(well)
	}
}

func wellDowtimes(wg *sync.WaitGroup) {
	// функция возвращяет массив структуры LastUpdate по заданным кретериям из odoo
	downtimes := entity.GetUpdateByAssetLine("well", "downtime_ids")
	for _, down := range downtimes {
		wg.Add(1)
		fmt.Println("+=====================================+")
		fmt.Println("|  Checking for update.........  |")
		fmt.Printf("LastUpdate: %s, ID: %v, Number: %s -----Простои Скважин-----\n", down.LastUpdateDate, down.AssetID, down.WellNumber)
		fmt.Println("+=====================================+")
		// анонимная функция запрос в базу данных
		go func(down models.LastUpdates) {
			defer wg.Done()
			err := entity.DowtimeQuery(down)
			if err != nil {
				log.Println(err)
			}
		}(down)
	}
}

func wellWaterVols(wg *sync.WaitGroup) {
	lastUpdates := entity.GetUpdateByAssetLine("well", "water_vol_lines")
	for _, well := range lastUpdates {
		wg.Add(1)
		fmt.Println("+=====================================+")
		fmt.Println("|  Checking for update.........  |")
		fmt.Printf("LastUpdate: %s, ID: %v, Number: %s ----Закачка воды----\n", well.LastUpdateDate, well.AssetID, well.WellNumber)
		fmt.Println("+=====================================+")
		go func(well models.LastUpdates) {
			defer wg.Done()
			err := entity.WaterVolQuery(well)
			if err != nil {
				log.Println(err)
			}
		}(well)
	}
}
