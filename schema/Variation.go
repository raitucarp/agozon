package schema

import (
	"encoding/xml"
)

// VariationAttribute is a container for a variation name and value.
//
// Ancestry: Item/VariationAttributes/
//
// More details: http://docs.aws.amazon.com/AWSECommerceService/latest/DG/CHAP_response_elements.html#VariationAttribute
type VariationAttribute struct {
	XMLName xml.Name `xml:"VariationAttribute"`
	Name    string   `xml:"Name,omitempty"`
	Value   []string `xml:"Value,omitempty"`
}

// VariationDimensions is container for dimensions
//
// Ancestry: Variations/VariationDimensions/
//
// A variation is a child ASIN. The parent ASIN is an abstraction of the children items. For example, Shirt is a parent ASIN. Parent ASINs cannot be sold. A child ASIN of it would be a blue shirt, size 16, sold by MyApparelStore. This child ASIN is one of potentially many variations. The ways in which variations differ are called dimensions.
//
// In the preceding example, size and color are the dimensions. Parent ASINs therefore return two related elements:
// 	- VariationDimensions
// 	- VariationDimension
//
// More details: http://docs.aws.amazon.com/AWSECommerceService/latest/DG/CHAP_response_elements.html#VariationDimensions
type VariationDimensions struct {
	XMLName            xml.Name `xml:"VariationDimensions"`
	VariationDimension []string `xml:"VariationDimension,omitempty"`
}

// Variations schema.
//
// Often, an item comes in a variety of sizes and colors. A shirt, for example, might come in four different sizes and six different colors.
//
// Each color and size combination is called a variation. Each variation, such as a medium, blue shirt, is an item that a customer can buy. For that reason, each variation has its own ASIN. For example, if a shirt came in four sizes and six colors, there would be 24 variations, each with a unique ASIN.
//
// More details: http://docs.aws.amazon.com/AWSECommerceService/latest/DG/Variations.html
type Variations struct {
	XMLName             xml.Name             `xml:"Variations"`
	TotalVariations     *uint64              `xml:"TotalVariations,omitempty"`
	TotalVariationPages *uint64              `xml:"TotalVariationPages,omitempty"`
	VariationDimensions *VariationDimensions `xml:"VariationDimensions,omitempty"`
	Item                []*Item              `xml:"Item,omitempty"`
}

// VariationSummary response group provides the lowest price, highest price, lowest sale price, and highest sale price for all child ASINs in a response.
//
// Parent ASINs do not have offers, but their children do.
// For example, you cannot buy a shirt (the parent ASIN) but you can buy a shirt that is a specific color and size (the child ASIN).
//
// More details: http://docs.aws.amazon.com/AWSECommerceService/latest/DG/RG_VariationSummary.html
type VariationSummary struct {
	XMLName          xml.Name `xml:"VariationSummary"`
	LowestPrice      *Price   `xml:"LowestPrice,omitempty"`
	HighestPrice     *Price   `xml:"HighestPrice,omitempty"`
	LowestSalePrice  *Price   `xml:"LowestSalePrice,omitempty"`
	HighestSalePrice *Price   `xml:"HighestSalePrice,omitempty"`
}
