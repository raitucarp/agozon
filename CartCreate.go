package agozon

import (
	"strconv"
)

type CartCreateRequest struct {
	*CartCommon
	Items CartItems
}

type CartCreateResponse struct {
	OperationRequest OperationRequest `xml:"OperationRequest,omitempty" json:",omitempty"`
	Cart             Cart             `xml:",omitempty" json:",omitempty"`
}

/* CartCreate */
// The CartCreate operation enables you to create a remote shopping cart.
// A shopping cart is the metaphor used by most e-commerce solutions.
// It is a temporary data storage structure that resides on Amazon servers.
// The structure contains the items a customer wants to buy.
// In Product Advertising API, the shopping cart is considered remote because it is hosted by Amazon servers.
// In this way, the cart is remote to the vendor's web site where the customer views and selects the
// items they want to purchase.
func (c *remoteCart) Create(items ...Item) (response CartCreateResponse, err error) {
	q := map[string]string{}

	for index, item := range items {
		key := "Item." + strconv.Itoa(index+1)
		if item.OfferListingId != "" {
			q[key+".OfferListingId"] = item.OfferListingId
			// add quantity
			if item.Quantity < 1 {
				item.Quantity = 1
			}
			q[key+".Quantity"] = strconv.Itoa(item.Quantity)
		}
	}

	// call cart do
	response, err = c.do("CartCreate", q)
	if err != nil {
		return
	}

	// detect error from response
	if response.Cart.Request.Errors != nil {
		return response, response.Cart.Request.Errors
	}

	// return response and error
	return
}
