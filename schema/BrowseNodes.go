package schema

import (
	"encoding/xml"
)

// BrowseNodes response group returns the browse node names and IDs associated with the items returned in the response.
// The response group also returns the names and IDs of the child and parent browse nodes of the items returned in the response.
//
// It's possible for an item to belong to multiple browse nodes, so it's common to see multiple hierarchies of browse nodes for a single item.
//
// Some products, such as parent ASINs, do not return information in the BrowsesNodes response group.
//
// This response group is similar to the BrowseNodeInfo response group.
// The difference is that the BrowseNodes response group is used with operations that are based on item attributes,
// search indices, and lists.
// These operations typically return multiple items.
// BrowseNodeInfo can only be used with BrowseNodeLookup and the search is always keyed on a browse node ID.
//
// Read more: http://docs.aws.amazon.com/AWSECommerceService/latest/DG/RG_BrowseNodes.html
type BrowseNodes struct {
	XMLName    xml.Name      `xml:"BrowseNodes"`
	Request    *Request      `xml:"Request,omitempty"`
	BrowseNode []*BrowseNode `xml:"BrowseNode,omitempty"`
}
