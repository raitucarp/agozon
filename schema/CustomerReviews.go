package schema

import (
	"encoding/xml"
)

// CustomerReviews provides customer reviews type.
//
// More details: http://docs.aws.amazon.com/AWSECommerceService/latest/DG/CustomerandSellerReviews.html#GettingCustomerReviews
type CustomerReviews struct {
	XMLName    xml.Name `xml:"CustomerReviews"`
	IFrameURL  string   `xml:"IFrameURL,omitempty"`
	HasReviews bool     `xml:"HasReviews,omitempty"`
}
