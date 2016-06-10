package schema

import (
	"encoding/xml"
)

// Merchant is schema for merchant.
type Merchant struct {
	XMLName xml.Name `xml:"Merchant"`
	Name    string   `xml:"Name,omitempty"`
}
