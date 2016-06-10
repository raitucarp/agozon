package schema

import (
	"encoding/xml"
)

// Property schema.
type Property struct {
	XMLName xml.Name `xml:"Property"`
	Name    string   `xml:"Name,omitempty"`
	Value   string   `xml:"Value,omitempty"`
}
