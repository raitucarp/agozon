package schema

import (
	"encoding/xml"
)

// Cart response group provides information about a specified remote shopping cart and the items in it.
// The cart information includes:
//
// 	1) CartId
//	2) HMAC
//	3) PurchaseURL
//
// For each item in the cart, including SavedForLaterItems, the response group returns:
// 	- CartItemId
// 	- ProductName
// 	- ASIN
// 	- Quantity
// 	- ListPrice
// 	- OurPrice
//
// 	For more information read:
// 	http://docs.aws.amazon.com/AWSECommerceService/latest/DG/RG_Cart.html
type Cart struct {
	XMLName                        xml.Name                        `xml:"Cart"`
	Request                        *Request                        `xml:"Request,omitempty"`
	CartID                         string                          `xml:"CartId,omitempty"`
	HMAC                           string                          `xml:"HMAC,omitempty"`
	URLEncodedHMAC                 string                          `xml:"URLEncodedHMAC,omitempty"`
	PurchaseURL                    string                          `xml:"PurchaseURL,omitempty"`
	MobileCartURL                  string                          `xml:"MobileCartURL,omitempty"`
	SubTotal                       *Price                          `xml:"SubTotal,omitempty"`
	CartItems                      *CartItems                      `xml:"CartItems,omitempty"`
	SavedForLaterItems             *SavedForLaterItems             `xml:"SavedForLaterItems,omitempty"`
	SimilarProducts                *SimilarProducts                `xml:"SimilarProducts,omitempty"`
	TopSellers                     *TopSellers                     `xml:"TopSellers,omitempty"`
	NewReleases                    *NewReleases                    `xml:"NewReleases,omitempty"`
	SimilarViewedProducts          *SimilarViewedProducts          `xml:"SimilarViewedProducts,omitempty"`
	OtherCategoriesSimilarProducts *OtherCategoriesSimilarProducts `xml:"OtherCategoriesSimilarProducts,omitempty"`
}
