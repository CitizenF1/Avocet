package odoo

import (
	"fmt"
)

// AssetReadingsDowntime represents asset.readings.downtime model.
type AssetReadingsDowntime struct {
	AssetId       *Many2One `xmlrpc:"asset_id,omptempty"`
	Comment       *String   `xmlrpc:"comment,omptempty"`
	CreateDate    *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid     *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName   *String   `xmlrpc:"display_name,omptempty"`
	DowntimeType  *String   `xmlrpc:"downtime_type,omptempty"`
	Duration      *String   `xmlrpc:"duration,omptempty"`
	EndDatetime   *Time     `xmlrpc:"end_datetime,omptempty"`
	Id            *Int      `xmlrpc:"id,omptempty"`
	LastUpdate    *Time     `xmlrpc:"__last_update,omptempty"`
	StartDatetime *Time     `xmlrpc:"start_datetime,omptempty"`
	WriteDate     *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid      *Many2One `xmlrpc:"write_uid,omptempty"`
}

// AssetReadingsDowntimes represents array of asset.readings.downtime model.
type AssetReadingsDowntimes []AssetReadingsDowntime

// AssetReadingsDowntimeModel is the odoo model name.
const AssetReadingsDowntimeModel = "asset.readings.downtime"

// Many2One convert AssetReadingsDowntime to *Many2One.
func (ard *AssetReadingsDowntime) Many2One() *Many2One {
	return NewMany2One(ard.Id.Get(), "")
}

// CreateAssetReadingsDowntime creates a new asset.readings.downtime model and returns its id.
func (c *Client) CreateAssetReadingsDowntime(ard *AssetReadingsDowntime) (int64, error) {
	return c.Create(AssetReadingsDowntimeModel, ard)
}

// UpdateAssetReadingsDowntime updates an existing asset.readings.downtime record.
func (c *Client) UpdateAssetReadingsDowntime(ard *AssetReadingsDowntime) error {
	return c.UpdateAssetReadingsDowntimes([]int64{ard.Id.Get()}, ard)
}

// UpdateAssetReadingsDowntimes updates existing asset.readings.downtime records.
// All records (represented by ids) will be updated by ard values.
func (c *Client) UpdateAssetReadingsDowntimes(ids []int64, ard *AssetReadingsDowntime) error {
	return c.Update(AssetReadingsDowntimeModel, ids, ard)
}

// DeleteAssetReadingsDowntime deletes an existing asset.readings.downtime record.
func (c *Client) DeleteAssetReadingsDowntime(id int64) error {
	return c.DeleteAssetReadingsDowntimes([]int64{id})
}

// DeleteAssetReadingsDowntimes deletes existing asset.readings.downtime records.
func (c *Client) DeleteAssetReadingsDowntimes(ids []int64) error {
	return c.Delete(AssetReadingsDowntimeModel, ids)
}

// GetAssetReadingsDowntime gets asset.readings.downtime existing record.
func (c *Client) GetAssetReadingsDowntime(id int64) (*AssetReadingsDowntime, error) {
	ards, err := c.GetAssetReadingsDowntimes([]int64{id})
	if err != nil {
		return nil, err
	}
	if ards != nil && len(*ards) > 0 {
		return &((*ards)[0]), nil
	}
	return nil, fmt.Errorf("id %v of asset.readings.downtime not found", id)
}

// GetAssetReadingsDowntimes gets asset.readings.downtime existing records.
func (c *Client) GetAssetReadingsDowntimes(ids []int64) (*AssetReadingsDowntimes, error) {
	ards := &AssetReadingsDowntimes{}
	if err := c.Read(AssetReadingsDowntimeModel, ids, nil, ards); err != nil {
		return nil, err
	}
	return ards, nil
}

// FindAssetReadingsDowntime finds asset.readings.downtime record by querying it with criteria.
func (c *Client) FindAssetReadingsDowntime(criteria *Criteria) (*AssetReadingsDowntime, error) {
	ards := &AssetReadingsDowntimes{}
	if err := c.SearchRead(AssetReadingsDowntimeModel, criteria, NewOptions().Limit(1), ards); err != nil {
		return nil, err
	}
	if ards != nil && len(*ards) > 0 {
		return &((*ards)[0]), nil
	}
	return nil, fmt.Errorf("asset.readings.downtime was not found")
}

// FindAssetReadingsDowntimes finds asset.readings.downtime records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAssetReadingsDowntimes(criteria *Criteria, options *Options) (*AssetReadingsDowntimes, error) {
	ards := &AssetReadingsDowntimes{}
	if err := c.SearchRead(AssetReadingsDowntimeModel, criteria, options, ards); err != nil {
		return nil, err
	}
	return ards, nil
}

// FindAssetReadingsDowntimeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAssetReadingsDowntimeIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AssetReadingsDowntimeModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAssetReadingsDowntimeId finds record id by querying it with criteria.
func (c *Client) FindAssetReadingsDowntimeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AssetReadingsDowntimeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("asset.readings.downtime was not found")
}
