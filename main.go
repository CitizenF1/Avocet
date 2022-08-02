package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
	go_odoo "github.com/skilld-labs/go-odoo"
)

var (
	ClientOdoo go_odoo.Client
)

var (
	odooLogin    = "LOGIN"
	odooPassword = "PASSWORD"
	odooDatabase = "DATABASE"
	odooUrl      = "URL"
)

var (
	server   = "SERVERADDRES"
	port     = 1433
	user     = "USER"
	password = "PASSWORD"
	database = "DATABASE"
)

func main() {
	odooClientOdoo, err := go_odoo.NewClient(&go_odoo.ClientConfig{
		Admin:    odooLogin,
		Password: odooPassword,
		Database: odooDatabase,
		URL:      odooUrl,
	})
	if err != nil {
		fmt.Println("[Odoo connection] error: ", err)
		// return err
	}
	ClientOdoo = *odooClientOdoo

	// GetWellReadings()
	// GetWaterVolQuery()
	// getDowtimeReadingsQuery()
	// DailyProductionWells()
	// wellsPressureQuery()
	// ofm_master()
	// getPSTN()
	// WellsPressureQuery()
}

func databaseConnection() *sql.DB {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)

	conn, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	fmt.Printf("Connected!\n")
	return conn
}

type DowntimeType struct {
	Downtime_type string
	Downtime_text string
}

// SELECT DISTINCT [DOWNTIME_TYPE]
//       ,[DOWNTIME_TYPE_TEXT]
//   FROM [AVM_CN].[dbo].[VT_DOWNTIME_en_US]

//--------CREATE тип насоса----------
// value := make(map[string]interface{})
// value["name"] = "ПНШ-60"
// value["production_method"] = "2"
// id, err := ClientOdoo.Create("asset.pump.type", value)
// if err != nil {
// 	fmt.Println(err)
// }
// fmt.Println(id)
