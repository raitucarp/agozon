package schema

import (
	"encoding/xml"
)

// TopSeller is element of TopSellers type.
type TopSeller struct {
	ASIN  string `xml:"ASIN,omitempty"`
	Title string `xml:"Title,omitempty"`
}

// TopSellers response group returns the ASINs and titles of the 10 best sellers within a specified browse node.
//
// More details: http://docs.aws.amazon.com/AWSECommerceService/latest/DG/RG_TopSellers.html
type TopSellers struct {
	XMLName   xml.Name     `xml:"TopSellers"`
	TopSeller []*TopSeller `xml:"TopSeller,omitempty"`
}
