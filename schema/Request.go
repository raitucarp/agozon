package schema

import (
	"encoding/xml"
)

// Request response group returns all of the parameters and their values that were submitted in a request.
// Use this information to debug requests.
// All Product Advertising API operations return this response group by default. There can be up to 10 parameters in each request.
//
// More details:
// 	- http://docs.aws.amazon.com/AWSECommerceService/latest/DG/RG_Request.html
// 	- http://docs.aws.amazon.com/AWSECommerceService/latest/DG/AnatomyOfaRESTRequest.html#GeneralRequestFormat
type Request struct {
	XMLName                 xml.Name                 `xml:"Request"`
	IsValid                 string                   `xml:"IsValid,omitempty"`
	BrowseNodeLookupRequest *BrowseNodeLookupRequest `xml:"BrowseNodeLookupRequest,omitempty"`
	ItemSearchRequest       *ItemSearchRequest       `xml:"ItemSearchRequest,omitempty"`
	ItemLookupRequest       *ItemLookupRequest       `xml:"ItemLookupRequest,omitempty"`
	SimilarityLookupRequest *SimilarityLookupRequest `xml:"SimilarityLookupRequest,omitempty"`
	CartGetRequest          *CartGetRequest          `xml:"CartGetRequest,omitempty"`
	CartAddRequest          *CartAddRequest          `xml:"CartAddRequest,omitempty"`
	CartCreateRequest       *CartCreateRequest       `xml:"CartCreateRequest,omitempty"`
	CartModifyRequest       *CartModifyRequest       `xml:"CartModifyRequest,omitempty"`
	CartClearRequest        *CartClearRequest        `xml:"CartClearRequest,omitempty"`
	Errors                  *Errors                  `xml:"Errors,omitempty"`
}
