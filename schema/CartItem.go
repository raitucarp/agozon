package schema

// CartItem is a parent element for many child elements, including CartItemId, Quantity, Title, ProductGroup, Price, and ItemTotal.
//
// Ancestry: Cart/CartItems
type CartItem struct {
	//XMLName        xml.Name `xml:"CartItem"`
	CartItemID     string      `xml:"CartItemId,omitempty"`
	ASIN           string      `xml:"ASIN,omitempty"`
	SellerNickname string      `xml:"SellerNickname,omitempty"`
	Quantity       string      `xml:"Quantity,omitempty"`
	Title          string      `xml:"Title,omitempty"`
	ProductGroup   string      `xml:"ProductGroup,omitempty"`
	MetaData       *[]Metadata `xml:"MetaData>KeyValuePair,omitempty"`
	Price          *Price      `xml:"Price,omitempty"`
	ItemTotal      *Price      `xml:"ItemTotal,omitempty"`
}
