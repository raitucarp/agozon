package schema

import (
	"encoding/xml"
)

// BrowseNodeLookup enables you to traverse the browse node hierarchy to find a browse node.
//
// Given a browse node ID, BrowseNodeLookup returns the specified browse node’s name, children, and ancestors.
// The names and browse node IDs of the children and ancestor browse nodes are also returned.
//
// As you traverse down the hierarchy, you refine your search and limit the number of items returned.
// For example, you might traverse the following hierarchy: DVD>Used DVDs>Kids and Family,
// to select out of all the DVDs offered by Amazon only those that are appropriate for family viewing.
// Returning the items associated with Kids and Family produces a much more targeted result than a search based at the level of Used DVDs.
//
// Alternatively, by traversing up the browse node tree, you can determine the root category of an item.
// You might do that, for example, to return the top seller of the root product category using the TopSeller response group in an ItemSearch request.
//
// You can use BrowseNodeLookup iteratively to navigate through the browse node hierarchy to reach the node that most appropriately suits your search.
// Then you can use the browse node ID in an ItemSearch request.
// This response would be far more targeted than, for example, searching through all of the browse nodes in a search index.
//
// for more information:
// http://docs.aws.amazon.com/AWSECommerceService/latest/DG/BrowseNodeLookup.html#BrowseNodeLookup-desc
type BrowseNodeLookup struct {
	XMLName           xml.Name                   `xml:"BrowseNodeLookup"`
	MarketplaceDomain string                     `xml:"MarketplaceDomain,omitempty"`
	AWSAccessKeyID    string                     `xml:"AWSAccessKeyId,omitempty"`
	AssociateTag      string                     `xml:"AssociateTag,omitempty"`
	Validate          string                     `xml:"Validate,omitempty"`
	XMLEscaping       string                     `xml:"XMLEscaping,omitempty"`
	Shared            *BrowseNodeLookupRequest   `xml:"Shared,omitempty"`
	Request           []*BrowseNodeLookupRequest `xml:"Request,omitempty"`
}

// BrowseNodeLookupRequest is request parameter for BrowseNodeLookup
//
// read more:
// http://docs.aws.amazon.com/AWSECommerceService/latest/DG/BrowseNodeLookup.html#BrowseNodeLookup-req
type BrowseNodeLookupRequest struct {
	//XMLName       xml.Name `xml:"BrowseNodeLookupRequest"`
	BrowseNodeID  []string `xml:"BrowseNodeId,omitempty"`
	ResponseGroup []string `xml:"ResponseGroup,omitempty"`
}

// BrowseNodeLookupResponse is response for BrowseNodeLookup
//
// read more:
// http://docs.aws.amazon.com/AWSECommerceService/latest/DG/BrowseNodeLookup.html#ResponseTags
type BrowseNodeLookupResponse struct {
	XMLName          xml.Name          `xml:"BrowseNodeLookupResponse"`
	OperationRequest *OperationRequest `xml:"OperationRequest,omitempty"`
	BrowseNodes      *BrowseNodes      `xml:"BrowseNodes,omitempty"`
}
