package agozon
import "strconv"

type CartAddRequest struct {
	*CartCommon
	Items CartItems
}

type CartAddResponse struct {
	OperationRequest OperationRequest `xml:"OperationRequest,omitempty" json:",omitempty"`
	Cart             Cart             `xml:",omitempty" json:",omitempty"`
}


func (c *remoteCart) CartAdd(cart Cart, items ...Item) (response CartAddResponse, err error) {
	q := map[string]string{}

	q["CartId"] = cart.CartId
	q["HMAC"] = cart.HMAC

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
	response, err = c.do("CartAdd", q)

	// catch the error
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
