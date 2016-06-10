package schema

import (
	"encoding/xml"
)

// CorrectedQuery is a parent element that contains the elements related to a corrected keyword. See Keywords.
//
// Ancestry: Items
type CorrectedQuery struct {
	XMLName  xml.Name `xml:"CorrectedQuery"`
	Keywords string   `xml:"Keywords,omitempty"`
	Message  string   `xml:"Message,omitempty"`
}
