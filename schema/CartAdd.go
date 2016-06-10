package schema

import (
	"encoding/xml"
)

// CartAdd operation enables you to add items to an existing remote shopping cart. CartAdd can only be used to place a new item in a shopping cart.
// It cannot be used to increase the quantity of an item already in the cart.
// If you would like to increase the quantity of an item that is already in the cart, you must use the CartModify operation.
//
// You add an item to a cart by specifying the item's OfferListingId, or ASIN and ListItemId.
// Once in a cart, an item can only be identified by its CartItemId.
// That is, an item in a cart cannot be accessed by its ASIN or OfferListingId.
// CartItemId is returned by CartCreate, CartGet, and CartAdd.
// To add items to a cart, you must specify the cart using the CartId and HMAC values,
// which are returned by the CartCreate operation.
// If the associated CartCreate request specified an AssociateTag,
// all CartAdd requests must also include a value for Associate Tag otherwise the request will fail.
//
// Read more: http://docs.aws.amazon.com/AWSECommerceService/latest/DG/CartAdd.html#CartAdd-des
type CartAdd struct {
	XMLName           xml.Name          `xml:"CartAdd"`
	MarketplaceDomain string            `xml:"MarketplaceDomain,omitempty"`
	AWSAccessKeyID    string            `xml:"AWSAccessKeyId,omitempty"`
	AssociateTag      string            `xml:"AssociateTag,omitempty"`
	Validate          string            `xml:"Validate,omitempty"`
	XMLEscaping       string            `xml:"XMLEscaping,omitempty"`
	Shared            *CartAddRequest   `xml:"Shared,omitempty"`
	Request           []*CartAddRequest `xml:"Request,omitempty"`
}

// CartAddRequest is request parameters for CartAdd.
//
// More details: http://docs.aws.amazon.com/AWSECommerceService/latest/DG/CartAdd.html#CartAdd-req
type CartAddRequest struct {
	CartID        string                `xml:"CartId,omitempty"`
	HMAC          string                `xml:"HMAC,omitempty"`
	MergeCart     string                `xml:"MergeCart,omitempty"`
	Items         []*CartAddRequestItem `xml:"Items>Item,omitempty"`
	ResponseGroup []string              `xml:"ResponseGroup,omitempty"`
}

// CartAddRequestItem is item in CartAddRequestItems
//
// See CartAddRequest and CartAddRequestItems
type CartAddRequestItem struct {
	ASIN           string  `xml:"ASIN,omitempty"`
	OfferListingID string  `xml:"OfferListingId,omitempty"`
	Quantity       *uint64 `xml:"Quantity,omitempty"`
	AssociateTag   string  `xml:"AssociateTag,omitempty"`
	ListItemID     string  `xml:"ListItemId,omitempty"`
}

// CartAddResponse is a response for CartAdd
//
// read more: http://docs.aws.amazon.com/AWSECommerceService/latest/DG/CartAdd.html#CartAdd-resp
type CartAddResponse struct {
	XMLName          xml.Name          `xml:"CartAddResponse"`
	OperationRequest *OperationRequest `xml:"OperationRequest,omitempty"`
	Cart             []*Cart           `xml:"Cart,omitempty"`
}
