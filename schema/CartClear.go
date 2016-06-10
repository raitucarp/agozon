package schema

import (
	"encoding/xml"
)

// CartClear operation enables you to remove all of the items in a remote shopping cart, including SavedForLater items.
// To remove only some of the items in a cart or to reduce the quantity of one or more items, use CartModify.
//
// To delete all of the items from a remote shopping cart, you must specify the cart using the CartId and HMAC values,
// which are returned by the CartCreate operation. A value similar to the HMAC, URLEncodedHMAC, is also returned.
// This value is the URL encoded version of the HMAC.
// This encoding is necessary because some characters, such as + and /, cannot be included in a URL.
// Rather than encoding the HMAC yourself, use the URLEncodedHMAC value for the HMAC parameter.
// CartClear does not work after the customer has used the PurchaseURL to either purchase the items or merge them with the items in their Amazon cart.
//
// Carts exist even though they have been emptied. The lifespan of a cart is 7 days since the last time it was acted upon.
// For example, if a cart created 6 days ago is modified, the cart lifespan is reset to 7 days.
//
// read more: http://docs.aws.amazon.com/AWSECommerceService/latest/DG/CartClear.html#CartClear-desc
type CartClear struct {
	XMLName           xml.Name            `xml:"CartClear"`
	MarketplaceDomain string              `xml:"MarketplaceDomain,omitempty"`
	AWSAccessKeyID    string              `xml:"AWSAccessKeyId,omitempty"`
	AssociateTag      string              `xml:"AssociateTag,omitempty"`
	Validate          string              `xml:"Validate,omitempty"`
	XMLEscaping       string              `xml:"XMLEscaping,omitempty"`
	Shared            *CartClearRequest   `xml:"Shared,omitempty"`
	Request           []*CartClearRequest `xml:"Request,omitempty"`
}

// CartClearRequest is request parameters for CartClear
//
// read more: http://docs.aws.amazon.com/AWSECommerceService/latest/DG/CartClear.html#CartClear-req
type CartClearRequest struct {
	CartID        string   `xml:"CartId,omitempty"`
	HMAC          string   `xml:"HMAC,omitempty"`
	MergeCart     string   `xml:"MergeCart,omitempty"`
	ResponseGroup []string `xml:"ResponseGroup,omitempty"`
}

// CartClearResponse is response for CartClear
//
// read more: http://docs.aws.amazon.com/AWSECommerceService/latest/DG/CartClear.html#CartClear-resp
type CartClearResponse struct {
	XMLName          xml.Name          `xml:"CartClearResponse"`
	OperationRequest *OperationRequest `xml:"OperationRequest,omitempty"`
	Cart             []*Cart           `xml:"Cart,omitempty"`
}
