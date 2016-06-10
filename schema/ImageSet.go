package schema

import (
	"encoding/xml"
)

// ImageSet is set of images.
type ImageSet struct {
	XMLName        xml.Name `xml:"ImageSet"`
	Category       string   `xml:"Category,attr,omitempty"`
	SwatchImage    *Image   `xml:"SwatchImage,omitempty"`
	SmallImage     *Image   `xml:"SmallImage,omitempty"`
	ThumbnailImage *Image   `xml:"ThumbnailImage,omitempty"`
	TinyImage      *Image   `xml:"TinyImage,omitempty"`
	MediumImage    *Image   `xml:"MediumImage,omitempty"`
	LargeImage     *Image   `xml:"LargeImage,omitempty"`
	HiResImage     *Image   `xml:"HiResImage,omitempty"`
}

// ImageSets is bundle of ImageSet
type ImageSets struct {
	ImageSet []*ImageSet `xml:"ImageSet,omitempty"`
}
