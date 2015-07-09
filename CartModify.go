package agozon
import "strconv"

type CartModifyRequest struct {
	*CartCommon
	Items CartItems
}

type CartModifyRequest struct {
	OperationRequest OperationRequest `xml:"OperationRequest,omitempty" json:",omitempty"`
	Cart             Cart             `xml:",omitempty" json:",omitempty"`
}

func (c *remoteCart) CartModify(cart Cart, items ...CartItem) (response CartAddResponse, err error) {
	q := map[string]string{}

	q["CartId"] = cart.CartId
	q["HMAC"] = cart.HMAC

	for index, item := range items {
		key := "Item." + strconv.Itoa(index+1)
		if item.CartItemId != "" {
			q[key+".CartItemId"] = item.CartItemId
			// action
			if item.Action == "MoveToCart" || item.Action == "SaveForLater" {
				q[key+".Action"] = item.Action
			}
			// add quantity
			if item.Quantity < 1 {
				item.Quantity = 1
			}
			q[key+".Quantity"] = strconv.Itoa(item.Quantity)
		}
	}

	// call cart do
	response, err = c.do("CartModify", q)
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
