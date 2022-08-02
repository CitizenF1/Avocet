package main

import (
	"fmt"
	"strconv"
	"time"
)

// TODO: add variable to set in query ITEM_NAME
// Простои скважин
func GetDowtimeReadingsQuery() {
	conn := databaseConnection()
	defer conn.Close()
	tsql := fmt.Sprintf(`SELECT ITEM_NAME
	,START_DATETIME
	,END_DATETIME
	,DURATION
	,DOWNTIME_TYPE
	,DOWNTIME_TYPE_TEXT
	,COMMENT
FROM VT_DOWNTIME_en_US
WHERE ITEM_NAME = '101'`) // HERE Replace '101'
	rows, err := conn.Query(tsql)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	downtimes := []Downtime{}
	for rows.Next() {
		dow := Downtime{}
		err := rows.Scan(&dow.Item_name, &dow.Start_datetime, &dow.End_datetime, &dow.Duration, &dow.Downtime_type, &dow.Downtime_text, &dow.Comment)
		if err != nil {
			// fmt.Println(err)
		}
		downtimes = append(downtimes, Downtime{
			Item_name:      dow.Item_name,
			Start_datetime: dow.Start_datetime,
			End_datetime:   dow.End_datetime,
			Duration:       dow.Duration / 3600,
			Downtime_type:  dow.Downtime_type,
			Downtime_text:  dow.Downtime_text,
			Comment:        dow.Comment,
		})
		// fmt.Printf("Item_name: %s, Start: %s, End: %s, Duration: %v, Type: %s, Text: %s, Comment: %s \n", dow.Item_name, dow.Start_datetime, dow.End_datetime, dow.Duration/3600, dow.Downtime_type, dow.Downtime_text, dow.Comment)
	}
	setDowtimeReadings(downtimes)
}

func setDowtimeReadings(downtime []Downtime) {
	for _, v := range downtime {
		// fmt.Printf("Item_name: %s, Start: %s, End: %s, Duration: %v, Type: %s, Text: %s, Comment: %s \n", v.Item_name, v.Start_datetime, v.End_datetime, v.Duration, v.Downtime_type, v.Downtime_text, v.Comment)
		Update := make(map[string]interface{})
		start, err := time.Parse("2006-01-02T15:04:05Z", v.Start_datetime)
		if err != nil {
			fmt.Println(err)
		}
		end, err := time.Parse("2006-01-02T15:04:05Z", v.End_datetime)
		if err != nil {
			fmt.Println(err)
		}
		Update["asset_id"] = 2
		Update["start_datetime"] = start.Format("2006-01-02 03:04:05") // Add(-time.Hour * 6).
		Update["end_datetime"] = end.Format("2006-01-02 03:04:05")     //Add(-time.Hour * 6).
		Update["duration"] = strconv.Itoa(int(v.Duration))
		Update["downtime_type"] = v.Downtime_type
		if v.Comment != "" {
			Update["comment"] = v.Comment
		} else {
			Update["comment"] = v.Downtime_text
		}
		var args []interface{}
		args = append(args, Update)
		id, err := ClientOdoo.ExecuteKw("create", "asset.readings.downtime", args, nil)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(id)
	}
}

// var dur float64 = 16
// 	Update := make(map[string]interface{})
// 	str := "2021-09-02T09:00:00"
// 	t, err := time.Parse("2006-01-02T15:04:05", str)
// 	fmt.Println(t.Add(-time.Hour * 6).Format("2006-01-02 03:04:05")) //2006-01-02 //2006-01-02 03:04:05 //2006-01-02 15:04:0
// 	Update["asset_id"] = 2
// 	Update["start_datetime"] = t.Add(-time.Hour * 6).Format("2006-01-02 03:04:05")
// 	Update["duration"] = strconv.Itoa(int(dur))
// 	var args []interface{}
// 	args = append(args, Update)
// 	id, err := ClientOdoo.ExecuteKw("create", "asset.readings.downtime", args, nil)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(id)
