package odoo

import (
	"fmt"
)

// AssetGaugeNew represents asset.gauge.new model.
type AssetGaugeNew struct {
	AssetId        *Many2One `xmlrpc:"asset_id,omptempty"`
	CassingPress1  *Float    `xmlrpc:"cassing_press1,omptempty"`
	CassingPress2  *Float    `xmlrpc:"cassing_press2,omptempty"`
	CassingPress3  *Float    `xmlrpc:"cassing_press3,omptempty"`
	CassingPress4  *Float    `xmlrpc:"cassing_press4,omptempty"`
	Comment        *String   `xmlrpc:"comment,omptempty"`
	CreateDate     *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid      *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName    *String   `xmlrpc:"display_name,omptempty"`
	FlowPress      *Float    `xmlrpc:"flow_press,omptempty"`
	Id             *Int      `xmlrpc:"id,omptempty"`
	LastUpdate     *Time     `xmlrpc:"__last_update,omptempty"`
	PumpCurrent    *Float    `xmlrpc:"pump_current,omptempty"`
	PumpEfficiency *Float    `xmlrpc:"pump_efficiency,omptempty"`
	PumpFrequency  *Float    `xmlrpc:"pump_frequency,omptempty"`
	PumpRpm        *Float    `xmlrpc:"pump_rpm,omptempty"`
	StartDatetime  *Time     `xmlrpc:"start_datetime,omptempty"`
	WaterOrder     *Float    `xmlrpc:"water_order,omptempty"`
	WhPress        *Float    `xmlrpc:"wh_press,omptempty"`
	WhTemp         *Float    `xmlrpc:"wh_temp,omptempty"`
	WriteDate      *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid       *Many2One `xmlrpc:"write_uid,omptempty"`
}

// AssetGaugeNews represents array of asset.gauge.new model.
type AssetGaugeNews []AssetGaugeNew

// AssetGaugeNewModel is the odoo model name.
const AssetGaugeNewModel = "asset.gauge.new"

// Many2One convert AssetGaugeNew to *Many2One.
func (agn *AssetGaugeNew) Many2One() *Many2One {
	return NewMany2One(agn.Id.Get(), "")
}

// CreateAssetGaugeNew creates a new asset.gauge.new model and returns its id.
func (c *Client) CreateAssetGaugeNew(agn *AssetGaugeNew) (int64, error) {
	return c.Create(AssetGaugeNewModel, agn)
}

// UpdateAssetGaugeNew updates an existing asset.gauge.new record.
func (c *Client) UpdateAssetGaugeNew(agn *AssetGaugeNew) error {
	return c.UpdateAssetGaugeNews([]int64{agn.Id.Get()}, agn)
}

// UpdateAssetGaugeNews updates existing asset.gauge.new records.
// All records (represented by ids) will be updated by agn values.
func (c *Client) UpdateAssetGaugeNews(ids []int64, agn *AssetGaugeNew) error {
	return c.Update(AssetGaugeNewModel, ids, agn)
}

// DeleteAssetGaugeNew deletes an existing asset.gauge.new record.
func (c *Client) DeleteAssetGaugeNew(id int64) error {
	return c.DeleteAssetGaugeNews([]int64{id})
}

// DeleteAssetGaugeNews deletes existing asset.gauge.new records.
func (c *Client) DeleteAssetGaugeNews(ids []int64) error {
	return c.Delete(AssetGaugeNewModel, ids)
}

// GetAssetGaugeNew gets asset.gauge.new existing record.
func (c *Client) GetAssetGaugeNew(id int64) (*AssetGaugeNew, error) {
	agns, err := c.GetAssetGaugeNews([]int64{id})
	if err != nil {
		return nil, err
	}
	if agns != nil && len(*agns) > 0 {
		return &((*agns)[0]), nil
	}
	return nil, fmt.Errorf("id %v of asset.gauge.new not found", id)
}

// GetAssetGaugeNews gets asset.gauge.new existing records.
func (c *Client) GetAssetGaugeNews(ids []int64) (*AssetGaugeNews, error) {
	agns := &AssetGaugeNews{}
	if err := c.Read(AssetGaugeNewModel, ids, nil, agns); err != nil {
		return nil, err
	}
	return agns, nil
}

// FindAssetGaugeNew finds asset.gauge.new record by querying it with criteria.
func (c *Client) FindAssetGaugeNew(criteria *Criteria) (*AssetGaugeNew, error) {
	agns := &AssetGaugeNews{}
	if err := c.SearchRead(AssetGaugeNewModel, criteria, NewOptions().Limit(1), agns); err != nil {
		return nil, err
	}
	if agns != nil && len(*agns) > 0 {
		return &((*agns)[0]), nil
	}
	return nil, fmt.Errorf("asset.gauge.new was not found")
}

// FindAssetGaugeNews finds asset.gauge.new records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAssetGaugeNews(criteria *Criteria, options *Options) (*AssetGaugeNews, error) {
	agns := &AssetGaugeNews{}
	if err := c.SearchRead(AssetGaugeNewModel, criteria, options, agns); err != nil {
		return nil, err
	}
	return agns, nil
}

// FindAssetGaugeNewIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAssetGaugeNewIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AssetGaugeNewModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAssetGaugeNewId finds record id by querying it with criteria.
func (c *Client) FindAssetGaugeNewId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AssetGaugeNewModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("asset.gauge.new was not found")
}
