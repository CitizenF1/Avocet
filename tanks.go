package main

import "fmt"

// 	SELECT TOP (1000) [LEGACY_ID]
//       ,[ITEM_NAME]
//       ,[START_DATETIME]
//       ,[OIL_MASS]
//   FROM [AVM_CN].[dbo].[CN_INV_PSTN]
//   ORDER BY [START_DATETIME] desc

func getTanksReadingsQuery() {
	conn := databaseConnection()
	defer conn.Close()
	tsql := fmt.Sprintf(`SELECT TOP (100) ITEM_NAME
	,START_DATETIME
	,PRESSURE
	,TEMPERATURE
	,LEVEL
	,LEVEL_PH
	,FLW_OIL_DENS
FROM VT_INVENTORY_DAY_ru_RU
WHERE (LEVEL IS NOT NULL AND PRESSURE IS NOT NULL)
ORDER BY START_DATETIME desc`)
	rows, err := conn.Query(tsql)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		var pstn Pstn
		err := rows.Scan(&pstn.Item_name, &pstn.Start_datetime, &pstn.Pressure, &pstn.Temperature, &pstn.Level, &pstn.Level_ph, &pstn.Oil_dens)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("ItemName: %s, Date: %s, Pressure: %s, Tempeture: %s, Level: %s, Level_PH: %s, Oil_dens: %s \n", pstn.Item_name, pstn.Start_datetime, pstn.Pressure, pstn.Temperature, pstn.Level, pstn.Level_ph, pstn.Oil_dens)
	}
}

func TanksReadings(pstn []Pstn) {

}
