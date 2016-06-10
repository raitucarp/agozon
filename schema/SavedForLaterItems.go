package schema

import (
	"encoding/xml"
)

// SavedForLaterItems is A positive integer that uniquely identifies an item in Saved For Later.
//
// Ancestry
//
// Cart/SavedForLaterItems
//
// Cart
//
// More details:
// 	- http://docs.aws.amazon.com/AWSECommerceService/latest/DG/CHAP_response_elements.html
// 	- http://docs.aws.amazon.com/AWSECommerceService/latest/DG/ShoppingCartConcepts.html#ActiveandSaveForLaterAreas
// 	- http://docs.aws.amazon.com/AWSECommerceService/latest/DG/AddingItemsAsSavedForLater.html
// 	- http://docs.aws.amazon.com/AWSECommerceService/latest/DG/CartCreate.html#SampleSavedForLaterItemXMLSnippet
type SavedForLaterItems struct {
	XMLName           xml.Name    `xml:"SavedForLaterItems"`
	SubTotal          *Price      `xml:"SubTotal,omitempty"`
	SavedForLaterItem []*CartItem `xml:"SavedForLaterItem,omitempty"`
}
