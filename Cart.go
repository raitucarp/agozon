package agozon

type remoteCart struct {
	request       *Request
	responseGroup []string `xml:",omitempty" json:",omitempty"`
}

func (c *remoteCart) do(operationName string, q map[string]string) (response interface{}, err error) {
	c.request.create()
	c.validate(operationName)
	// add queries to api request
	c.request.AddParams(q)
	err = c.request.CallOperation(operationName)
	if err != nil {
		return
	}

	// returning response
	err = c.request.Get(&response)
	if err != nil {
		return
	}

	return
}

func (c *remoteCart) validate(operationName string)  {
	responseGroup := []string{}
	for _, val := range c.responseGroup {
		for _, validRG := range ValidResponseGroup[operationName] {
			if val == validRG {
				responseGroup = append(responseGroup, val)
			}
		}
	}
	c.request.SetResponseGroup(responseGroup...)
}

func (c *remoteCart) ResponseGroup(responseGroup ...string) {
	c.responseGroup = responseGroup
}

func (r *Request) NewRemoteCart() *remoteCart {
	c := &remoteCart{request: &r}
	return c
}

type Cart struct {
	Request                        *Request                        `xml:",omitempty" json:",omitempty"`
	CartId                         string                          `xml:",omitempty" json:",omitempty"`
	HMAC                           string                          `xml:",omitempty" json:",omitempty"`
	URLEncodedHMAC                 string                          `xml:",omitempty" json:",omitempty"`
	PurchaseURL                    string                          `xml:",omitempty" json:",omitempty"`
	MobileCartURL                  string                          `xml:",omitempty" json:",omitempty"`
	SubTotal                       *Price                          `xml:",omitempty" json:",omitempty"`
	CartItems                      *CartItems                      `xml:",omitempty" json:",omitempty"`
	SavedForLaterItems             *SavedForLaterItems             `xml:",omitempty" json:",omitempty"`
	SimilarProducts                *SimilarProducts                `xml:"SimilarProducts>SimilarProduct,omitempty" json:",omitempty"`
	TopSellers                     *TopSellers                     `xml:"TopSellers>TopSeller,omitempty" json:",omitempty"`
	NewReleases                    *NewReleases                    `xml:",omitempty" json:",omitempty"`
	SimilarViewedProducts          *SimilarViewedProducts          `xml:"SimilarViewedProducts>SimilarViewedProduct,omitempty" json:",omitempty"`
	OtherCategoriesSimilarProducts *OtherCategoriesSimilarProducts `xml:",omitempty" json:",omitempty"`
}
