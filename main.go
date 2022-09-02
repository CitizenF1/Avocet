package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	_ "github.com/denisenkom/go-mssqldb"
	go_odoo "github.com/skilld-labs/go-odoo"
)

var wg sync.WaitGroup

func main() {
	// ====Route log to file====
	f, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)
	odooClientOdoo, err := go_odoo.NewClient(&go_odoo.ClientConfig{
		Admin:    odooLogin,
		Password: odooPassword,
		Database: odooDatabase,
		URL:      odooUrl,
	})
	if err != nil {
		fmt.Println("[Odoo connection] error: ", err)
	}
	ClientOdoo = *odooClientOdoo

	WorkWellDowtimes()

	WorkWellReadings()

	WorkWellWaterVols()

	wg.Wait()
	fmt.Println("Done")
}

func WorkWellReadings() {
	wells := FindInjectionWells()
	for _, v := range wells {
		wg.Add(1)
		fmt.Println("+=====================================+")
		fmt.Println("|  Checking for update.........  |")
		fmt.Printf("LastUpdate: %s, ID: %v, WellType: %s, Number: %s --Устьевые показания--\n", v.LastUpdateDate, v.ID, v.WellType, v.WellNumber)
		fmt.Println("+=====================================+")
		go func(v WellsUpdate) {
			defer wg.Done()
			GetWellInjectReadings(v.WellNumber, v.LastUpdateDate, v.ID)
		}(v)
	}
}

func WorkWellDowtimes() {
	wells := GetDowntimeWellsOdoo()
	for _, v := range wells {
		wg.Add(1)
		fmt.Println("+=====================================+")
		fmt.Println("|  Checking for update.........  |")
		fmt.Printf("LastUpdate: %s, ID: %v, Number: %s -----Простои Скважин-----\n", v.LastUpdateDate, v.ID, v.WellNumber)
		fmt.Println("+=====================================+")
		go func(v DowntimeWellsOdoo) {
			defer wg.Done()
			GetDowtimeReadings(v.WellNumber, v.LastUpdateDate, v.ID)
		}(v)
	}
}

func WorkWellWaterVols() {
	wells := FindWaterVolsWells()
	for _, v := range wells {
		wg.Add(1)
		fmt.Println("+=====================================+")
		fmt.Println("|  Checking for update.........  |")
		fmt.Printf("LastUpdate: %s, ID: %v, Number: %s ----Закачка воды----\n", v.LastUpdateDate, v.ID, v.WellNumber)
		fmt.Println("+=====================================+")
		go func(v WellsUpdate) {
			defer wg.Done()
			GetWaterVolQuery(v.WellNumber, v.LastUpdateDate, v.ID)
		}(v)
	}
}
