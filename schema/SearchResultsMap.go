package schema

import (
	"encoding/xml"
)

// SearchIndex schema, no description available from Amazon.
type SearchIndex struct {
	IndexName      string          `xml:"IndexName,omitempty"`
	Results        *uint64         `xml:"Results,omitempty"`
	Pages          *uint64         `xml:"Pages,omitempty"`
	CorrectedQuery *CorrectedQuery `xml:"CorrectedQuery,omitempty"`
	RelevanceRank  *uint64         `xml:"RelevanceRank,omitempty"`
	ASIN           []string        `xml:"ASIN,omitempty"`
}

// SearchResultsMap schema, no description available from Amazon.
type SearchResultsMap struct {
	XMLName     xml.Name       `xml:"SearchResultsMap"`
	SearchIndex []*SearchIndex `xml:"SearchIndex,omitempty"`
}
