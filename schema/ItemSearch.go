package schema

import (
	"encoding/xml"
	"github.com/raitucarp/agozon/schema/Condition"
	"github.com/raitucarp/agozon/schema/audience-rating"
)

// ItemSearch is schema for ItemSearch operation.
//
// The ItemSearch operation searches for items on Amazon. The Product Advertising API returns up to ten items per search results page.

// An ItemSearch request requires a search index and the value for at least one parameter.
// For example, you might use the BrowseNode parameter for Harry Potter books and specify the Books search index.
//
// http://docs.aws.amazon.com/AWSECommerceService/latest/DG/ItemSearch.html#Description
type ItemSearch struct {
	XMLName           xml.Name             `xml:"ItemSearch"`
	MarketplaceDomain string               `xml:"MarketplaceDomain,omitempty"`
	AWSAccessKeyID    string               `xml:"AWSAccessKeyId,omitempty"`
	AssociateTag      string               `xml:"AssociateTag,omitempty"`
	XMLEscaping       string               `xml:"XMLEscaping,omitempty"`
	Validate          string               `xml:"Validate,omitempty"`
	Shared            *ItemSearchRequest   `xml:"Shared,omitempty"`
	Request           []*ItemSearchRequest `xml:"Request,omitempty"`
}

// ItemSearchRequest is request parameters for ItemSearch
//
// More details: http://docs.aws.amazon.com/AWSECommerceService/latest/DG/ItemSearch.html#ItemSearch-rp
type ItemSearchRequest struct {
	Actor                 string               `xml:"Actor,omitempty"`
	Artist                string               `xml:"Artist,omitempty"`
	Availability          string               `xml:"Availability,omitempty"`
	AudienceRating        []*audience.Rating   `xml:"AudienceRating,omitempty"`
	Author                string               `xml:"Author,omitempty"`
	Brand                 string               `xml:"Brand,omitempty"`
	BrowseNode            string               `xml:"BrowseNode,omitempty"`
	Composer              string               `xml:"Composer,omitempty"`
	Condition             *condition.Condition `xml:"Condition,omitempty"`
	Conductor             string               `xml:"Conductor,omitempty"`
	Director              string               `xml:"Director,omitempty"`
	ItemPage              *uint64              `xml:"ItemPage,omitempty"`
	Keywords              string               `xml:"Keywords,omitempty"`
	Manufacturer          string               `xml:"Manufacturer,omitempty"`
	MaximumPrice          *uint64              `xml:"MaximumPrice,omitempty"`
	MerchantID            string               `xml:"MerchantId,omitempty"`
	MinimumPrice          *uint64              `xml:"MinimumPrice,omitempty"`
	MinPercentageOff      *uint64              `xml:"MinPercentageOff,omitempty"`
	MusicLabel            string               `xml:"MusicLabel,omitempty"`
	Orchestra             string               `xml:"Orchestra,omitempty"`
	Power                 string               `xml:"Power,omitempty"`
	Publisher             string               `xml:"Publisher,omitempty"`
	RelatedItemPage       *string              `xml:"RelatedItemPage,omitempty"`
	RelationshipType      []string             `xml:"RelationshipType,omitempty"`
	ResponseGroup         []string             `xml:"ResponseGroup,omitempty"`
	SearchIndex           string               `xml:"SearchIndex,omitempty"`
	Sort                  string               `xml:"Sort,omitempty"`
	Title                 string               `xml:"Title,omitempty"`
	ReleaseDate           string               `xml:"ReleaseDate,omitempty"`
	IncludeReviewsSummary string               `xml:"IncludeReviewsSummary,omitempty"`
	TruncateReviewsAt     *uint64              `xml:"TruncateReviewsAt,omitempty"`
}

// ItemSearchResponse is response for ItemSearch
//
// More details: http://docs.aws.amazon.com/AWSECommerceService/latest/DG/ItemSearch.html#ItemSearch-resp
type ItemSearchResponse struct {
	XMLName          xml.Name          `xml:"ItemSearchResponse"`
	OperationRequest *OperationRequest `xml:"OperationRequest,omitempty"`
	Items            []*Items          `xml:"Items,omitempty"`
}
