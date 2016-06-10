package schema

// Price is A parent element for Amount, CurrencyCode, and FormattedPrice.
// In this case, the price is for an item in the Active or Saved For Later areas, respectively.
//
// Ancestry: CartItem
//
// SavedForLaterItem
//
// More details: http://docs.aws.amazon.com/AWSECommerceService/latest/DG/CHAP_response_elements.html
type Price struct {
	Amount         int32  `xml:"Amount,omitempty"`
	CurrencyCode   string `xml:"CurrencyCode,omitempty"`
	FormattedPrice string `xml:"FormattedPrice,omitempty"`
}
