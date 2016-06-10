package schema

import (
	"encoding/xml"
)

// SimilarViewedProduct is children of SimilarViewedProducts schema.
type SimilarViewedProduct struct {
	ASIN  string `xml:"ASIN,omitempty"`
	Title string `xml:"Title,omitempty"`
}

// SimilarViewedProducts is parent element for Title and ASIN of similar products in the same product group that have been viewed
//
// Ancestry: Cart
//
// More details: http://docs.aws.amazon.com/AWSECommerceService/latest/DG/CHAP_response_elements.html
type SimilarViewedProducts struct {
	XMLName              xml.Name                `xml:"SimilarViewedProducts"`
	SimilarViewedProduct []*SimilarViewedProduct `xml:"SimilarViewedProduct,omitempty"`
}
