package schema

import (
	"encoding/xml"
)

// Accessory is single accesory property.
// See Accessories
type Accessory struct {
	ASIN  string `xml:"ASIN,omitempty"`
	Title string `xml:"Title,omitempty"`
}

// Accessories is associated accesories.
//
// Some items have associated accessories. For example, a camera might have a camera case, flash card, and battery.
// Each of these accessories has an item ID, such as an ASIN, as shown in the following response snippet.
// In this example, the main item, B00008OE6I, the camera, returned in the response comes with two accessories, B00003G1RG,
// a compact flash card, and B00004WCCT, a leather camera case.
//
// 		<Item>
// 		  <ASIN>B00008OE6I</ASIN>
// 		    <Accessories>
// 			  <Accessory>
// 			    <ASIN>B00003G1RG</ASIN>
// 				<Title>Viking 128 MB CompactFlash Card (CF128M)</Title>
// 			  </Accessory>
// 			  <Accessory>
// 			    <ASIN>B00004WCCT</ASIN>
// 				<Title>Canon Soft Leather Case for Canon Digital ELPH Cameras(Black)</Title>
// 			  </Accessory>
//
// Read more: http://docs.aws.amazon.com/AWSECommerceService/latest/DG/Accessories.html
type Accessories struct {
	XMLName   xml.Name    `xml:"Accessories"`
	Accessory []Accessory `xml:"Accessory,omitempty"`
}
