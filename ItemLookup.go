package agozon
import (
	"strconv"
	"strings"
	"fmt"
)

type ItemLookupRequest struct {
	Condition             string `xml:",omitempty" json:",omitempty"`
	IdType                string `xml:",omitempty" json:",omitempty"`
	IncludeReviewsSummary bool `xml:",omitempty" json:",omitempty"`
	ItemId                []string `xml:",omitempty" json:",omitempty"`
	MerchantId            string `xml:",omitempty" json:",omitempty"`
	RelatedItemPage       string `xml:",omitempty" json:",omitempty"`
	RelationshipType      string `xml:",omitempty" json:",omitempty"`
	responseGroup         []string `xml:",omitempty" json:",omitempty"`
	SearchIndex           string `xml:",omitempty" json:",omitempty"`
	TruncateReviewsAt     uint64 `xml:",omitempty" json:",omitempty"`
	VariationPage         string `xml:",omitempty" json:",omitempty"`
	locale                string
	do                    func() (ItemLookupResponse, error)
	validate              func() map[string]string
}

// Set ResponseGroup via safe method. It will takes one or more arguments
func (r *ItemLookupRequest) ResponseGroup(responseGroup ...string) {
	validResponseGroup := []string{}
	for _, validRG := range ValidResponseGroup["ItemLookup"] {
		for _, rg := range responseGroup {
			if rg == validRG {
				validResponseGroup = append(validResponseGroup, rg)
			}
		}
	}
	r.responseGroup = validResponseGroup
}

/* ItemSearchResponse */
type ItemLookupResponse struct {
	OperationRequest OperationRequest `xml:"OperationRequest,omitempty" json:",omitempty"`
	Items            Items            `xml:",omitempty" json:",omitempty"`
}

func (r *ItemLookupResponse) GetItems() []Item {
	return r.Items.Item
}

func (r *ItemLookupRequest) Do()  (results ItemLookupResponse, err error) {
	results, err = r.do()
	return
}

func (r *Request) ItemLookup(itemids ...string) *ItemLookupRequest {
	r.create()
	// set params
	req := ItemLookupRequest{}
	req.locale = r.locale

	req.validate = func() map[string]string{
		m := map[string]string{}

		m["Condition"] = req.Condition

		// IdType
		// Type of item identifier used to look up an item.
		// All IdTypes except ASINx require a SearchIndex to be specified.
		if req.SearchIndex == "" {
			req.IdType = "ASIN"
		}

		m["IdType"] = req.IdType

		// When set to true, returns the reviews summary within the Reviews iframe.
		m["IncludeReviewsSummary"] = strings.ToTitle(strconv.FormatBool(req.IncludeReviewsSummary))

		// One or more (up to ten) positive integers that uniquely identify an item.
		// The meaning of the number is specified by IdType.
		// That is, if IdType is ASIN, the ItemId value is an ASIN.
		// If ItemIdis an ASIN, a search index cannot be specified in the request.
		if len(itemids) > 0 {
			req.ItemId = itemids
		}
		m["ItemId"] = strings.Join(req.ItemId, ",")

		// An optional parameter you can use to filter search results
		// and offer listings to only include items sold by Amazon.
		// By default, the API will return items sold by various merchants including Amazon.
		// Enter Amazon if you only want to see items sold by Amazon in the response.
		if req.MerchantId != "Amazon" {
			req.MerchantId = ""
		}
		m["MerchantId"] = req.MerchantId

		// This optional parameter is only valid when the RelatedItems response group is used.
		// Each ItemLookup request can return, at most, ten related items.
		// The RelatedItemPage value specifies the set of ten related items to return.
		// A value of 2, for example, returns the second set of ten related items.
		for _, rg := range req.responseGroup {
			if rg == "RelatedItems" {
				m["RelatedItemPage"] = req.RelatedItemPage
				m["RelationshipType"] = req.RelationshipType
			}
		}

		// The product category to search.
		m["SearchIndex"] = "All"
		if val, ok := LocaleInformation[req.locale]; ok {
			if _, ok := val[req.SearchIndex]; ok {
				m["SearchIndex"] = req.SearchIndex
			}
		}

		// By default, reviews are truncated to 1000 characters within the Reviews iframe.
		// To specify a different length, enter the value.
		// To return complete reviews, specify 0.
		if req.TruncateReviewsAt == 0 {
			m["TruncateReviewsAt"] = ""
		} else {
			m["TruncateReviewsAt"] = strconv.FormatUint(req.TruncateReviewsAt, 10)
		}

		// Retrieves a specific page of variations returned by ItemSearch.
		// By default, ItemSearch returns all variations.
		// Use VariationPage to return a subsection of the response.
		// There are 10 variations per page.
		// To examine offers 11 trough 20, for example, set VariationPage to 2.
		// The total number of pages is returned in the TotalPages element.
		m["VariationPage"] = req.VariationPage

		// validate all params by search index to be more safety use
		validParams := map[string]string{}
		if locale, ok := LocaleInformation[req.locale]; ok {
			searchIndex := locale[m["SearchIndex"]]
			for key, val := range m {
				for _, valParams := range searchIndex.ItemSearchParameters {
					if key == valParams {
						validParams[key] = val
					}
				}
			}
			if req.IdType == "ASIN" {
				validParams["SearchIndex"] = ""
			}
			validParams["ItemId"] = m["ItemId"]
		}

		// Specifies the types of values to return.
		// You can specify multiple response groups in one request by separating them with commas.
		responseGroup := []string{}
		for _, val := range req.responseGroup {
			for _, validRG := range ValidResponseGroup["ItemLookup"] {
				if val == validRG {
					responseGroup = append(responseGroup, val)
				}
			}
		}
		r.SetResponseGroup(responseGroup...)
		return validParams
	}
	// do the request
	req.do = func() (response ItemLookupResponse, err error) {
		// validate the request parameters
		q := req.validate()

		// add queries to api request
		r.AddParams(q)
		// call operation ItemSearch
		err = r.CallOperation("ItemLookup")
		if err != nil {
			return
		}
		// get the response
		err = r.Get(&response)
		fmt.Println("get response error", err)
		// catch the error
		if err != nil {
			return
		}
		// detect error from response
		if response.Items.Request.Errors != nil {
			return response, response.Items.Request.Errors
		}

		// return response and error
		return
	}

	return &req
}
