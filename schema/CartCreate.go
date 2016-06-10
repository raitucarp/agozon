package schema

import (
	"encoding/xml"
)

// CartCreate operation enables you to create a remote shopping cart.
// A shopping cart is the metaphor used by most e-commerce solutions.
// It is a temporary data storage structure that resides on Amazon servers.
// The structure contains the items a customer wants to buy.
// In Product Advertising API, the shopping cart is considered remote because it is hosted by Amazon servers.
// In this way, the cart is remote to the vendor's web site where the customer views and selects the items they want to purchase.
//
// Once you add an item to a cart by specifying the item's ASIN or OfferListing ID,
// the item is assigned a CartItemId and is accessible only by that value.
// That is, in subsequent requests, an item in a cart cannot be accessed by its ASIN or OfferListingId.
// CartItemId is returned by CartCreate, CartGet, and CartAdd.
//
// Because the contents of a cart can change for different reasons, such as item availability, you should not keep a copy of a cart locally.
// Instead, use the other cart operations to modify the cart contents.
// For example, to retrieve contents of the cart, which are represented by CartItemIds, use CartGet.
//
// Available products are added as cart items. Unavailable items, for example,
// items out of stock, discontinued, or future releases, are added as SaveForLaterItems. No error is generated.
// The Amazon database changes regularly.
// You may find a product with an offer listing ID but by the time the item is added to the cart the product is no longer available.
// The checkout page in the Order Pipeline clearly lists items that are available and those that are SaveForLaterItems.
//
// It is impossible to create an empty shopping cart. You have to add at least one item to a shopping cart using a single CartCreate request.
// You can add specific quantities (up to 999) of each item.
//
// CartCreate can be used only once in the life cycle of a cart.
// To modify the contents of the cart, use one of the other cart operations.
//
// Carts cannot be deleted. They expire automatically after being unused for 7 days.
// The lifespan of a cart restarts, however, every time a cart is modified.
// In this way, a cart can last for more than 7 days.
// If, for example, on day 6, the customer modifies a cart, the 7 day countdown starts over.
//
// More details: http://docs.aws.amazon.com/AWSECommerceService/latest/DG/CartCreate.html#CartCreate-desc
type CartCreate struct {
	XMLName           xml.Name             `xml:"CartCreate"`
	MarketplaceDomain string               `xml:"MarketplaceDomain,omitempty"`
	AWSAccessKeyID    string               `xml:"AWSAccessKeyId,omitempty"`
	AssociateTag      string               `xml:"AssociateTag,omitempty"`
	Validate          string               `xml:"Validate,omitempty"`
	XMLEscaping       string               `xml:"XMLEscaping,omitempty"`
	Shared            *CartCreateRequest   `xml:"Shared,omitempty"`
	Request           []*CartCreateRequest `xml:"Request,omitempty"`
}

// CartCreateRequest is request parameters for CartCreate
//
// More details: http://docs.aws.amazon.com/AWSECommerceService/latest/DG/CartCreate.html#CartCreate-rp
type CartCreateRequest struct {
	XMLName       xml.Name                 `xml:"CartCreateRequest"`
	MergeCart     string                   `xml:"MergeCart,omitempty"`
	Items         []*CartCreateRequestItem `xml:"Items>Item,omitempty"`
	ResponseGroup []string                 `xml:"ResponseGroup,omitempty"`
}

// CartCreateRequestItem is item for CartCreateRequestItems
type CartCreateRequestItem struct {
	ASIN           string      `xml:"ASIN,omitempty"`
	OfferListingID string      `xml:"OfferListingId,omitempty"`
	Quantity       *uint64     `xml:"Quantity,omitempty"`
	AssociateTag   string      `xml:"AssociateTag,omitempty"`
	ListItemID     string      `xml:"ListItemId,omitempty"`
	MetaData       *[]Metadata `xml:"MetaData,omitempty"`
}

// CartCreateResponse is response for CartCreate
//
// More details: http://docs.aws.amazon.com/AWSECommerceService/latest/DG/CartCreate.html#CartCreate-resp
type CartCreateResponse struct {
	XMLName          xml.Name          `xml:"CartCreateResponse"`
	OperationRequest *OperationRequest `xml:"OperationRequest,omitempty"`
	Cart             []*Cart           `xml:"Cart,omitempty"`
}
