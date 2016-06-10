package schema

import (
	"encoding/xml"
)

// LoyaltyPoints is the number of points awarded as part of a purchase. Points translate into rewards.
// In the JP locale only, loyalty points are returned. Loyalty points are used to encourage patronage and stimulate sales.
//
// Ancestry: Offers, VariationOffers
//
// More details:
// http://docs.aws.amazon.com/AWSECommerceService/latest/DG/CHAP_response_elements.html
// http://docs.aws.amazon.com/AWSECommerceService/latest/DG/RG_Offers.html#RG_Offers_LoyaltyPoints
type LoyaltyPoints struct {
	XMLName                xml.Name `xml:"LoyaltyPoints"`
	Points                 uint64   `xml:"Points"`
	TypicalRedemptionValue *Price   `xml:"TypicalRedemptionValue,omitempty"`
}
