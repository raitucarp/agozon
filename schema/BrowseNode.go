package schema

import (
	"encoding/xml"
)

// BrowseNode enables you to search a specified browse node for associated items
//
// Amazon uses a hierarchy of nodes to organize its items for sale.
// Each node represents a collection of items for sale, such as Harry Potter books, not the items themselves.
// Product Advertising API calls the nodes,
// browse nodes because the customer can browse through the nodes to find the collection of items that interests them.
// For example, the customer might be interested in the browse nodes
// Literature & Fiction, Medicine, Mystery & Thrillers:, Nonfiction:, or Outdoors & Nature.
type BrowseNode struct {
	XMLName xml.Name `xml:"BrowseNode"`

	// BrowseNodeId A positive integer that uniquely identifies a parent product category.
	//
	// Ancestry: BrowseNode/Ancestors/BrowseNode BrowseNode/Children/BrowseNode
	ID             string      `xml:"BrowseNodeId,omitempty"`
	Name           string      `xml:"Name,omitempty"`
	IsCategoryRoot bool        `xml:"IsCategoryRoot,omitempty"`
	Properties     *[]Property `xml:"Properties>Property,omitempty"`

	// Nodes that are subsets of the current node.
	Children *[]BrowseNode `xml:"Children>BrowseNode,omitempty"`

	// Nodes that are supersets of the current node.
	Ancestors   *[]BrowseNode `xml:"Ancestors>BrowseNode,omitempty"`
	TopSellers  *TopSellers   `xml:"TopSellers,omitempty"`
	NewReleases *NewReleases  `xml:"NewReleases,omitempty"`
	TopItemSet  *[]TopItemSet `xml:"TopItemSet,omitempty"`
}

// GetProperties get all properties of nodes.
func (b *BrowseNode) GetProperties() (p []Property) {
	for _, v := range *b.Properties {
		p = append(p, v)
	}
	return
}

// GetChildren get all children of nodes.
func (b *BrowseNode) GetChildren() (nodes []BrowseNode) {
	for _, children := range *b.Children {
		nodes = append(nodes, children)
	}
	return
}

// GetAncestors get ancestors of nodes.
func (b *BrowseNode) GetAncestors() (nodes []BrowseNode) {
	for _, ancestors := range *b.Ancestors {
		nodes = append(nodes, ancestors)
	}
	return
}
