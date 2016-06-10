package schema

import (
	"encoding/xml"
)

// ItemLink is link of an item
type ItemLink struct {
	XMLName     xml.Name `xml:"ItemLink"`
	Description string   `xml:"Description,omitempty"`
	URL         string   `xml:"URL,omitempty"`
}

// ItemLinks is collection of ItemLink
type ItemLinks struct {
	XMLName  xml.Name    `xml:"ItemLinks"`
	ItemLink []*ItemLink `xml:"ItemLink,omitempty"`
}
