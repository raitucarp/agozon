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
	searchModernBooks.SearchIndex = "Books"
	searchModernBooks.Keywords("modern")
	searchModernBooks.Sort = "price"
	searchModernBooks.OnlyAvailable(true)
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
		fmt.Println(item.ASIN, item.ItemAttributes.Title)
	}
}

func TestItemLookupRequest(t *testing.T) {
	lookupItem := testRequest.ItemLookup("B00IJYII4E", "B00IEG7ULO")
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

func TestCartCreate(t *testing.T) {
	cart := testRequest.Cart()
	response, err := cart.Create(
		Item{
			ASIN:     "B00IJYII4E",
			Quantity: 1,
		},
		Item{
			ASIN: "B00IEG7ULO",
			Quantity: 2,
		},
	)

	if err != nil {
		t.Error("It's an error", err)
		return
	}
	fmt.Println("print carts")
	for _, cartInfo := range response.Cart{
		fmt.Println("--------------------")
		fmt.Println("CartID", cartInfo.CartID)
		fmt.Println("PurchaseURL", cartInfo.PurchaseURL)
		fmt.Println("HMAC", cartInfo.HMAC)
		fmt.Println("MobileCartURL", cartInfo.MobileCartURL)
	}

}
