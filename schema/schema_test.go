package schema

import (
	"encoding/xml"
	"fmt"
	"github.com/raitucarp/agozon/schema/Condition"
	"github.com/raitucarp/agozon/schema/audience-rating"
	"io/ioutil"
	"regexp"
	"strings"
	"testing"
)

// ReadSampleXML is read xml file in sample-xml
func ReadSampleXML(filename string) (content string, err error) {
	// format filename to be match with sample-xml
	filename = fmt.Sprintf("./sample-xml/%s.xml", filename)

	// read file into data
	data, err := ioutil.ReadFile(filename)

	// if error then return it with error
	if err != nil {
		return
	}

	// convert []byte of data to string, and return nil error
	return string(data), nil
}

// Unmarshal read sample-xml file and unmarshal it to type
func Unmarshal(filename string, v interface{}) (data string, err error) {
	// read data
	data, err = ReadSampleXML(filename)
	if err != nil {
		return
	}

	// unmarshaling to type
	err = xml.Unmarshal([]byte(data), &v)
	if err != nil {
		return
	}
	return
}

// selfClosing pattern is for replacement tags with no data to be self-closing tag
var selfClosingPattern = regexp.MustCompile(`></(.*)>`)

// MarshalAndCompare is marshaling type and compare it with sample-xml
func MarshalAndCompare(data string, v interface{}) (outputString string, equal bool, err error) {
	// replace \r from windows new line
	data = strings.Replace(data, "\r", "", -1)

	// marshal with indent two space and no prefix
	output, err := xml.MarshalIndent(v, "", "  ")
	if err != nil {
		return
	}

	// lower xml header to be match with sample-xml
	modifiedHeader := strings.ToLower(xml.Header)

	// outputString is modified header with string of output
	outputString = modifiedHeader + string(output)
	// replace tags with no data with self closing tag
	// to be match with sample-xml
	outputString = selfClosingPattern.ReplaceAllString(outputString, ` />`)

	// compare output string with data, and return equal if it is equal
	if strings.Compare(outputString, data) == 0 {
		equal = true
		return
	}

	// return default value
	return
}

// test is the main task for doing testing.
// the strategy is compare the marshal content with original sample-xml.
// If it's matched then the schema is correct
func test(filename string, v interface{}, t *testing.T) {
	// Unmarshal file to v
	data, err := Unmarshal(filename, v)
	// if error then give testing error
	if err != nil {
		t.Error(err)
	}

	// get output string from marshaling
	outputString, equal, err := MarshalAndCompare(data, v)
	// if error then give testing error
	if err != nil {
		t.Error(err)
	}

	// check if before and after marshaling equal or not,
	// raise error when it's not
	if !equal {
		t.Error(filename, " > Before and after marshaling not equal")
		//t.Log(data)
		t.Log(outputString)
		t.Fail()
	}

}

func TestAccessories(t *testing.T) {
	var v Accessories
	test("Accessories", &v, t)
}

func TestArguments(t *testing.T) {
	var v Arguments
	test("Arguments", &v, t)
}

func TestAudienceRating(t *testing.T) {
	var rating audience.Rating
	test("AudienceRating", &rating, t)

	if rating.Valid() == false {
		t.Error("It seems not valid audience rating")
	}

	// test invalid rating
	invalidRating := "<AudienceRating>invalidRating</AudienceRating>"
	err := xml.Unmarshal([]byte(invalidRating), &rating)

	if !rating.Valid() == false {
		t.Error(err)
	}

}

func TestBin(t *testing.T) {
	var v Bin
	test("Bin", &v, t)
}

func TestBrowseNode(t *testing.T) {
	var node BrowseNode
	test("BrowseNode", &node, t)

	if len(node.GetProperties()) < 0 {
		t.Error("Node less than 0")
	}

	if len(node.GetChildren()) < 0 {
		t.Error("Children less than 0")
	}

	if len(node.GetAncestors()) < 0 {
		t.Error("Ancestors less than 0")
	}
}

func TestBrowseNodeLookup(t *testing.T) {
	var v BrowseNodeLookup
	test("BrowseNodeLookup", &v, t)

	//t.Run("Request", func(t *testing.T) {
	var request BrowseNodeLookupRequest
	test("BrowseNodeLookupRequest", &request, t)
	//})

	//t.Run("Response", func(t *testing.T) {
	var response BrowseNodeLookupResponse
	test("BrowseNodeLookupResponse", &response, t)
	//})
}

func TestBrowseNodes(t *testing.T) {
	var v BrowseNodes
	test("BrowseNodes", &v, t)
}

func TestCart(t *testing.T) {
	var c Cart
	test("Cart", &c, t)
}

func TestCartAdd(t *testing.T) {
	var c CartAdd
	test("CartAdd", &c, t)

	//t.Run("Request", func(t *testing.T) {
	var request CartAddRequest
	test("CartAddRequest", &request, t)
	//})

	//t.Run("Response", func(t *testing.T) {
	var response CartAddResponse
	test("CartAddResponse", &response, t)
	//})
}

func TestCartClear(t *testing.T) {
	var c CartClear
	test("CartClear", &c, t)

	// CartClear request
	//t.Run("Request", func(t *testing.T) {
	var request CartClearRequest
	test("CartClearRequest", &request, t)
	//})

	// Response of CartClear
	//t.Run("Response", func(t *testing.T) {
	var response CartClearResponse
	test("CartClearResponse", &response, t)
	//})
}

func TestCartGet(t *testing.T) {
	var c CartGet
	test("CartGet", &c, t)

	// request of CartGet
	//t.Run("Request", func(t *testing.T) {
	var request CartGetRequest
	test("CartGetRequest", &request, t)
	//})

	// response of CartGet
	//t.Run("Response", func(t *testing.T) {
	var response CartGetResponse
	test("CartGetResponse", &response, t)
	//})
}

func TestCartItem(t *testing.T) {
	var i CartItem

	test("CartItem", &i, t)
}

func TestCartItems(t *testing.T) {
	var i CartItems

	test("CartItems", &i, t)
}

func TestCartModify(t *testing.T) {
	var c CartModify

	test("CartModify", &c, t)

	// request of CartModify
	//t.Run("Request", func(t *testing.T) {
	var request CartModifyRequest

	test("CartModifyRequest", &request, t)
	//})

	// response of CartModify
	//t.Run("Response", func(t *testing.T) {
	var response CartModifyResponse

	test("CartModifyResponse", &response, t)
	//})
}

func TestCollections(t *testing.T) {
	var c Collections

	test("Collections", &c, t)
}

func TestCondition(t *testing.T) {
	var c condition.Condition = 0

	test("Condition", &c, t)
}

func TestCorrectedQuery(t *testing.T) {
	var cq CorrectedQuery
	test("CorrectedQuery", &cq, t)
}

func TestCustomerReviews(t *testing.T) {
	var cr CustomerReviews
	test("CustomerReviews", &cr, t)
}

func TestDecimalWithUnits(t *testing.T) {
	var d DecimalWithUnits
	test("DecimalWithUnits", &d, t)

	if d.Value != "1" {
		t.Error("Value is not equal to 1")
	}
}

func TestEditorialReview(t *testing.T) {
	var er EditorialReview
	test("EditorialReview", &er, t)
}

func TestEditorialReviews(t *testing.T) {
	var er EditorialReviews
	test("EditorialReviews", &er, t)
}

func TestErrors(t *testing.T) {
	var e Errors
	test("Errors", &e, t)
}

func TestHTTPHeaders(t *testing.T) {
	var header HTTPHeaders
	test("HTTPHeaders", &header, t)
}

func TestImage(t *testing.T) {
	var i Image
	test("Image", &i, t)
}

func TestImageSet(t *testing.T) {
	var i ImageSet
	test("ImageSet", &i, t)
}

func TestItem(t *testing.T) {
	var item Item
	test("Item", &item, t)
}

func TestItemAttributes(t *testing.T) {
	var itemAttr ItemAttributes
	test("ItemAttributes", &itemAttr, t)
}

func TestItemLink(t *testing.T) {
	var i ItemLink
	test("ItemLink", &i, t)
}

func TestItemLinks(t *testing.T) {
	var i ItemLinks
	test("ItemLinks", &i, t)
}

func TestItemLookup(t *testing.T) {
	var i ItemLookup
	test("ItemLookup", &i, t)

	//t.Run("Request", func(t *testing.T) {
	var request ItemLookupRequest
	test("ItemLookupRequest", &request, t)
	//})

	//t.Run("Response", func(t *testing.T) {
	var response ItemLookupResponse
	test("ItemLookupResponse", &response, t)
	//})
}

func TestItems(t *testing.T) {
	var i Items
	test("Items", &i, t)
}

func TestItemSearch(t *testing.T) {
	var i ItemSearch
	test("ItemSearch", &i, t)

	// request of ItemSearch
	//t.Run("Request", func(t *testing.T) {
	var request ItemSearchRequest
	test("ItemSearchRequest", &request, t)
	//})

	// response of ItemSearch
	//t.Run("Response", func(t *testing.T) {
	var response ItemSearchResponse
	test("ItemSearchResponse", &response, t)
	//})
}

func TestLoyaltyPoints(t *testing.T) {
	var l LoyaltyPoints
	test("LoyaltyPoints", &l, t)
}

func TestMerchant(t *testing.T) {
	var m Merchant
	test("Merchant", &m, t)
}

func TestNewReleases(t *testing.T) {
	var n NewReleases
	test("NewReleases", &n, t)
}

func TestNonNegativeIntegerWithUnits(t *testing.T) {
	var i NonNegativeIntegerWithUnits
	test("NonNegativeIntegerWithUnits", &i, t)
}

func TestOffer(t *testing.T) {
	var o Offer
	test("Offer", &o, t)
}

func TestOfferAttributes(t *testing.T) {
	var attribute OfferAttributes
	test("OfferAttributes", &attribute, t)
}

func TestOfferListing(t *testing.T) {
	var listing OfferListing
	test("OfferListing", &listing, t)
}

func TestOffers(t *testing.T) {
	var o Offers
	test("Offers", &o, t)
}

func TestOfferSummary(t *testing.T) {
	var summary OfferSummary
	test("OfferSummary", &summary, t)
}

func TestOperationRequest(t *testing.T) {
	var request OperationRequest
	test("OperationRequest", &request, t)
}

func TestOtherCategoriesSimilarProducts(t *testing.T) {
	var p OtherCategoriesSimilarProducts
	test("OtherCategoriesSimilarProducts", &p, t)
}

func TestPrice(t *testing.T) {
	var p Price
	test("Price", &p, t)
}

func TestPromotion(t *testing.T) {
	var p Promotion
	test("Promotion", &p, t)
}

func TestPromotions(t *testing.T) {
	var p Promotions
	test("Promotions", &p, t)
}

func TestProperty(t *testing.T) {
	var prop Property
	test("Property", &prop, t)
}

func TestRelatedItem(t *testing.T) {
	var item RelatedItem
	test("RelatedItem", &item, t)
}

func TestRelatedItems(t *testing.T) {
	var items RelatedItems
	test("RelatedItems", &items, t)
}

func TestRentalListing(t *testing.T) {
	var listing RentalListing
	test("RentalListing", &listing, t)
}

func TestRentalOffer(t *testing.T) {
	var offer RentalOffer
	test("RentalOffer", &offer, t)
}

func TestRentalOffers(t *testing.T) {
	var offers RentalOffers
	test("RentalOffers", &offers, t)
}

func TestRequest(t *testing.T) {
	var r Request
	test("Request", &r, t)
}

func TestSavedForLaterItems(t *testing.T) {
	var i SavedForLaterItems
	test("SavedForLaterItems", &i, t)
}

func TestSearchBinSet(t *testing.T) {
	var s SearchBinSet
	test("SearchBinSet", &s, t)
}

func TestSearchBinSets(t *testing.T) {
	var s SearchBinSets
	test("SearchBinSets", &s, t)
}

func TestSearchResultsMap(t *testing.T) {
	var result SearchResultsMap
	test("SearchResultsMap", &result, t)
}

func TestSimilarityLookup(t *testing.T) {
	var result SimilarityLookup
	test("SimilarityLookup", &result, t)

	//t.Run("Request", func(t *testing.T) {
	var request SimilarityLookupRequest
	test("SimilarityLookupRequest", &request, t)
	return
	//})

	//t.Run("Response", func(t *testing.T) {
	var response SimilarityLookupResponse
	test("SimilarityLookupResponse", &response, t)
	return
	//})
}

func TestSimilarProducts(t *testing.T) {
	var products SimilarProducts
	test("SimilarProducts", &products, t)
}

func TestStringWithUnits(t *testing.T) {
	var s StringWithUnits
	test("StringWithUnits", &s, t)
}

func TestTopItemSet(t *testing.T) {
	var itemSet TopItemSet
	test("TopItemSet", &itemSet, t)
}

func TestTopSellers(t *testing.T) {
	var sellers TopSellers
	test("TopSellers", &sellers, t)
}

func TestTracks(t *testing.T) {
	var tr Tracks
	test("Tracks", &tr, t)
}

func TestVariationAttribute(t *testing.T) {
	var v VariationAttribute
	test("VariationAttribute", &v, t)
}

func TestVariationDimensions(t *testing.T) {
	var v VariationDimensions
	test("VariationDimensions", &v, t)
}

func TestVariations(t *testing.T) {
	var v Variations
	test("Variations", &v, t)
}

func TestVariationSummary(t *testing.T) {
	var v VariationSummary
	test("VariationSummary", &v, t)
}
