package odoo

import (
	"avocet/config"
	"fmt"

	go_odoo "github.com/skilld-labs/go-odoo"
)

var (
	ClientOdoo = go_odoo.Client{}
)

func NewOdooClient() *go_odoo.Client {
	odooClientOdoo, err := go_odoo.NewClient(&go_odoo.ClientConfig{
		Admin:    config.OdooLogin,
		Password: config.OdooPassword,
		Database: config.OdooDatabase,
		URL:      config.OdooUrl,
	})
	if err != nil {
		fmt.Println("[Odoo connection] error: ", err)
	}
	return odooClientOdoo
}
