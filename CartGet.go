package agozon

type CartGetRequest struct {
	*CartCommon
	Items CartItems
}

type CartGetResponse struct {
	OperationRequest OperationRequest `xml:"OperationRequest,omitempty" json:",omitempty"`
	Cart             Cart             `xml:",omitempty" json:",omitempty"`
}

func (c *remoteCart) Get(cart Cart) (response CartGetResponse, err error) {
	q := map[string]string{}

	q["CartId"] = cart.CartId
	q["HMAC"] = cart.HMAC
	// call cart do
	response, err = c.do("CartGet", q)
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
