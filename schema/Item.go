package schema

import (
	"encoding/xml"
)

// Item is Container for item information, including ASIN, DetailPageURL, ItemLinks,  ItemAttributes, Title, ProductGroup, and Manufacturer.
type Item struct {
	XMLName             xml.Name        `xml:"Item"`
	ASIN                string          `xml:"ASIN,omitempty"`
	ParentASIN          string          `xml:"ParentASIN,omitempty"`
	Errors              *Errors         `xml:"Errors,omitempty"`
	DetailPageURL       string          `xml:"DetailPageURL,omitempty"`
	ItemLinks           *ItemLinks      `xml:"ItemLinks,omitempty"`
	SalesRank           string          `xml:"SalesRank,omitempty"`
	SmallImage          *Image          `xml:"SmallImage,omitempty"`
	MediumImage         *Image          `xml:"MediumImage,omitempty"`
	LargeImage          *Image          `xml:"LargeImage,omitempty"`
	ImageSets           []*ImageSets    `xml:"ImageSets,omitempty"`
	Attributes          *ItemAttributes `xml:"ItemAttributes,omitempty"`
	VariationAttributes struct {
		VariationAttribute []*VariationAttribute `xml:"VariationAttribute,omitempty"`
	} `xml:"VariationAttributes,omitempty"`
	RelatedItems []*RelatedItems `xml:"RelatedItems,omitempty"`
	Collections  *Collections    `xml:"Collections,omitempty"`
	Subjects     struct {
		Subject []string `xml:"Subject,omitempty"`
	} `xml:"Subjects,omitempty"`
	OfferSummary      *OfferSummary      `xml:"OfferSummary,omitempty"`
	Offers            *Offers            `xml:"Offers,omitempty"`
	RentalOffers      *RentalOffers      `xml:"RentalOffers,omitempty"`
	VariationSummary  *VariationSummary  `xml:"VariationSummary,omitempty"`
	Variations        *Variations        `xml:"Variations,omitempty"`
	CustomerReviews   *CustomerReviews   `xml:"CustomerReviews,omitempty"`
	EditorialReviews  *EditorialReviews  `xml:"EditorialReviews,omitempty"`
	SimilarProducts   *SimilarProducts   `xml:"SimilarProducts,omitempty"`
	Accessories       *Accessories       `xml:"Accessories,omitempty"`
	Tracks            *Tracks            `xml:"Tracks,omitempty"`
	BrowseNodes       *BrowseNodes       `xml:"BrowseNodes,omitempty"`
	AlternateVersions *AlternateVersions `xml:"AlternateVersions>AlternateVersion,omitempty"`
}

// AlternateVersions is alternate version
type AlternateVersions struct {
	ASIN    string `xml:"ASIN,omitempty"`
	Title   string `xml:"Title,omitempty"`
	Binding string `xml:"Binding,omitempty"`
}
