package schema

import (
	"encoding/xml"
)

// SearchBinSet is a means by which to categorize results, such as price range.
//
// Ancestry: SearchBinSets
//
// More details: http://docs.aws.amazon.com/AWSECommerceService/latest/DG/CHAP_response_elements.html
type SearchBinSet struct {
	XMLName  xml.Name `xml:"SearchBinSet"`
	Bin      []*Bin   `xml:"Bin,omitempty"`
	NarrowBy string   `xml:"NarrowBy,attr,omitempty"`
}

// SearchBinSets is Parent element for SearchBins element.
//
// Ancestry: SearchBinSets
//
// More details: http://docs.aws.amazon.com/AWSECommerceService/latest/DG/CHAP_response_elements.html
type SearchBinSets struct {
	XMLName      xml.Name        `xml:"SearchBinSets"`
	SearchBinSet []*SearchBinSet `xml:"SearchBinSet,omitempty"`
}
