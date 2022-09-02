package main

import (
	"fmt"
	"log"
	"os"
	"rest/db"
	"strconv"
	"time"

	"github.com/xuri/excelize/v2"

	"github.com/joho/godotenv"

	go_odoo "github.com/skilld-labs/go-odoo"
)

var (
	ClientOdoo go_odoo.Client
)

var (
	odooLogin    string
	odooPassword string
	odooDatabase string
	odooUrl      string
)

func SumArray(arr []int) int {
	sum := 0
	for _, value := range arr {
		sum += value
	}
	return sum
}

func init() {
	err := godotenv.Overload(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	odooLogin = os.Getenv("ODOOLOGIN")
	odooPassword = os.Getenv("ODOOPASSWORD")
	odooDatabase = os.Getenv("ODOODATABASE")
	odooUrl = os.Getenv("ODOOURL")

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		fmt.Println(err)
	}
	db.Server = os.Getenv("SERVER")
	db.Password = os.Getenv("PASSWORD")
	db.Port = port
	db.User = os.Getenv("USER")
	db.Database = os.Getenv("DATABASE")
}

type Production struct {
	Name           string
	StartDate      string
	EndDate        string
	Target         string
	Measured       string
	TotalShorfall  string
	ExplShortFall  string
	UnexpShartFall string
}

func Prod() {
	conn := db.DatabaseConnection()
	defer conn.Close()
	rows, err := conn.Query(`SELECT ITEM_NAME
	,START_DATETIME
	,END_DATETIME
	,TARGET
	,MEASURED
	,TOTAL_SHORTFALL
	,EXPL_SHORTFALL
	,UNEXP_SHORTFALL
FROM VT_WL_SF_OILMASSDAY_ru_RU
WHERE ITEM_NAME = '190'`)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	prod := []Production{}
	for rows.Next() {
		var ITEM_NAME, START_DATETIME, END_DATETIME, TARGET, MEASURED, TOTAL_SHORTFALL, EXPL_SHORTFALL, UNEXP_SHORTFALL string
		err := rows.Scan(&ITEM_NAME, &START_DATETIME, &END_DATETIME, &TARGET, &MEASURED, &TOTAL_SHORTFALL, &EXPL_SHORTFALL, &UNEXP_SHORTFALL)
		if err != nil {
			fmt.Println(err)
		}
		prod = append(prod, Production{
			Name:           ITEM_NAME,
			StartDate:      START_DATETIME,
			EndDate:        END_DATETIME,
			Target:         TARGET,
			Measured:       MEASURED,
			TotalShorfall:  TOTAL_SHORTFALL,
			ExplShortFall:  EXPL_SHORTFALL,
			UnexpShartFall: UNEXP_SHORTFALL,
		})
	}
	defer Write(prod)
}

func Write(prod []Production) {
	f, err := excelize.OpenFile("Production.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	for i := 0; i < len(prod); i++ {
		start, err := time.Parse("2006-01-02T15:04:05Z", prod[i].StartDate)
		if err != nil {
			fmt.Println(err)
		}
		end, err := time.Parse("2006-01-02T15:04:05Z", prod[i].EndDate)
		if err != nil {
			fmt.Println(err)
		}
		f.SetCellValue("PRODUCTION", "A"+strconv.Itoa(i+2), prod[i].Name)
		f.SetCellValue("PRODUCTION", "B"+strconv.Itoa(i+2), start.Format("2006-01-02"))
		f.SetCellValue("PRODUCTION", "C"+strconv.Itoa(i+2), end.Format("2006-01-02"))
		f.SetCellValue("PRODUCTION", "D"+strconv.Itoa(i+2), prod[i].Target)
		f.SetCellValue("PRODUCTION", "E"+strconv.Itoa(i+2), prod[i].Measured)
		f.SetCellValue("PRODUCTION", "F"+strconv.Itoa(i+2), prod[i].TotalShorfall)
		f.SetCellValue("PRODUCTION", "G"+strconv.Itoa(i+2), prod[i].ExplShortFall)
		f.SetCellValue("PRODUCTION", "H"+strconv.Itoa(i+2), prod[i].UnexpShartFall)
	}
	err = f.SaveAs("Production.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}
