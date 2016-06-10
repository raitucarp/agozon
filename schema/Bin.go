package schema

import (
	"encoding/xml"
)

// BinParameter is a container for the BrowseNodeId and it's value.
//
// Ancestry: Bin
//
// Read more: http://docs.aws.amazon.com/AWSECommerceService/latest/DG/CHAP_response_elements.html
type BinParameter struct {
	Name  string `xml:"Name,omitempty"`
	Value string `xml:"Value,omitempty"`
}

// Bin is a container for Bin elements.
//
// Children: BinItemCount, BinName, BinParameter
type Bin struct {
	XMLName   xml.Name       `xml:"Bin"`
	Name      string         `xml:"BinName,omitempty"`
	ItemCount *uint64        `xml:"BinItemCount,omitempty"`
	Parameter []BinParameter `xml:"BinParameter,omitempty"`
}
