package odoo

import (
	"fmt"
)

// AssetFactIndicators represents asset.fact.indicators model.
type AssetFactIndicators struct {
	AssetFactIndicator *Many2One `xmlrpc:"asset_fact_indicator,omptempty"`
	BswVolFrac         *String   `xmlrpc:"bsw_vol_frac,omptempty"`
	CreateDate         *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid          *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName        *String   `xmlrpc:"display_name,omptempty"`
	FactIndicatorsUom  *Many2One `xmlrpc:"fact_indicators_uom,omptempty"`
	FlwOilDens         *String   `xmlrpc:"flw_oil_dens,omptempty"`
	FwLvl              *String   `xmlrpc:"fw_lvl,omptempty"`
	Id                 *Int      `xmlrpc:"id,omptempty"`
	LastUpdate         *Time     `xmlrpc:"__last_update,omptempty"`
	Level              *String   `xmlrpc:"level,omptempty"`
	LevelPh            *String   `xmlrpc:"level_ph,omptempty"`
	Name               *String   `xmlrpc:"name,omptempty"`
	ObsTemp            *String   `xmlrpc:"obs_temp,omptempty"`
	StartDatetime      *String   `xmlrpc:"start_datetime,omptempty"`
	Temperature        *String   `xmlrpc:"temperature,omptempty"`
	WriteDate          *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid           *Many2One `xmlrpc:"write_uid,omptempty"`
}

// AssetFactIndicatorss represents array of asset.fact.indicators model.
type AssetFactIndicatorss []AssetFactIndicators

// AssetFactIndicatorsModel is the odoo model name.
const AssetFactIndicatorsModel = "asset.fact.indicators"

// Many2One convert AssetFactIndicators to *Many2One.
func (afi *AssetFactIndicators) Many2One() *Many2One {
	return NewMany2One(afi.Id.Get(), "")
}

// CreateAssetFactIndicators creates a new asset.fact.indicators model and returns its id.
func (c *Client) CreateAssetFactIndicators(afi *AssetFactIndicators) (int64, error) {
	return c.Create(AssetFactIndicatorsModel, afi)
}

// UpdateAssetFactIndicators updates an existing asset.fact.indicators record.
func (c *Client) UpdateAssetFactIndicators(afi *AssetFactIndicators) error {
	return c.UpdateAssetFactIndicatorss([]int64{afi.Id.Get()}, afi)
}

// UpdateAssetFactIndicatorss updates existing asset.fact.indicators records.
// All records (represented by ids) will be updated by afi values.
func (c *Client) UpdateAssetFactIndicatorss(ids []int64, afi *AssetFactIndicators) error {
	return c.Update(AssetFactIndicatorsModel, ids, afi)
}

// DeleteAssetFactIndicators deletes an existing asset.fact.indicators record.
func (c *Client) DeleteAssetFactIndicators(id int64) error {
	return c.DeleteAssetFactIndicatorss([]int64{id})
}

// DeleteAssetFactIndicatorss deletes existing asset.fact.indicators records.
func (c *Client) DeleteAssetFactIndicatorss(ids []int64) error {
	return c.Delete(AssetFactIndicatorsModel, ids)
}

// GetAssetFactIndicators gets asset.fact.indicators existing record.
func (c *Client) GetAssetFactIndicators(id int64) (*AssetFactIndicators, error) {
	afis, err := c.GetAssetFactIndicatorss([]int64{id})
	if err != nil {
		return nil, err
	}
	if afis != nil && len(*afis) > 0 {
		return &((*afis)[0]), nil
	}
	return nil, fmt.Errorf("id %v of asset.fact.indicators not found", id)
}

// GetAssetFactIndicatorss gets asset.fact.indicators existing records.
func (c *Client) GetAssetFactIndicatorss(ids []int64) (*AssetFactIndicatorss, error) {
	afis := &AssetFactIndicatorss{}
	if err := c.Read(AssetFactIndicatorsModel, ids, nil, afis); err != nil {
		return nil, err
	}
	return afis, nil
}

// FindAssetFactIndicators finds asset.fact.indicators record by querying it with criteria.
func (c *Client) FindAssetFactIndicators(criteria *Criteria) (*AssetFactIndicators, error) {
	afis := &AssetFactIndicatorss{}
	if err := c.SearchRead(AssetFactIndicatorsModel, criteria, NewOptions().Limit(1), afis); err != nil {
		return nil, err
	}
	if afis != nil && len(*afis) > 0 {
		return &((*afis)[0]), nil
	}
	return nil, fmt.Errorf("asset.fact.indicators was not found")
}

// FindAssetFactIndicatorss finds asset.fact.indicators records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAssetFactIndicatorss(criteria *Criteria, options *Options) (*AssetFactIndicatorss, error) {
	afis := &AssetFactIndicatorss{}
	if err := c.SearchRead(AssetFactIndicatorsModel, criteria, options, afis); err != nil {
		return nil, err
	}
	return afis, nil
}

// FindAssetFactIndicatorsIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAssetFactIndicatorsIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AssetFactIndicatorsModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAssetFactIndicatorsId finds record id by querying it with criteria.
func (c *Client) FindAssetFactIndicatorsId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AssetFactIndicatorsModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("asset.fact.indicators was not found")
}
