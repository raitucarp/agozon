package agozon

import (
	"net/url"
	"strconv"
	"strings"
)

/* ItemSearchRequest */
type ItemSearchRequest struct {
	Actor                 string     `xml:",omitempty" json:",omitempty"`
	Artist                string     `xml:",omitempty" json:",omitempty"`
	availability          string     `xml:"Availability,omitempty" json:",omitempty"`
	audienceRating        string     `xml:"AudienceRating,omitempty" json:",omitempty"`
	Author                string     `xml:",omitempty" json:",omitempty"`
	Brand                 string     `xml:",omitempty" json:",omitempty"`
	BrowseNode            int        `xml:",omitempty" json:",omitempty"`
	Composer              string     `xml:",omitempty" json:",omitempty"`
	Condition             string     `xml:",omitempty" json:",omitempty"`
	Conductor             string     `xml:",omitempty" json:",omitempty"`
	Director              string     `xml:",omitempty" json:",omitempty"`
	ItemPage              uint64     `xml:",omitempty" json:",omitempty"`
	keywords              string     `xml:"Keywords,omitempty" json:",omitempty"`
	Manufacturer          string     `xml:",omitempty" json:",omitempty"`
	MaximumPrice          uint64     `xml:",omitempty" json:",omitempty"`
	MerchantId            string     `xml:",omitempty" json:",omitempty"`
	MinimumPrice          uint64     `xml:",omitempty" json:",omitempty"`
	MinPercentageOff      uint64     `xml:",omitempty" json:",omitempty"`
	MusicLabel            string     `xml:",omitempty" json:",omitempty"`
	Orchestra             string     `xml:",omitempty" json:",omitempty"`
	Power                 url.Values `xml:",omitempty" json:",omitempty"`
	Publisher             string     `xml:",omitempty" json:",omitempty"`
	RelatedItemPage       string     `xml:",omitempty" json:",omitempty"`
	RelationshipType      string   `xml:",omitempty" json:",omitempty"`
	responseGroup         []string   `xml:"ResponseGroup,omitempty" json:",omitempty"`
	SearchIndex           string     `xml:",omitempty" json:",omitempty"`
	Sort                  string     `xml:",omitempty" json:",omitempty"`
	Title                 string     `xml:",omitempty" json:",omitempty"`
	ReleaseDate           string     `xml:",omitempty" json:",omitempty"`
	TruncateReviewsAt     uint64     `xml:",omitempty" json:",omitempty"`
	IncludeReviewsSummary bool       `xml:",omitempty" json:",omitempty"`
	VariationPage         string
	locale                string
	do                    func() (ItemSearchResponse, error)
	validate              func() map[string]string
}

// Set audience rating by method.
// This is safe method, with checking valid value per locale
func (r *ItemSearchRequest) AudienceRating(ratings ...string) {
	// loop over audience rating list
	if ratingsVal, ok := AudienceRatingList[r.locale]; ok {
		// create array of audience rating
		audienceRating := []string{}
		// loop over ratings value
		for _, rv := range ratingsVal {
			for _, r := range ratings {
				// if rating value in rating,
				// then append it
				if rv == r {
					audienceRating = append(audienceRating, r)
				}
			}
		}
		// set audiencerating in ItemSearchRequest,
		// with joining audiencerating with comma
		// this is will safe for rest api request.
		r.audienceRating = strings.Join(audienceRating, ",")
	}
}

// Set availability to available, via this safe method.
// It will set availability to "Available", it is the
// only value that valid.
func (r *ItemSearchRequest) OnlyAvailable(yes bool) {
	// if yes is true, set availability to available
	if yes == true {
		r.availability = "Available"
	}
}

// Set ResponseGroup via safe method. It will takes one or more arguments
func (r *ItemSearchRequest) ResponseGroup(responseGroup ...string) {
	validResponseGroup := []string{}
	for _, validRG := range ValidResponseGroup["ItemSearch"] {
		for _, rg := range responseGroup {
			if rg == validRG {
				validResponseGroup = append(validResponseGroup, rg)
			}
		}
	}
	r.responseGroup = validResponseGroup
}

// Set keywords
func (r *ItemSearchRequest) Keywords(keywords string) {
	r.keywords = keywords
}

// Do the search.
func (r *ItemSearchRequest) Do() (results ItemSearchResponse, err error) {
	results, err = r.do()
	return
}

/* ItemSearchResponse */
type ItemSearchResponse struct {
	OperationRequest OperationRequest `xml:"OperationRequest,omitempty" json:",omitempty"`
	Items            Items            `xml:",omitempty" json:",omitempty"`
}

func (response *ItemSearchResponse) GetItems() []Item {
	return response.Items.Item
}

/* ItemSearch operation */
func (r *Request) ItemSearch(keywords ...string) *ItemSearchRequest {
	r.create()
	// set params
	req := ItemSearchRequest{}
	req.locale = r.locale

	// validate item search request,
	// returning maps of string to add as queries
	req.validate = func() map[string]string {
		m := map[string]string{}
		// Name of an actor associated with the item.
		// You can enter all or part of the name.
		m["Actor"] = req.Actor

		//Name of an artist associated with the item.
		// You can enter all or part of the name.
		m["Artist"] = req.Artist

		// Validate AudienceRating by Locale
		// Movie ratings based on MPAA ratings or age,
		// depending upon the locale.
		// You may specify one or more values in a comma-separated list in a REST request or
		// by using multiple elements in a SOAP request.
		m["AudienceRating"] = req.audienceRating

		// Name of an author associated with the item.
		// You can enter all or part of the name.
		m["Author"] = req.Author

		// Enables ItemSearch to return only those items that are available.
		// This parameter must be used in combination with a merchant ID and Condition.
		// For more information, see Availability Parameter, which follows this table.
		// When Availability is set to "Available,"
		// the Condition parameter cannot be set to "New."
		m["Availability"] = req.availability

		// Name of a brand associated with the item. You can enter all or part of the name.
		//Type: String, for example, Tim ex, Seiko, Rolex.
		m["Brand"] = req.Brand

		// Browse nodes are positive integers that identify product categories, for example,
		// Literature & Fiction: (17), Medicine: (13996),
		// Mystery & Thrillers: (18), Nonfiction: (53), Outdoors & Nature: (290060).
		if req.BrowseNode == 0 {
			m["BrowseNode"] = ""
		} else {
			m["BrowseNode"] = strconv.Itoa(req.BrowseNode)
		}

		// Name of an composer associated with the item. You can enter all or part of the name.
		m["Composer"] = req.Composer

		// Use the Condition parameter to filter the offers returned in the product list by condition type.
		// By default, Condition equals "New".
		// If you do not get results, consider changing the value to "All".
		// When the Availability parameter is set to "Available," the Condition parameter cannot be set to "New."
		//
		// ItemSearch returns up to ten search results at a time.
		// When condition equals "All," ItemSearch returns up to three offers per condition (if they exist),
		// for example, three new, three used, three refurbished, and three collectible items.
		// Or, for example, if there are no collectible or refurbished offers,
		// ItemSearch returns three new and three used offers.
		if req.availability == "Available" {
			// cannot set condition to new,
			// while availability is available
			if req.Condition == ConditionNew {
				req.Condition = ""
			}
		} else {
			// if Condition is not All
			// and not empty, and not New
			// then set it to empty, it will be valid request
			if (req.Condition != ConditionNew) &&
				(req.Condition != ConditionAll) &&
				(req.Condition != ConditionCollectible) &&
				(req.Condition != ConditionRefurbished) {
				req.Condition = ""
			}
		}
		m["Condition"] = req.Condition

		// Name of a conductor associated with the item.
		// You can enter all or part of the name.
		m["Conductor"] = req.Conductor

		// Name of a director associated with the item.
		// You can enter all or part of the name.
		m["Director"] = req.Director

		// When set to true, returns the reviews summary within the Reviews iframe.
		if req.IncludeReviewsSummary == false {
			m["IncludeReviewsSummary"] = ""
		} else {
			m["IncludeReviewsSummary"] = "true"
		}

		// Retrieves a specific page of items from all of the items in a response.
		// Up to ten items are returned on a page unless Condition equals "All."
		// In that case, ItemSearch returns additional offers for those items,
		// one offer per condition type (if an offer exists)—for example,
		// one new, one used, one refurbished, and one collectible item.
		//
		// Or, for example, if there are no collectible or refurbished offers,
		// ItemSearch returns one new and one used offer.
		// If you do not include ItemPage in your request, the first page is returned.
		// The total number of pages of items found is returned in the TotalPages response tag.
		// Valid Values: 1 to 10 (1 to 5 when the search index = "All")
		var ItemPage string
		// if SearchIndex is All
		if req.SearchIndex == "All" {
			// ItemPage should not greater than 5
			if req.ItemPage > 5 {
				req.ItemPage = 5
			}
		}
		// if ItemPage more than 10,
		// set it to 10
		if req.ItemPage > 10 {
			req.ItemPage = 10
		}

		// if ItemPage is 0,
		// convert it to empty string
		if req.ItemPage == 0 {
			ItemPage = ""
		} else {
			ItemPage = strconv.FormatUint(req.ItemPage, 10)
		}
		m["ItemPage"] = ItemPage

		// A word or phrase associated with an item.
		// The word or phrase can be in various product fields,
		// including product title, author, artist, description, manufacturer, and so forth.
		// When, for example, the search index equals "MusicTracks",
		// the Keywords parameter enables you to search by song title.
		// If you enter a phrase, the spaces must be URL-encoded as %20.
		if req.keywords == "" {
			m["Keywords"] = strings.Join(keywords, " ")
		} else {
			m["Keywords"] = req.keywords
		}

		// Name of a manufacturer associated with the item. You can enter all or part of the name.
		m["Manufacturer"] = req.Manufacturer

		// specifies the maximum price of the items in the response.
		// Prices are in terms of the lowest currency denomination, for example, pennies. For example,
		// 3241 represents $32.41.
		if req.MaximumPrice == 0 {
			m["MaximumPrice"] = ""
		} else {
			m["MaximumPrice"] = strconv.FormatUint(req.MaximumPrice, 10)
		}

		// An optional parameter you can use to filter search results and offer listings
		// to only include items sold by Amazon.
		// By default, Product Advertising API returns items sold by various merchants including Amazon.
		// Use the Amazon to limit the response to only items sold by Amazon.
		if req.MerchantId != "Amazon" {
			req.MerchantId = ""
		}
		m["MerchantId"] = req.MerchantId

		// Specifies the minimum price of the items to return.
		// Prices are in terms of the lowest currency denomination, for example, pennies,
		// for example, 3241 represents $32.41.
		if req.MinimumPrice == 0 {
			m["MinimumPrice"] = ""
		} else {
			m["MinimumPrice"] = strconv.FormatUint(req.MinimumPrice, 10)
		}

		// Specifies the minimum percentage off for the items to return
		if req.MinPercentageOff == 0 {
			m["MinPercentageOff"] = ""
		} else {
			m["MinPercentageOff"] = strconv.FormatUint(req.MinPercentageOff, 10)
		}
		// Name of an orchestra associated with the item. You can enter all or part of the name.
		m["Orchestra"] = req.Orchestra

		// Performs a book search using a complex query string.
		// Only works when the search index is set equal to "Books."
		m["Power"] = req.Power.Encode()

		// Name of a publisher associated with the item. You can enter all or part of the name.
		m["Publisher"] = req.Publisher

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

		m["SearchIndex"] = "All"
		m["Sort"] = ""
		if val, ok := LocaleInformation[req.locale]; ok {
			if searchIndex, ok := val[req.SearchIndex]; ok {
				m["SearchIndex"] = req.SearchIndex
				for _, v := range searchIndex.SortValues {
					if req.Sort == v {
						m["Sort"] = req.Sort
					}
				}
			}
		}



		// The title associated with the item. You can enter all or part of the title.
		// Title searches are a subset of Keyword searches.
		// If a Title search yields insufficient results, consider using a Keywords search.
		m["Title"] = req.Title

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
			validParams["SearchIndex"] = m["SearchIndex"]
		}

		// Specifies the types of values to return.
		// You can specify multiple response groups in one request by separating them with commas.
		responseGroup := []string{}
		for _, val := range req.responseGroup {
			for _, validRG := range ValidResponseGroup["ItemSearch"] {
				if val == validRG {
					responseGroup = append(responseGroup, val)
				}
			}
		}
		r.SetResponseGroup(responseGroup...)

		return validParams
	}

	// do the request
	req.do = func() (response ItemSearchResponse, err error) {
		// validate the request parameters
		q := req.validate()

		// add queries to api request
		r.AddParams(q)
		// call operation ItemSearch
		err = r.CallOperation("ItemSearch")
		if err != nil {

			return
		}
		// get the response
		err = r.Get(&response)

		// catch the error
		if err != nil {
			return
		}
		// detect error from resposne
		if response.Items.Request.Errors != nil {
			return response, response.Items.Request.Errors
		}

		// return response and error
		return
	}

	// return ItemSearchRequest,
	// as request
	return &req
}
