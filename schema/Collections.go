package schema

import (
	"encoding/xml"
)

// Collection is container for items that are part of a collection.
//
// Ancestry: None
type Collection struct {
	Summary struct {
		LowestListPrice  *Price `xml:"LowestListPrice,omitempty"`
		HighestListPrice *Price `xml:"HighestListPrice,omitempty"`
		LowestSalePrice  *Price `xml:"LowestSalePrice,omitempty"`
		HighestSalePrice *Price `xml:"HighestSalePrice,omitempty"`
	} `xml:"CollectionSummary,omitempty"`
	Parent struct {
		ASIN  string `xml:"ASIN,omitempty"`
		Title string `xml:"Title,omitempty"`
	} `xml:"CollectionParent,omitempty"`
	Item []*CollectionItem `xml:"CollectionItem,omitempty"`
}

// CollectionItem is collection of items
type CollectionItem struct {
	ASIN  string `xml:"ASIN,omitempty"`
	Title string `xml:"Title,omitempty"`
}

// Collections are bulk of Collection
type Collections struct {
	XMLName    xml.Name      `xml:"Collections"`
	Collection []*Collection `xml:"Collection,omitempty"`
}
