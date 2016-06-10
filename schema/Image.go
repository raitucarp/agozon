package schema

// Image is an image associated with a seller's item.
//
// Ancestry: SellerListing/Image
type Image struct {
	URL        string            `xml:"URL,omitempty"`
	Height     *DecimalWithUnits `xml:"Height,omitempty"`
	Width      *DecimalWithUnits `xml:"Width,omitempty"`
	IsVerified string            `xml:"IsVerified,omitempty"`
}
