package agozon

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Default User agent.
// User can customize it later.
var UserAgent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/43.0.2357.81 Safari/537.36"

// Get response, from url
func HTTPGet(theURL string) (response []byte, err error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", theURL, nil)
	req.Close = true
	req.Header.Set("User-Agent", UserAgent)

	resp, err := client.Do(req)

	if err != nil {
		return
	}

	defer resp.Body.Close()

	response, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return
}

// type Request from default XML Schema
type Request struct {
	// Exported Fields
	IsValid                 bool                     `xml:",omitempty" json:",omitempty"`
	BrowseNodeLookupRequest *BrowseNodeLookupRequest `xml:",omitempty" json:",omitempty"`
	ItemSearchRequest       *ItemSearchRequest       `xml:",omitempty" json:",omitempty"`
	ItemLookupRequest       *ItemLookupRequest       `xml:",omitempty" json:",omitempty"`
	SimilarityLookupRequest *SimilarityLookupRequest `xml:",omitempty" json:",omitempty"`
	CartGetRequest          *CartGetRequest          `xml:",omitempty" json:",omitempty"`
	CartAddRequest          *CartAddRequest          `xml:",omitempty" json:",omitempty"`
	CartCreateRequest       *CartCreateRequest       `xml:",omitempty" json:",omitempty"`
	CartModifyRequest       *CartModifyRequest       `xml:",omitempty" json:",omitempty"`
	CartClearRequest        *CartClearRequest        `xml:",omitempty" json:",omitempty"`
	Errors                  *Errors                  `xml:"Errors>Error,omitempty" json:",omitempty"`
	URL                     url.URL

	// Unexported fields, for internal use only
	response  string
	locale    string
	secretKey string
}

// Create new Request, with default params
func (r *Request) create(AssociateTag string, AWSAccessKeyId string, SecretAccessKey string) {
	r.secretKey = SecretAccessKey
	// build default query
	q := url.Values{
		"Service":        []string{ServiceName},
		"AWSAccessKeyId": []string{AWSAccessKeyId},
		"AssociateTag":   []string{AssociateTag},
		"Version":        []string{Version},
	}.Encode()

	// add query
	r.URL.RawQuery = q
}

// Set locale information.
func (r *Request) SetLocale(localeName string) (err error) {
	r.locale = localeName
	// base url
	endpoint := locales.Select(localeName)
	u, err := url.Parse(endpoint)
	r.URL = *u
	return
}

// Add single param, to the request.
func (r *Request) AddParam(key string, val string) {
	// add single query to URL
	q := r.URL.Query()
	q.Add(key, val)

	// encode url again
	r.URL.RawQuery = q.Encode()
}

// Add multiple params, to the request.
func (r *Request) AddParams(queries map[string]string) {
	// range the queries
	for key, val := range queries {
		// if value is empty,
		// delete the query from queries.
		if val == "" {
			delete(queries, key)
		}
	}
	// add multiple params, to the query
	q := r.URL.Query()
	for key, val := range queries {
		q.Add(key, val)
	}
	// apply to raw query
	r.URL.RawQuery = q.Encode()
}

// Set Response Group
func (r *Request) SetResponseGroup(responseGroup ...string) {
	q := r.URL.Query()
	q.Set("ResponseGroup", strings.Join(responseGroup, ","))
	r.URL.RawQuery = q.Encode()
}

// Call the operation
func (r *Request) CallOperation(operationName string) (err error) {
	r.AddParam("Timestamp", time.Now().Format(time.RFC3339))
	r.AddParam("Operation", operationName)
	signedURL := new(SignedURL)
	signedURL.Create(r.URL.String(), r.secretKey)

	// get request
	theURL := signedURL.String()

	fmt.Println(theURL)

	// Get response via REST http get call
	response, err := HTTPGet(theURL)
	// if error not null return error
	if err != nil {
		return
	}
	// convert response to string,
	// and add to response
	r.response = string(response)
	return
}

// Get Response from Amazon.
// And convert it to corresponding struct.
func (r *Request) Get(response interface{}) (err error) {
	// unmarshal the response xml, to defined
	// response struct
	err = xml.Unmarshal([]byte(r.response), response)

	// if error return error
	if err != nil {
		return
	}

	// return the struct
	return
}
