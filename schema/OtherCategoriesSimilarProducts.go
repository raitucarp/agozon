package schema

import (
	"encoding/xml"
)

// OtherCategoriesSimilarProduct is single similar product item for OtherCategoriesSimilarProducts
//
// Ancestry: OtherCategoriesSimilarProducts
type OtherCategoriesSimilarProduct struct {
	ASIN  string `xml:"ASIN,omitempty"`
	Title string `xml:"Title,omitempty"`
}

// OtherCategoriesSimilarProducts is parent element for Title and ASIN of similar products in other product groups
//
// Ancestry: Cart
//
// More details: http://docs.aws.amazon.com/AWSECommerceService/latest/DG/CHAP_response_elements.html
type OtherCategoriesSimilarProducts struct {
	XMLName                       xml.Name                         `xml:"OtherCategoriesSimilarProducts"`
	OtherCategoriesSimilarProduct []*OtherCategoriesSimilarProduct `xml:"OtherCategoriesSimilarProduct,omitempty"`
}
