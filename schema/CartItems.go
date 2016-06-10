package schema

import (
	"encoding/xml"
)

// CartItems is a parent element for many child elements, including SubTotal, and CartItem.
//
// Ancestry: Cart
type CartItems struct {
	XMLName  xml.Name    `xml:"CartItems"`
	SubTotal *Price      `xml:"SubTotal,omitempty"`
	CartItem []*CartItem `xml:"CartItem,omitempty"`
}
