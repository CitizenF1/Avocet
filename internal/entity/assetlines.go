package entity

import (
	"avocet/models"
	"avocet/odoo"
	"log"
	"regexp"

	go_odoo "github.com/skilld-labs/go-odoo"
)

// запрос в odoo по кретериям тип assettype и assetfields в зависимости от заданных аргументов возвращяется структура lastupdate
func GetUpdateByAssetLine(assetType, assetFields string) []models.LastUpdates {
	var updates []models.LastUpdates
	assets, err := odoo.ClientOdoo.FindAssetAssets(go_odoo.NewCriteria().Add("asset_class_choice", "=", assetType),
		go_odoo.NewOptions().FetchFields("name", "asset_type_equipment", assetFields))
	if err != nil {
		log.Println(err)
	}
	for _, asset := range *assets {
		if asset.AssetTypeEquipment.Name == "Injection" || assetFields == "downtime_ids" {
			reg := regexp.MustCompile(`\d+`)
			wellnumber := reg.FindString(asset.Name.Get())
			switch assetFields {
			case "water_vol_lines":
				updates = append(updates, waterVolLines(asset, wellnumber))
			case "gauge_new_lines":
				updates = append(updates, gaugesLines(asset, wellnumber))
			case "downtime_ids":
				updates = append(updates, downtimeIds(asset, wellnumber))
			}
		}
	}
	return updates
}

//lastupdate обновления в модели asset.readings.downtime
func downtimeIds(asset go_odoo.AssetAsset, wellnumber string) models.LastUpdates {
	var updates models.LastUpdates
	if len(asset.DowntimeIds.Get()) != 0 {
		downtime, err := odoo.ClientOdoo.GetAssetReadingsDowntime(asset.DowntimeIds.Get()[len(asset.DowntimeIds.Get())-1])
		if err != nil {
			log.Println(err)
		}
		updates = models.LastUpdates{
			AssetID:        asset.Id.Get(),
			WellNumber:     wellnumber,
			LastUpdateDate: downtime.StartDatetime.Get().Format("2006-01-02 03:04:05"),
		}
	} else {
		updates = models.LastUpdates{
			AssetID:        asset.Id.Get(),
			WellNumber:     wellnumber,
			LastUpdateDate: "1980-01-01 00:00:00",
		}
	}
	return updates
}

// lastupdate обновления в модели asset.water.vol
func waterVolLines(asset go_odoo.AssetAsset, wellnumber string) models.LastUpdates {
	var update models.LastUpdates
	if len(asset.WaterVolLines.Get()) != 0 {
		waterVol, err := odoo.ClientOdoo.GetAssetWaterVol(asset.WaterVolLines.Get()[len(asset.WaterVolLines.Get())-1])
		if err != nil {
			log.Println(err)
		}
		if waterVol != nil {
			update = models.LastUpdates{
				AssetID:        asset.Id.Get(),
				WellType:       "Injection",
				WellNumber:     wellnumber,
				LastUpdateDate: waterVol.StartDatetime.Get() + " 00:00:00",
			}
		}
	} else {
		update = models.LastUpdates{
			AssetID:        asset.Id.Get(),
			WellType:       "Injection",
			WellNumber:     wellnumber,
			LastUpdateDate: "2012-01-01 00:00:00",
		}
	}
	return update
}

// lastupdate обновления в модели asset.gauge.new
func gaugesLines(asset go_odoo.AssetAsset, wellnumber string) models.LastUpdates {
	var updates models.LastUpdates
	if len(asset.GaugeNewLine.Get()) != 0 {
		gauges, err := odoo.ClientOdoo.GetAssetGaugeNew(asset.GaugeNewLine.Get()[len(asset.GaugeNewLine.Get())-1])
		if err != nil {
			log.Println(err)
		}
		if gauges != nil {
			updates = models.LastUpdates{
				AssetID:        asset.Id.Get(),
				WellType:       "Injection",
				WellNumber:     wellnumber,
				LastUpdateDate: gauges.StartDatetime.Get().Format("2006-01-02 00:00:00"),
			}
		}
	} else {
		return models.LastUpdates{
			AssetID:        asset.Id.Get(),
			WellType:       "Injection",
			WellNumber:     wellnumber,
			LastUpdateDate: "2012-01-01 00:00:00",
		}
	}
	return updates
}
