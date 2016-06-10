package schema

import (
	"encoding/xml"
	"github.com/raitucarp/agozon/schema/Condition"
	"github.com/raitucarp/agozon/schema/id"
)

// ItemLookup is schema for operation ItemLookup.
//
// Given an Item identifier, the ItemLookup operation returns some or all of the item attributes,
// depending on the response group specified in the request.
// By default, ItemLookup returns an item’s ASIN, Manufacturer, ProductGroup, and Title of the item.
//
// ItemLookup supports many response groups.
// Response groups return product information, called item attributes.
// Item attributes include product reviews, variations, similar products, pricing, availability,
// images of products, accessories, and other information.
//
// To look up more than one item at a time, separate the item identifiers by commas.
//
// More details: http://docs.aws.amazon.com/AWSECommerceService/latest/DG/ItemLookup.html#ItemLookup-desc
type ItemLookup struct {
	XMLName           xml.Name             `xml:"ItemLookup"`
	MarketplaceDomain string               `xml:"MarketplaceDomain,omitempty"`
	AWSAccessKeyID    string               `xml:"AWSAccessKeyId,omitempty"`
	AssociateTag      string               `xml:"AssociateTag,omitempty"`
	Validate          string               `xml:"Validate,omitempty"`
	XMLEscaping       string               `xml:"XMLEscaping,omitempty"`
	Shared            *ItemLookupRequest   `xml:"Shared,omitempty"`
	Request           []*ItemLookupRequest `xml:"Request,omitempty"`
}

// ItemLookupRequest is request parameters for ItemLookup
//
// Read more: http://docs.aws.amazon.com/AWSECommerceService/latest/DG/ItemLookup.html#ItemLookup-rp
type ItemLookupRequest struct {
	Condition *condition.Condition `xml:"Condition,omitempty"`
	// IDType is type of ID whether is ASIN, UPC, SKU, EAN, or ISBN
	IDType                *id.Type `xml:"IdType,omitempty"`
	MerchantID            string   `xml:"MerchantId,omitempty"`
	ItemID                []string `xml:"ItemId,omitempty"`
	ResponseGroup         []string `xml:"ResponseGroup,omitempty"`
	SearchIndex           string   `xml:"SearchIndex,omitempty"`
	VariationPage         *string  `xml:"VariationPage,omitempty"`
	RelatedItemPage       *string  `xml:"RelatedItemPage,omitempty"`
	RelationshipType      []string `xml:"RelationshipType,omitempty"`
	IncludeReviewsSummary string   `xml:"IncludeReviewsSummary,omitempty"`
	TruncateReviewsAt     *uint64  `xml:"TruncateReviewsAt,omitempty"`
}

// ItemLookupResponse is a response for ItemLookup
//
// Read more: http://docs.aws.amazon.com/AWSECommerceService/latest/DG/ItemLookup.html#ItemLookup-resp
type ItemLookupResponse struct {
	XMLName          xml.Name          `xml:"ItemLookupResponse"`
	OperationRequest *OperationRequest `xml:"OperationRequest,omitempty"`
	Items            []*Items          `xml:"Items,omitempty"`
}
