package schema

import (
	"encoding/xml"
)

// RentalOffers schema
type RentalOffers struct {
	XMLName     xml.Name       `xml:"RentalOffers"`
	RentalOffer []*RentalOffer `xml:"RentalOffer,omitempty"`
}

// RentalOffer schema
type RentalOffer struct {
	XMLName       xml.Name         `xml:"RentalOffer"`
	Merchant      *Merchant        `xml:"Merchant,omitempty"`
	RentalListing []*RentalListing `xml:"RentalListing,omitempty"`
}
