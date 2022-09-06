package entity

import (
	"avocet/internal/db"
	"avocet/models"
	"avocet/odoo"
	"fmt"
	"log"
	"strconv"
	"time"
)

// Запрос в базе данных Avocet просто скважин VT_DOWNTIME_en_US
func DowtimeQuery(well models.LastUpdates) error {
	rows, err := db.Conn.Query(`SELECT ITEM_NAME, START_DATETIME, END_DATETIME, DURATION, DOWNTIME_TYPE, DOWNTIME_TYPE_TEXT, COMMENT
	FROM VT_DOWNTIME_en_US
	WHERE ITEM_NAME = ? AND START_DATETIME > ?`, well.WellNumber, well.LastUpdateDate)
	if err != nil {
		log.Printf("Error Query: %v", err)
		return err
	}
	defer rows.Close()
	downtimes := []models.DowntimeUpdate{}
	for rows.Next() {
		dow := models.DowntimeUpdate{}
		err := rows.Scan(&dow.ItemName, &dow.Start_datetime, &dow.EndDatetime, &dow.Duration, &dow.DowntimeType, &dow.DowntimeText, &dow.Comment)
		if err != nil {
			log.Println(err)
			return err
		}
		downtimes = append(downtimes, models.DowntimeUpdate{
			AssetID:        well.AssetID,
			ItemName:       dow.ItemName,
			Start_datetime: dow.Start_datetime,
			EndDatetime:    dow.EndDatetime,
			Duration:       dow.Duration / 3600,
			DowntimeType:   dow.DowntimeType,
			DowntimeText:   dow.DowntimeText,
			Comment:        dow.Comment,
		})
	}
	setDowtimeReadings(downtimes)
	return nil
}

func setDowtimeReadings(downtimes []models.DowntimeUpdate) {
	if len(downtimes) != 0 {
		for _, downtime := range downtimes {
			Update := make(map[string]interface{})
			start, err := time.Parse("2006-01-02T15:04:05Z", downtime.Start_datetime)
			if err != nil {
				log.Printf("Error parse time: %v", err)
			}
			end, err := time.Parse("2006-01-02T15:04:05Z", downtime.EndDatetime)
			if err != nil {
				log.Printf("Error parse time: %v", err)
			}
			Update["asset_id"] = downtime.AssetID
			Update["start_datetime"] = start.Format("2006-01-02 03:04:05")
			Update["end_datetime"] = end.Format("2006-01-02 03:04:05")
			Update["duration"] = strconv.Itoa(int(downtime.Duration))
			Update["downtime_type"] = downtime.DowntimeType
			if downtime.Comment != "" {
				Update["comment"] = downtime.Comment
			} else {
				Update["comment"] = downtime.DowntimeText
			}
			var args []interface{}
			args = append(args, Update)
			id, err := odoo.ClientOdoo.ExecuteKw("create", "asset.readings.downtime", args, nil)
			if err != nil {
				log.Printf("Error ExecuteKw: %v", err)
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
