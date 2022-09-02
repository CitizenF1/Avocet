package main

import (
	"fmt"
	"log"
	"rest/db"
	"strconv"
	"strings"
	"time"

	go_odoo "github.com/skilld-labs/go-odoo"
)

type DowntimeWellsOdoo struct {
	ID             int64
	WellNumber     string
	LastUpdateDate string
}

func GetDowntimeWellsOdoo() []DowntimeWellsOdoo {
	var wellsDowntimes []DowntimeWellsOdoo
	downtimeIDs, err := ClientOdoo.FindAssetAssets(go_odoo.NewCriteria().Add("asset_class_choice", "=", "well"), go_odoo.NewOptions().FetchFields("id", "name", "downtime_ids"))
	if err != nil {
		fmt.Println(err, "---------")
	}
	for _, v := range *downtimeIDs {
		wellnumber := strings.TrimPrefix(v.Name.Get(), "Скважина №")
		if len(v.DowntimeIds.Get()) != 0 {
			dow, err := ClientOdoo.GetAssetReadingsDowntime(v.DowntimeIds.Get()[len(v.DowntimeIds.Get())-1])
			if err != nil {
				log.Println(err)
			}
			wellsDowntimes = append(wellsDowntimes, DowntimeWellsOdoo{
				ID:             v.Id.Get(),
				WellNumber:     wellnumber,
				LastUpdateDate: dow.StartDatetime.Get().Format("2006-01-02 03:04:05"),
			})
		} else {
			wellsDowntimes = append(wellsDowntimes, DowntimeWellsOdoo{
				ID:             v.Id.Get(),
				WellNumber:     wellnumber,
				LastUpdateDate: "1980-01-01 00:00:00",
			})
		}
	}
	return wellsDowntimes
}

func GetDowtimeReadings(wellNuber string, lastUpdate string, assetID int64) {
	conn := db.DatabaseConnection()
	defer conn.Close()
	rows, err := conn.Query(`SELECT ITEM_NAME, START_DATETIME, END_DATETIME, DURATION, DOWNTIME_TYPE, DOWNTIME_TYPE_TEXT, COMMENT
	FROM VT_DOWNTIME_en_US
	WHERE ITEM_NAME = ? AND START_DATETIME > ?`, wellNuber, lastUpdate)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	downtimes := []DowntimeUpdate{}
	for rows.Next() {
		dow := DowntimeUpdate{}
		err := rows.Scan(&dow.Item_name, &dow.Start_datetime, &dow.End_datetime, &dow.Duration, &dow.Downtime_type, &dow.Downtime_text, &dow.Comment)
		if err != nil {
			log.Println(err)
		}
		downtimes = append(downtimes, DowntimeUpdate{
			AssetID:        assetID,
			Item_name:      dow.Item_name,
			Start_datetime: dow.Start_datetime,
			End_datetime:   dow.End_datetime,
			Duration:       dow.Duration / 3600,
			Downtime_type:  dow.Downtime_type,
			Downtime_text:  dow.Downtime_text,
			Comment:        dow.Comment,
		})
	}
	setDowtimeReadings(downtimes)
}

func setDowtimeReadings(downtime []DowntimeUpdate) {
	if len(downtime) != 0 {
		for _, v := range downtime {
			Update := make(map[string]interface{})
			start, err := time.Parse("2006-01-02T15:04:05Z", v.Start_datetime)
			if err != nil {
				log.Println(err)
			}
			end, err := time.Parse("2006-01-02T15:04:05Z", v.End_datetime)
			if err != nil {
				log.Println(err)
			}
			Update["asset_id"] = v.AssetID
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
				log.Println(err)
			}
			fmt.Println("+==============DOWTIMES UPDATE=======================+")
			fmt.Println(id)
			fmt.Println("+====================================================+")
		}
	} else {
		fmt.Println("+==============DOWTIMES=======================+")
		fmt.Println("+=======Nothin to update", time.Now().Format("2006-01-02 03:04:05"), "=======+")
		fmt.Println("+=====================================+")
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
