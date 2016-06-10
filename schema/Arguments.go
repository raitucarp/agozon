package schema

import (
	"encoding/xml"
)

// Argument for request
type Argument struct {
	Name  string `xml:"Name,attr,omitempty"`
	Value string `xml:"Value,attr,omitempty"`
}

// Arguments provides one or argument of request
type Arguments struct {
	XMLName  xml.Name   `xml:"Arguments"`
	Argument []Argument `xml:"Argument,omitempty"`
}
