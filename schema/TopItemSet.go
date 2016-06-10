package schema

import (
	"encoding/xml"
)

// TopItem is container object for information related to ranked responses, including MostGifted, MostWishedFor, TopSellers, and NewReleases. Information contained includes ASIN, Title, ProductGroup, Author, Artist, and Actor.
//
// Ancestry: MostGifted, MostWishedFor, TopSellers, NewReleases
//
// Children: ASIN, Title, ProductGroup, Actor, Artist, Author
//
// More details: http://docs.aws.amazon.com/AWSECommerceService/latest/DG/CHAP_response_elements.html
type TopItem struct {
	ASIN          string   `xml:"ASIN,omitempty"`
	Title         string   `xml:"Title,omitempty"`
	DetailPageURL string   `xml:"DetailPageURL,omitempty"`
	ProductGroup  string   `xml:"ProductGroup,omitempty"`
	Author        []string `xml:"Author,omitempty"`
	Artist        []string `xml:"Artist,omitempty"`
	Actor         []string `xml:"Actor,omitempty"`
}

// TopItemSet is container for one or more TopItem elements.
//
// More details: http://docs.aws.amazon.com/AWSECommerceService/latest/DG/CHAP_response_elements.html
type TopItemSet struct {
	XMLName xml.Name   `xml:"TopItemSet"`
	Type    string     `xml:"Type,omitempty"`
	TopItem []*TopItem `xml:"TopItem,omitempty"`
}
