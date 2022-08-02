package odoo

import (
	"fmt"
)

// WellWaterLines represents well.water.lines model.
type WellWaterLines struct {
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Fact        *Float    `xmlrpc:"fact,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	Month       *Time     `xmlrpc:"month,omptempty"`
	Plan        *Float    `xmlrpc:"plan,omptempty"`
	ProductId   *Many2One `xmlrpc:"product_id,omptempty"`
	RequestId   *Many2One `xmlrpc:"request_id,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// WellWaterLiness represents array of well.water.lines model.
type WellWaterLiness []WellWaterLines

// WellWaterLinesModel is the odoo model name.
const WellWaterLinesModel = "well.water.lines"

// Many2One convert WellWaterLines to *Many2One.
func (wwl *WellWaterLines) Many2One() *Many2One {
	return NewMany2One(wwl.Id.Get(), "")
}

// CreateWellWaterLines creates a new well.water.lines model and returns its id.
func (c *Client) CreateWellWaterLines(wwl *WellWaterLines) (int64, error) {
	return c.Create(WellWaterLinesModel, wwl)
}

// UpdateWellWaterLines updates an existing well.water.lines record.
func (c *Client) UpdateWellWaterLines(wwl *WellWaterLines) error {
	return c.UpdateWellWaterLiness([]int64{wwl.Id.Get()}, wwl)
}

// UpdateWellWaterLiness updates existing well.water.lines records.
// All records (represented by ids) will be updated by wwl values.
func (c *Client) UpdateWellWaterLiness(ids []int64, wwl *WellWaterLines) error {
	return c.Update(WellWaterLinesModel, ids, wwl)
}

// DeleteWellWaterLines deletes an existing well.water.lines record.
func (c *Client) DeleteWellWaterLines(id int64) error {
	return c.DeleteWellWaterLiness([]int64{id})
}

// DeleteWellWaterLiness deletes existing well.water.lines records.
func (c *Client) DeleteWellWaterLiness(ids []int64) error {
	return c.Delete(WellWaterLinesModel, ids)
}

// GetWellWaterLines gets well.water.lines existing record.
func (c *Client) GetWellWaterLines(id int64) (*WellWaterLines, error) {
	wwls, err := c.GetWellWaterLiness([]int64{id})
	if err != nil {
		return nil, err
	}
	if wwls != nil && len(*wwls) > 0 {
		return &((*wwls)[0]), nil
	}
	return nil, fmt.Errorf("id %v of well.water.lines not found", id)
}

// GetWellWaterLiness gets well.water.lines existing records.
func (c *Client) GetWellWaterLiness(ids []int64) (*WellWaterLiness, error) {
	wwls := &WellWaterLiness{}
	if err := c.Read(WellWaterLinesModel, ids, nil, wwls); err != nil {
		return nil, err
	}
	return wwls, nil
}

// FindWellWaterLines finds well.water.lines record by querying it with criteria.
func (c *Client) FindWellWaterLines(criteria *Criteria) (*WellWaterLines, error) {
	wwls := &WellWaterLiness{}
	if err := c.SearchRead(WellWaterLinesModel, criteria, NewOptions().Limit(1), wwls); err != nil {
		return nil, err
	}
	if wwls != nil && len(*wwls) > 0 {
		return &((*wwls)[0]), nil
	}
	return nil, fmt.Errorf("well.water.lines was not found")
}

// FindWellWaterLiness finds well.water.lines records by querying it
// and filtering it with criteria and options.
func (c *Client) FindWellWaterLiness(criteria *Criteria, options *Options) (*WellWaterLiness, error) {
	wwls := &WellWaterLiness{}
	if err := c.SearchRead(WellWaterLinesModel, criteria, options, wwls); err != nil {
		return nil, err
	}
	return wwls, nil
}

// FindWellWaterLinesIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindWellWaterLinesIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(WellWaterLinesModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindWellWaterLinesId finds record id by querying it with criteria.
func (c *Client) FindWellWaterLinesId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(WellWaterLinesModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("well.water.lines was not found")
}
