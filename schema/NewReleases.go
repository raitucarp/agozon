package schema

import (
	"encoding/xml"
)

// NewRelease is single description of NewReleases response group.
type NewRelease struct {
	ASIN  string `xml:"ASIN,omitempty"`
	Title string `xml:"Title,omitempty"`
}

// NewReleases response group returns the ASIN and title of new releases in a specified browse node.
// This response group works only with BrowseNodeLookup requests, as shown in the following request.
//
// NewReleases response group returns the ASIN and title of newly released items in a specified browse node.
//
// You can return new releases with the NewReleases and CartNewReleases response groups.
// More details: http://docs.aws.amazon.com/AWSECommerceService/latest/DG/NewReleases.html
type NewReleases struct {
	XMLName    xml.Name      `xml:"NewReleases"`
	NewRelease []*NewRelease `xml:"NewRelease,omitempty"`
}
