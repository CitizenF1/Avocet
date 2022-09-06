package models

type Tanks struct {
	AssetID       int64  `json:"asset_iD"`
	ItemName      string `json:"item_name"`
	StartDatetime string `json:"start_datetime"`
	Pressure      string `json:"pressure"`
	Temperature   string `json:"temperature"`
	Level         string `json:"level"`
	LevelPh       string `json:"level_ph"`
	OilDens       string `json:"oil_dens"`
}
