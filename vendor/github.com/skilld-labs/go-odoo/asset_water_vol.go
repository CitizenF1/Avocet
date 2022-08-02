package odoo

import (
	"fmt"
)

// AssetWaterVol represents asset.water.vol model.
type AssetWaterVol struct {
	AssetWaterVol *Many2One `xmlrpc:"asset_water_vol,omptempty"`
	CreateDate    *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid     *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName   *String   `xmlrpc:"display_name,omptempty"`
	Id            *Int      `xmlrpc:"id,omptempty"`
	LastUpdate    *Time     `xmlrpc:"__last_update,omptempty"`
	Name          *String   `xmlrpc:"name,omptempty"`
	StartDatetime *String   `xmlrpc:"start_datetime,omptempty"`
	WaterVol      *String   `xmlrpc:"water_vol,omptempty"`
	WaterVolUom   *Many2One `xmlrpc:"water_vol_uom,omptempty"`
	WriteDate     *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid      *Many2One `xmlrpc:"write_uid,omptempty"`
}

// AssetWaterVols represents array of asset.water.vol model.
type AssetWaterVols []AssetWaterVol

// AssetWaterVolModel is the odoo model name.
const AssetWaterVolModel = "asset.water.vol"

// Many2One convert AssetWaterVol to *Many2One.
func (awv *AssetWaterVol) Many2One() *Many2One {
	return NewMany2One(awv.Id.Get(), "")
}

// CreateAssetWaterVol creates a new asset.water.vol model and returns its id.
func (c *Client) CreateAssetWaterVol(awv *AssetWaterVol) (int64, error) {
	return c.Create(AssetWaterVolModel, awv)
}

// UpdateAssetWaterVol updates an existing asset.water.vol record.
func (c *Client) UpdateAssetWaterVol(awv *AssetWaterVol) error {
	return c.UpdateAssetWaterVols([]int64{awv.Id.Get()}, awv)
}

// UpdateAssetWaterVols updates existing asset.water.vol records.
// All records (represented by ids) will be updated by awv values.
func (c *Client) UpdateAssetWaterVols(ids []int64, awv *AssetWaterVol) error {
	return c.Update(AssetWaterVolModel, ids, awv)
}

// DeleteAssetWaterVol deletes an existing asset.water.vol record.
func (c *Client) DeleteAssetWaterVol(id int64) error {
	return c.DeleteAssetWaterVols([]int64{id})
}

// DeleteAssetWaterVols deletes existing asset.water.vol records.
func (c *Client) DeleteAssetWaterVols(ids []int64) error {
	return c.Delete(AssetWaterVolModel, ids)
}

// GetAssetWaterVol gets asset.water.vol existing record.
func (c *Client) GetAssetWaterVol(id int64) (*AssetWaterVol, error) {
	awvs, err := c.GetAssetWaterVols([]int64{id})
	if err != nil {
		return nil, err
	}
	if awvs != nil && len(*awvs) > 0 {
		return &((*awvs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of asset.water.vol not found", id)
}

// GetAssetWaterVols gets asset.water.vol existing records.
func (c *Client) GetAssetWaterVols(ids []int64) (*AssetWaterVols, error) {
	awvs := &AssetWaterVols{}
	if err := c.Read(AssetWaterVolModel, ids, nil, awvs); err != nil {
		return nil, err
	}
	return awvs, nil
}

// FindAssetWaterVol finds asset.water.vol record by querying it with criteria.
func (c *Client) FindAssetWaterVol(criteria *Criteria) (*AssetWaterVol, error) {
	awvs := &AssetWaterVols{}
	if err := c.SearchRead(AssetWaterVolModel, criteria, NewOptions().Limit(1), awvs); err != nil {
		return nil, err
	}
	if awvs != nil && len(*awvs) > 0 {
		return &((*awvs)[0]), nil
	}
	return nil, fmt.Errorf("asset.water.vol was not found")
}

// FindAssetWaterVols finds asset.water.vol records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAssetWaterVols(criteria *Criteria, options *Options) (*AssetWaterVols, error) {
	awvs := &AssetWaterVols{}
	if err := c.SearchRead(AssetWaterVolModel, criteria, options, awvs); err != nil {
		return nil, err
	}
	return awvs, nil
}

// FindAssetWaterVolIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAssetWaterVolIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AssetWaterVolModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAssetWaterVolId finds record id by querying it with criteria.
func (c *Client) FindAssetWaterVolId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AssetWaterVolModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("asset.water.vol was not found")
}
