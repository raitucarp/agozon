package schema

import (
	"encoding/xml"
	"github.com/raitucarp/agozon/schema/Condition"
)

// SimilarityLookup schema for SimilarityLookup operation
//
// Description
//
// The SimilarityLookup operation returns up to ten products per page that are similar to one or more items specified in the request.
// This operation is typically used to pique a customer’s interest in buying something similar to what they’ve already ordered.
//
// If you specify more than one item, SimilarityLookup returns the intersection of similar items each item would return separately.
// Alternatively, you can use the SimilarityType parameter to return the union of items that are similar to any of the specified items.
// A maximum of ten similar items are returned; the operation does not return additional pages of similar items.
// If there are more than ten similar items,
// running the same request can result in different answers because the ten that are included in the response are picked randomly.
// The results are picked randomly only when you specify multiple items and the results include more than ten similar items.
// When you specify multiple items, it is possible for there to be no intersection of similar items. In this case,
// the operation returns the following error:
// 	<Error>
// 		<Code>AWS.ECommerceService.NoSimilarities</Code>
// 		<Message>There are no similar items for this ASIN: B00006WREH.</Message>
// 	</Error>
//
// This result is very often the case if the items belong to different search indices. The error can occur, however, even when the items share the same search index.
//
// Similarity is a measurement of similar items purchased, that is, customers who bought X also bought Y and Z. It is not a measure, for example, of items viewed, that is, customers who viewed X also viewed Y and Z.
// More details:
// 	- http://docs.aws.amazon.com/AWSECommerceService/latest/DG/SimilarityLookup.html#SimilarityLookup-desc
// 	- http://docs.aws.amazon.com/AWSECommerceService/latest/DG/UsingSimilarityLookup.html
type SimilarityLookup struct {
	XMLName           xml.Name                   `xml:"SimilarityLookup"`
	MarketplaceDomain string                     `xml:"MarketplaceDomain,omitempty"`
	AWSAccessKeyID    string                     `xml:"AWSAccessKeyId,omitempty"`
	AssociateTag      string                     `xml:"AssociateTag,omitempty"`
	Validate          string                     `xml:"Validate,omitempty"`
	XMLEscaping       string                     `xml:"XMLEscaping,omitempty"`
	Shared            *SimilarityLookupRequest   `xml:"Shared,omitempty"`
	Request           []*SimilarityLookupRequest `xml:"Request,omitempty"`
}

// SimilarityLookupRequest is request parameters for SimilarityLookup
//
// More details: http://docs.aws.amazon.com/AWSECommerceService/latest/DG/SimilarityLookup.html
type SimilarityLookupRequest struct {
	Condition      *condition.Condition `xml:"Condition,omitempty"`
	ItemID         []string             `xml:"ItemId,omitempty"`
	MerchantID     string               `xml:"MerchantId,omitempty"`
	ResponseGroup  []string             `xml:"ResponseGroup,omitempty"`
	SimilarityType string               `xml:"SimilarityType,omitempty"`
}

// SimilarityLookupResponse is response for SimilarityLookup
//
// More details: http://docs.aws.amazon.com/AWSECommerceService/latest/DG/SimilarityLookup.html
type SimilarityLookupResponse struct {
	XMLName          xml.Name          `xml:"SimilarityLookupResponse"`
	OperationRequest *OperationRequest `xml:"OperationRequest,omitempty"`
	Items            []*Items          `xml:"Items,omitempty"`
}
