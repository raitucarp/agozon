package schema

import (
	"encoding/xml"
)

// StringWithUnits is string with units that describe it.
type StringWithUnits struct {
	XMLName xml.Name `xml:"StringWithUnits"`
	Value   string   `xml:",chardata"`
	Units   string   `xml:"Units,attr,omitempty"`
}
