package condition

import (
	"encoding/xml"
	"reflect"
)

// valid values of condition
var conditions = []string{"", "All", "New", "Used", "Collectible", "Refurbished"}

// Condition specifies the condition of the item, such as new, used, collectible, or refurbished.
//
// Ancestry: SellerListing/Condition Offers/Offer/OfferAttributes/Condition
type Condition int

const (
	All Condition = 1 + iota
	New
	Used
	Collectible
	Refurbished
)

// MarshalXML marshal condition into xml with custom xml structure
func (c Condition) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: start.Name})
	e.EncodeToken(xml.CharData(c.String()))
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

// UnmarshalXML unmarshal condition type to be exactly good structure.
func (c *Condition) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// create condition
	var condition string

	// decode to condition
	err := d.DecodeElement(&condition, &start)
	if err != nil {
		return err
	}

	// create index
	var index int

	// find index in conditions
	for k, v := range conditions {
		if v == condition {
			index = k
		}
	}

	// set condition
	*c = Condition(index)
	return nil
}

// String is a stringer for Condition type
func (c Condition) String() (s string) {
	index := reflect.ValueOf(c).Int()
	l := len(conditions)
	if index >= int64(l) {
		return "New"
	}
	return conditions[index]
}
