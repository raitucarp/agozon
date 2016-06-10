package schema

import (
	"encoding/xml"
)

// Offer schema
type Offer struct {
	XMLName         xml.Name         `xml:"Offer"`
	Merchant        *Merchant        `xml:"Merchant,omitempty"`
	OfferAttributes *OfferAttributes `xml:"OfferAttributes,omitempty"`
	OfferListing    []*OfferListing  `xml:"OfferListing,omitempty"`
	LoyaltyPoints   *LoyaltyPoints   `xml:"LoyaltyPoints,omitempty"`
	Promotions      *Promotions      `xml:"Promotions,omitempty"`
}

// OfferAttributes is attributes for Offer schema.
type OfferAttributes struct {
	XMLName   xml.Name `xml:"OfferAttributes"`
	Condition string   `xml:"Condition,omitempty"`
}

// OfferListing is part of OfferListings response group.
// It returns the OfferListings for items returned in the response.
// The values returned are similar to those returned by the Offers
// response group minus the values returned by the OfferSummary response group.
// OfferListings returns shipping options, including IsEligibleForSuperSavingShipping which specifies if the item qualifies for super saver shipping.
//
//	Note
//	This response group is not returned for Amazon Kindle digital books.
//	An Amazon Kindle ASIN can be verified through the Binding, Format, and ProductTypeName response elements.
//
// More details: http://docs.aws.amazon.com/AWSECommerceService/latest/DG/RG_OfferListings.html
type OfferListing struct {
	XMLName                xml.Name `xml:"OfferListing"`
	OfferListingID         string   `xml:"OfferListingId,omitempty"`
	PricePerUnit           string   `xml:"PricePerUnit,omitempty"`
	Price                  *Price   `xml:"Price,omitempty"`
	SalePrice              *Price   `xml:"SalePrice,omitempty"`
	AmountSaved            *Price   `xml:"AmountSaved,omitempty"`
	PercentageSaved        *uint64  `xml:"PercentageSaved,omitempty"`
	Availability           string   `xml:"Availability,omitempty"`
	AvailabilityAttributes struct {
		AvailabilityType string `xml:"AvailabilityType,omitempty"`
		IsPreorder       bool   `xml:"IsPreorder,omitempty"`
		MinimumHours     int32  `xml:"MinimumHours,omitempty"`
		MaximumHours     int32  `xml:"MaximumHours,omitempty"`
	} `xml:"AvailabilityAttributes,omitempty"`
	IsEligibleForSuperSaverShipping    bool `xml:"IsEligibleForSuperSaverShipping,omitempty"`
	IsEligibleForPrimeFreeDigitalVideo bool `xml:"IsEligibleForPrimeFreeDigitalVideo,omitempty"`
	IsEligibleForPrime                 bool `xml:"IsEligibleForPrime,omitempty"`
}

// Offers is parent element of Offer
type Offers struct {
	XMLName         xml.Name `xml:"Offers"`
	TotalOffers     *uint64  `xml:"TotalOffers,omitempty"`
	TotalOfferPages *uint64  `xml:"TotalOfferPages,omitempty"`
	MoreOffersURL   string   `xml:"MoreOffersUrl,omitempty"`
	Offer           []*Offer `xml:"Offer,omitempty"`
}

// OfferSummary response group returns the number of offer listings and the lowest price for each condition type for each item in the response.
// Condition types are New, Used, Collectible, and Refurbished. For example, this response group returns the lowest price for each Condition:
//
// 	- New item
// 	- Used item
// 	- Collectible item
// 	- Refurbished item
//
// Individual offer listings are not returned. The OfferSummary is dependent only on the ASIN parameter and is not affected by the MerchantId or Condition parameters (i.e. the OfferSummary will always be the same for a given ASIN independent of other parameters).
//
// 	Note
// 	This response group is not returned for Amazon Kindle digital books.
// 	An Amazon Kindle ASIN can be verified through the Binding, Format, and ProductTypeName response elements.
//
// More details: http://docs.aws.amazon.com/AWSECommerceService/latest/DG/RG_OfferSummary.html
type OfferSummary struct {
	XMLName                xml.Name `xml:"OfferSummary"`
	LowestNewPrice         *Price   `xml:"LowestNewPrice,omitempty"`
	LowestUsedPrice        *Price   `xml:"LowestUsedPrice,omitempty"`
	LowestCollectiblePrice *Price   `xml:"LowestCollectiblePrice,omitempty"`
	LowestRefurbishedPrice *Price   `xml:"LowestRefurbishedPrice,omitempty"`
	TotalNew               string   `xml:"TotalNew,omitempty"`
	TotalUsed              string   `xml:"TotalUsed,omitempty"`
	TotalCollectible       string   `xml:"TotalCollectible,omitempty"`
	TotalRefurbished       string   `xml:"TotalRefurbished,omitempty"`
}
