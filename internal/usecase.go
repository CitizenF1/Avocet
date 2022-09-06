package internal

import (
	"avocet/models"

	go_odoo "github.com/skilld-labs/go-odoo"
)

type UseCase interface {
	GetUpdates(asset go_odoo.AssetAsset, wellnumber string) models.LastUpdates
}
