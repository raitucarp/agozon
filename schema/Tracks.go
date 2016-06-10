package schema

import (
	"encoding/xml"
)

// Track refers to each track on a CD. On a music CD, each track corresponds to a song.
//
// Ancestry: Tracks/Disc
//
// More details: http://docs.aws.amazon.com/AWSECommerceService/latest/DG/CHAP_response_elements.html#Track
type Track struct {
	Value  string  `xml:",chardata"`
	Number *uint64 `xml:"Number,attr,omitempty"`
}

// Disc is element of Tracks type.
type Disc struct {
	Track  []*Track `xml:"Track,omitempty"`
	Number *uint64  `xml:"Number,attr,omitempty"`
}

// Tracks response group returns the title and number of
// each track on each CD in the response. For example,
// you could use ItemLookup to return Tracks information about a specified CD.
//
// More details: http://docs.aws.amazon.com/AWSECommerceService/latest/DG/RG_Tracks.html
type Tracks struct {
	XMLName xml.Name `xml:"Tracks"`
	Disc    []*Disc  `xml:"Disc,omitempty"`
}
