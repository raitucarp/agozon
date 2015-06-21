package agozon

import (
	"fmt"
	"testing"
)

func TestItemSearchRequest(t *testing.T) {
	testRequest := NewRequest(&Config{
		AssociateTag:    "hoxr0a-20",
		AWSAccessKeyId:  "AKIAJHWAQ26RR23FAORQ",
		SecretAccessKey: "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",
	})

	searchMelon := testRequest.ItemSearch()
	searchMelon.Keywords("Calculus")
	searchMelon.SearchIndex = "Books"
	searchMelon.ResponseGroup("Large", "ItemAttributes")
	response, err := searchMelon.Do()
	if err != nil {
		t.Error("It's an error", err)
		return
	}
	fmt.Println(response)

	/*for _, error := range response.OperationRequest.Errors {}*/
	fmt.Println("length of error", response.Items.Request.Errors)

	items := response.GetItems()
	for _, item := range items {
		ImageSets := item.ImageSets
		for _, imageset := range ImageSets {

			fmt.Println(imageset.LargeImage.URL)
		}
	}
	/*testRequest.ResponseGroup([]string{
		"Images",
		"ItemAttributes",
		"Offers",
		"Reviews",
	})
	testRequest.ItemLookup(map[string]string{
		"ItemId": "0679722769",
	})*/

}
