package schema

import (
	"encoding/xml"
)

// SimilarProduct is children of SimilarProduct
type SimilarProduct struct {
	ASIN  string `xml:"ASIN,omitempty"`
	Title string `xml:"Title,omitempty"`
}

// SimilarProducts is parent element for Title and ASIN of similar products in the same product group
//
// Ancestry: Cart
//
// More details: http://docs.aws.amazon.com/AWSECommerceService/latest/DG/CHAP_response_elements.html
type SimilarProducts struct {
	XMLName        xml.Name          `xml:"SimilarProducts"`
	SimilarProduct []*SimilarProduct `xml:"SimilarProduct,omitempty"`
}
