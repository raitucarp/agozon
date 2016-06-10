package schema

import (
	"encoding/xml"
)

// RentalListing schema, no description available from Amazon.
type RentalListing struct {
	XMLName             xml.Name                     `xml:"RentalListing"`
	Price               *Price                       `xml:"Price,omitempty"`
	Period              *NonNegativeIntegerWithUnits `xml:"Period,omitempty"`
	IsFulfilledByAmazon bool                         `xml:"IsFulfilledByAmazon,omitempty"`
	Disclaimer          string                       `xml:"Disclaimer,omitempty"`
}
