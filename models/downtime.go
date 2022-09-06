package models

type DowntimeUpdate struct {
	AssetID        int64   `json:"asset_iD"`
	ItemName       string  `json:"item_name"`
	Start_datetime string  `json:"start_datetime"`
	EndDatetime    string  `json:"end_datetime"`
	Duration       float64 `json:"duration"`
	DowntimeType   string  `json:"downtime_type"`
	DowntimeText   string  `json:"downtime_text"`
	Comment        string  `json:"comment"`
}
