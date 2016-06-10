package schema

import (
	"encoding/xml"
)

// PromotionSummary response group returns summary information about a promotion,
// including the type of promotion, the promotion ID,
// eligibility requirements, and text that describes the specifics of the promotion.
//
// PromotionSummary must be used with one of the following response groups:
// 	- Large
// 	- OfferFull
// 	- Offers
//
// An error is returned if ProductDetails is not accompanied by one of these response groups.
//
// PromotionSummary is summary for Promotion.
// PromotionSummary returns the promotion information for an item.
//
// More details: http://docs.aws.amazon.com/AWSECommerceService/latest/DG/CHAP_MotivatingCustomerstoBuy_Promotions.html
type PromotionSummary struct {
	PromotionID                       string `xml:"PromotionId,omitempty"`
	Message                           string `xml:"Message,omitempty"`
	Category                          string `xml:"Category,omitempty"`
	StartDate                         string `xml:"StartDate,omitempty"`
	EndDate                           string `xml:"EndDate,omitempty"`
	EligibilityRequirementDescription string `xml:"EligibilityRequirementDescription,omitempty"`
	BenefitDescription                string `xml:"BenefitDescription,omitempty"`
	TermsAndConditions                string `xml:"TermsAndConditions,omitempty"`
}

// Promotion is a container for one or more Details elements.
//
// Ancestry: Offers/OfferPromotions
//
// More details: http://docs.aws.amazon.com/AWSECommerceService/latest/DG/CHAP_response_elements.html
type Promotion struct {
	XMLName xml.Name         `xml:"Promotion"`
	Summary PromotionSummary `xml:"Summary,omitempty"`
}

// Promotions is a container for one or more Promotion elements.
//
// Ancestry: Offers/Offer
//
// More details: http://docs.aws.amazon.com/AWSECommerceService/latest/DG/CHAP_MotivatingCustomerstoBuy_Promotions.html
type Promotions struct {
	XMLName   xml.Name     `xml:"Promotions"`
	Promotion []*Promotion `xml:"Promotion,omitempty"`
}
