package schema

import (
	"encoding/xml"
)

// CartModify is a schema for CartModify operation
//
// The CartModify operation enables you to change the quantity of items that are already in a remote shopping cart
// and move items from the active area of a cart to the SaveForLater area or the reverse.
//
// To modify the number of items in a cart, you must specify the cart using the CartId and HMAC values that are returned in the CartCreate operation.
// A value similar to HMAC, URLEncodedHMAC, is also returned.
// This value is the URL encoded version of the HMAC. This encoding is necessary because some characters, such as + and /,
// cannot be included in a URL.
// Rather than encoding the HMAC yourself, use the URLEncodedHMAC value for the HMAC parameter.
//
// You can use CartModify to modify the number of items in a remote shopping cart by setting the value of the Quantity parameter appropriately.
// You can eliminate an item from a cart by setting the value of the Quantity parameter to zero.
// Or, you can double the number of a particular item in the cart by doubling its Quantity.
// You cannot, however, use CartModify to add new items to a cart.
//
// All CartModify requests must also include the value for AssociateTag
// that was used in the associated CartCreate request; otherwise, the request will fail.
type CartModify struct {
	XMLName           xml.Name             `xml:"CartModify"`
	MarketplaceDomain string               `xml:"MarketplaceDomain,omitempty"`
	AWSAccessKeyID    string               `xml:"AWSAccessKeyId,omitempty"`
	AssociateTag      string               `xml:"AssociateTag,omitempty"`
	Validate          string               `xml:"Validate,omitempty"`
	XMLEscaping       string               `xml:"XMLEscaping,omitempty"`
	Shared            *CartModifyRequest   `xml:"Shared,omitempty"`
	Request           []*CartModifyRequest `xml:"Request,omitempty"`
}

// CartModifyRequest are request parameters for CartModify
//
// More details: http://docs.aws.amazon.com/AWSECommerceService/latest/DG/CartModify.html#CartModify-rp
type CartModifyRequest struct {
	CartID        string                   `xml:"CartId,omitempty"`
	HMAC          string                   `xml:"HMAC,omitempty"`
	MergeCart     string                   `xml:"MergeCart,omitempty"`
	Items         []*CartModifyRequestItem `xml:"Items>Item,omitempty"`
	ResponseGroup []string                 `xml:"ResponseGroup,omitempty"`
}

// CartModifyRequestItem is Item type for CartModifyRequest
type CartModifyRequestItem struct {
	Action     string  `xml:"Action,omitempty"`
	CartItemID string  `xml:"CartItemId,omitempty"`
	Quantity   *uint64 `xml:"Quantity,omitempty"`
}

// CartModifyResponse is a response for CartModify
//
// More details: http://docs.aws.amazon.com/AWSECommerceService/latest/DG/CartModify.html#CartModify-resp
type CartModifyResponse struct {
	XMLName          xml.Name          `xml:"CartModifyResponse"`
	OperationRequest *OperationRequest `xml:"OperationRequest,omitempty"`
	Cart             []*Cart           `xml:"Cart,omitempty"`
}
