package schema

import (
	"encoding/xml"
)

// RelatedItem is Container for an item that is related to the one specified in the ItemLookup request.
//
// Ancestry: RelatedItems
//
// Children: Item, ASIN , ItemAttributes
//
// More details: http://docs.aws.amazon.com/AWSECommerceService/latest/DG/CHAP_response_elements.html
type RelatedItem struct {
	XMLName xml.Name `xml:"RelatedItem"`
	Item    *Item    `xml:"Item,omitempty"`
}

// RelatedItems response group returns items related to an item specified in an ItemLookup request. For example, related items could be all of the Unbox episodes in a TV season that are sold separately, or all of the MP3Download tracks on an MP3 album.
//
// The data returned for RelatedItems is limited to ASINs and ItemAttributes. This remains true even if you add additional response groups, such as Large, that would otherwise return additional data.
//
// The relationship between items is unidirectional. One item is the parent and one item is the child. Items, however, can have multiple children or multiple parents for a given relationship type.
//
// The way in which the items are related is specified by the RelationshipType parameter. This parameter is required when you use the RelatedItems response group. Some values include Episode, Season, Tracks, and Variation. For a list of all relationship types, go to the ItemLookup page.
//
// The relationship type is usually named after the child item in the relationship. For example, an MP3 Track is related to an MP3 album and the type of relationship is Tracks. In this relationship, the album is the parent. If you did an ItemLookup for an MP3 Track and requested RelatedItems using Tracks as the RelationshipType, you would receive the parent album (or albums) for that Track. Conversely, looking up an album using Tracks as the RelationshipType returns the list of Tracks on that album.
//
// Each ItemLookup request can return, at most, 10 related items. To return additional items, use the RelateditemsPage parameter. For example, a value of 2 returns the second set of 10 related items.
//
// More details: http://docs.aws.amazon.com/AWSECommerceService/latest/DG/RG_RelatedItems.html
type RelatedItems struct {
	XMLName              xml.Name       `xml:"RelatedItems"`
	Relationship         string         `xml:"Relationship,omitempty"`
	RelationshipType     string         `xml:"RelationshipType,omitempty"`
	RelatedItemCount     *uint64        `xml:"RelatedItemCount,omitempty"`
	RelatedItemPageCount *uint64        `xml:"RelatedItemPageCount,omitempty"`
	RelatedItemPage      *uint64        `xml:"RelatedItemPage,omitempty"`
	RelatedItem          []*RelatedItem `xml:"RelatedItem,omitempty"`
}
