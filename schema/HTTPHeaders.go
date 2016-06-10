package schema

import (
	"encoding/xml"
)

// Header is the header key value
type Header struct {
	Name  string `xml:"Name,attr,omitempty"`
	Value string `xml:"Value,attr,omitempty"`
}

// HTTPHeaders is an http headers requried for request.
type HTTPHeaders struct {
	XMLName xml.Name `xml:"HTTPHeaders"`
	Header  []Header `xml:"Header,omitempty"`
}
