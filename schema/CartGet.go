package schema

import (
	"encoding/xml"
)

// CartGet operation enables you to retrieve the IDs, quantities,
// and prices of all of the items,
// including SavedForLater items in a remote shopping cart.
//
// Because the contents of a cart can change for different reasons, such as availability,
// you should not keep a copy of a cart locally.
// Instead, use CartGet to retrieve the items in a remote shopping cart.
//
// To retrieve the items in a cart,
// you must specify the cart using the CartId and HMAC values,
// which are returned in the CartCreate operation.
// A value similar to HMAC, URLEncodedHMAC, is also returned.
// This value is the URL encoded version of the HMAC.
// This encoding is necessary because some characters, such as + and /, cannot be included in a URL.
// Rather than encoding the HMAC yourself, use the URLEncodedHMAC value for the HMAC parameter.
//
// CartGet does not work after the customer has used the PurchaseURL to either purchase the items or
// merge them with the items in their Amazon cart.
//
// All CartGet requests must also include a value for AssociateTag.
// Otherwise, the request will fail.
//
// More details: http://docs.aws.amazon.com/AWSECommerceService/latest/DG/CartGet.html#CartGet-desc
type CartGet struct {
	XMLName           xml.Name          `xml:"CartGet"`
	MarketplaceDomain string            `xml:"MarketplaceDomain,omitempty"`
	AWSAccessKeyID    string            `xml:"AWSAccessKeyId,omitempty"`
	AssociateTag      string            `xml:"AssociateTag,omitempty"`
	Validate          string            `xml:"Validate,omitempty"`
	XMLEscaping       string            `xml:"XMLEscaping,omitempty"`
	Shared            *CartGetRequest   `xml:"Shared,omitempty"`
	Request           []*CartGetRequest `xml:"Request,omitempty"`
}

// CartGetRequest is request parameters for CartGet
//
// More details: http://docs.aws.amazon.com/AWSECommerceService/latest/DG/CartGet.html#CartGet-rp
type CartGetRequest struct {
	CartID        string   `xml:"CartId,omitempty"`
	HMAC          string   `xml:"HMAC,omitempty"`
	MergeCart     string   `xml:"MergeCart,omitempty"`
	ResponseGroup []string `xml:"ResponseGroup,omitempty"`
}

// CartGetResponse is response for CartGet
//
// More details: http://docs.aws.amazon.com/AWSECommerceService/latest/DG/CartGet.html#CartGet-resp
type CartGetResponse struct {
	XMLName          xml.Name          `xml:"CartGetResponse"`
	OperationRequest *OperationRequest `xml:"OperationRequest,omitempty"`
	Cart             []*Cart           `xml:"Cart,omitempty"`
}
