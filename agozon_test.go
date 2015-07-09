package agozon

import (
	"fmt"
	"testing"
)

var testRequest = NewRequest(&Config{
	AssociateTag:    "hoxr0a-20",
	AWSAccessKeyId:  "AKIAJHWAQ26RR23FAORQ",
	SecretAccessKey: "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",
})

func TestItemSearchRequest(t *testing.T) {
	searchModernBooks := testRequest.ItemSearch()
	//searchMelon.Keywords("calculus")
	searchModernBooks.SearchIndex = "All"
	searchModernBooks.Keywords("modern")
	searchModernBooks.Sort = "price"
	//searchModernBooks.OnlyAvailable(true)
	searchModernBooks.ResponseGroup("Large", "ItemAttributes", "OfferFull")
	response, err := searchModernBooks.Do()
	if err != nil {
		t.Error("It's an error", err)
		return
	}

	/*for _, error := range response.OperationRequest.Errors {}*/
	fmt.Println("length of error", response.Items.Request.Errors)

	items := response.GetItems()
	for _, item := range items {
		fmt.Println("-----")
		fmt.Println(item.ASIN, item.ItemAttributes.Title)
		for _, offer := range *item.Offers.Offer {
			fmt.Println("Offer ListingId:", offer.OfferListing.OfferListingId)
		}
	}
}

func TestItemLookupRequest(t *testing.T) {
	lookupItem := testRequest.ItemLookup("1594206279", "B00A6O719S")
	lookupItem.ResponseGroup("Large")
	response, err := lookupItem.Do()
	if err != nil {
		t.Error("It's an error", err)
		return
	}
	/*for _, error := range response.OperationRequest.Errors {}*/
	fmt.Println("length of error", response.Items.Request.Errors)

	items := response.GetItems()
	for _, item := range items {
		fmt.Println(item.ASIN, item.ItemAttributes.Title)
	}
}

func TestCart(t *testing.T) {
	cart := testRequest.NewRemoteCart()
	cart.ResponseGroup("Cart", "CartSimilarities")

	// CartCreate
	fmt.Println("CartCreate start:")
	response, err := cart.Create(
		Item{
			OfferListingId: "2sv2itag6LJqu3Xj8C%2BB3xo%2FW6ROyQuuM0ecFd8x5ycKXdcOmvZQjrxYWJxQGVFszkoo5zwg4QJoUJjFchjotTr2vigST%2BwcHVUEZMotpUYblRXbc2fdPA%3D%3D",
			Quantity: 1,
		},
		Item{
			OfferListingId: "%2FBnJJLPPAnZNQrI0ogGqSE4H9269cqp6IltpITgNVNU50Qt6rnR74KgFjiZYJQ2nLqm9txuI%2Bjlc4%2ByAmFyilZv8JaUq023sQOamk8zv%2BE%2BzE7p0QCgcZCYBm4hWoPLdk9jCNA4nzrOli1RYsUh4KONMG%2FPPIeiF",
			Quantity: 2,
		},
	)

	if err != nil {
		t.Error("It's an error", err)
		return
	}
	fmt.Println("print carts")
	cartInfo := response.Cart
	fmt.Println("--------------------")
	fmt.Println("CartID", cartInfo.CartId)
	fmt.Println("PurchaseURL", cartInfo.PurchaseURL)
	fmt.Println("HMAC", cartInfo.HMAC)
	fmt.Println("URLEncodedHMAC", cartInfo.URLEncodedHMAC)
	fmt.Println("--------------------")

	// CartGet
	fmt.Println("CartGet")
	info, err := cart.Get(cartInfo)
	if err != nil {
		t.Error("It's an error", err)
		return
	}
	fmt.Println("CartGet Info", info.Cart.CartId, err, cartInfo.CartItems.CartItem)

	/*fmt.Println("-----")
	info, err = cart.Clear(info.Cart.CartId, info.Cart.HMAC)
	fmt.Println("CartClear", info.Cart, err)*/

}
