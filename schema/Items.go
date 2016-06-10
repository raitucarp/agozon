package schema

import (
	"encoding/xml"
)

// Items schema have no description
type Items struct {
	XMLName              xml.Name          `xml:"Items"`
	Request              *Request          `xml:"Request,omitempty"`
	CorrectedQuery       *CorrectedQuery   `xml:"CorrectedQuery,omitempty"`
	Qid                  string            `xml:"Qid,omitempty"`
	EngineQuery          string            `xml:"EngineQuery,omitempty"`
	TotalResults         *uint64           `xml:"TotalResults,omitempty"`
	TotalPages           *uint64           `xml:"TotalPages,omitempty"`
	MoreSearchResultsURL string            `xml:"MoreSearchResultsUrl,omitempty"`
	SearchResultsMap     *SearchResultsMap `xml:"SearchResultsMap,omitempty"`
	Item                 []*Item           `xml:"Item,omitempty"`
	SearchBinSets        *SearchBinSets    `xml:"SearchBinSets,omitempty"`
}
