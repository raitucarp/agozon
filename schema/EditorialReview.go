package schema

import (
	"encoding/xml"
)

// EditorialReview response group returns Amazon's review of the item, which appears on the Product Detail page for each item in the response.
//
// 	Note
// 	Copyrighted editorial reviews are not returned. For this reason, the reviews returned may be different than those returned by Amazon.com.
// More details: http://docs.aws.amazon.com/AWSECommerceService/latest/DG/RG_EditorialReview.html
type EditorialReview struct {
	XMLName          xml.Name `xml:"EditorialReview"`
	Source           string   `xml:"Source,omitempty"`
	Content          string   `xml:"Content,omitempty"`
	IsLinkSuppressed bool     `xml:"IsLinkSuppressed,omitempty"`
}

// EditorialReviews contains one or more EditorialReview
type EditorialReviews struct {
	XMLName         xml.Name           `xml:"EditorialReviews"`
	EditorialReview []*EditorialReview `xml:"EditorialReview,omitempty"`
}
