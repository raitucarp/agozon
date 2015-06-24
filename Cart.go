package agozon

func (r *Request) Cart() *Cart {
	c := &Cart{}

	// define do
	c.do = func(operationName string, q map[string]string) (get func(response interface{}) (err error), err error) {
		r.create()
		// add queries to api request
		r.AddParams(q)
		err = r.CallOperation(operationName)
		if err != nil {
			return
		}

		// returning response
		get = r.Get
		return get, nil
	}

	return c
}

type Cart struct {
	Request                        *Request                        `xml:",omitempty" json:",omitempty"`
	CartID                         string                          `xml:",omitempty" json:",omitempty"`
	CartItems                      *CartItems                      `xml:",omitempty" json:",omitempty"`
	HMAC                           string                          `xml:",omitempty" json:",omitempty"`
	MobileCartURL                  string                          `xml:",omitempty" json:",omitempty"`
	NewReleases                    *NewReleases                    `xml:",omitempty" json:",omitempty"`
	OtherCategoriesSimilarProducts *OtherCategoriesSimilarProducts `xml:",omitempty" json:",omitempty"`
	PurchaseURL                    string                          `xml:",omitempty" json:",omitempty"`
	SavedForLaterItems             *SavedForLaterItems             `xml:",omitempty" json:",omitempty"`
	SimilarProducts                *SimilarProducts                `xml:"SimilarProducts>SimilarProduct,omitempty" json:",omitempty"`
	SimilarViewedProducts          *SimilarViewedProducts          `xml:"SimilarViewedProducts>SimilarViewedProduct,omitempty" json:",omitempty"`
	SubTotal                       *Price                          `xml:",omitempty" json:",omitempty"`
	TopSellers                     *TopSellers                     `xml:"TopSellers>TopSeller,omitempty" json:",omitempty"`
	URLEncodedHMAC                 string                          `xml:",omitempty" json:",omitempty"`

	// function
	do func(string, map[string]string) (func(response interface{}) (err error), error)
}

func (c *Cart) Add() {

}

func (c *Cart) Get() {

}

func (c *Cart) Modify() {

}

func (c *Cart) Clear() {

}
