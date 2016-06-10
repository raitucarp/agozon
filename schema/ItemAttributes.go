package schema

import (
	"encoding/xml"
)

// ItemAttributes response group returns a potentially large number of attributes that describe an item.
// For example, an item in the Camera and Photo search index
// might return the attributes, height, width, weight, title, UPC, price, manufacturer, zoom ratio, number of megapixels, and carrying case.
//
// All search indices can return all item attributes.
// However, the number of item attributes returned varies by ASIN.
// Typically, ASINs within the same search index return the same item attributes.
// For example,
// the item attributes returned for an item in the "Books" search index will be different from those
// returned for an item in the "Camera and Photo" search index.
// However, items within a single search index do not necessarily return the same attributes.
//
// Container for many attributes that describe an item.
//
// Children: 88 elements.
//
// More details:
//
// http://docs.aws.amazon.com/AWSECommerceService/latest/DG/RG_ItemAttributes.html
// http://docs.aws.amazon.com/AWSECommerceService/latest/DG/CHAP_response_elements.html
type ItemAttributes struct {
	XMLName               xml.Name      `xml:"ItemAttributes"`
	Actor                 []string      `xml:"Actor,omitempty"`
	Artist                []string      `xml:"Artist,omitempty"`
	AspectRatio           string        `xml:"AspectRatio,omitempty"`
	AudienceRating        string        `xml:"AudienceRating,omitempty"`
	AudioFormat           []string      `xml:"AudioFormat,omitempty"`
	Author                []string      `xml:"Author,omitempty"`
	Binding               string        `xml:"Binding,omitempty"`
	Brand                 string        `xml:"Brand,omitempty"`
	CatalogNumberList     []string      `xml:"CatalogNumberList>CatalogNumberList,omitempty"`
	Category              []string      `xml:"Category,omitempty"`
	CEROAgeRating         string        `xml:"CEROAgeRating,omitempty"`
	ClothingSize          string        `xml:"ClothingSize,omitempty"`
	Color                 string        `xml:"Color,omitempty"`
	Creator               []ItemCreator `xml:"Creator,omitempty"`
	Department            string        `xml:"Department,omitempty"`
	Director              []string      `xml:"Director,omitempty"`
	EAN                   string        `xml:"EAN,omitempty"`
	EANList               []string      `xml:"EANList>EANList,omitempty"`
	Edition               string        `xml:"Edition,omitempty"`
	EISBN                 []string      `xml:"EISBN,omitempty"`
	EnergyEfficiencyClass string        `xml:"EnergyEfficiencyClass,omitempty"`
	EpisodeSequence       string        `xml:"EpisodeSequence,omitempty"`
	ESRBAgeRating         string        `xml:"ESRBAgeRating,omitempty"`
	Feature               []string      `xml:"Feature,omitempty"`
	Format                []string      `xml:"Format,omitempty"`
	Genre                 string        `xml:"Genre,omitempty"`
	HardwarePlatform      string        `xml:"HardwarePlatform,omitempty"`
	HazardousMaterialType string        `xml:"HazardousMaterialType,omitempty"`
	IsAdultProduct        bool          `xml:"IsAdultProduct,omitempty"`
	IsAutographed         bool          `xml:"IsAutographed,omitempty"`
	ISBN                  string        `xml:"ISBN,omitempty"`
	IsEligibleForTradeIn  bool          `xml:"IsEligibleForTradeIn,omitempty"`
	IsMemorabilia         bool          `xml:"IsMemorabilia,omitempty"`
	IssuesPerYear         string        `xml:"IssuesPerYear,omitempty"`
	Dimensions            struct {
		Height *DecimalWithUnits `xml:"Height,omitempty"`
		Length *DecimalWithUnits `xml:"Length,omitempty"`
		Weight *DecimalWithUnits `xml:"Weight,omitempty"`
		Width  *DecimalWithUnits `xml:"Width,omitempty"`
	} `xml:"ItemDimensions,omitempty"`
	ItemPartNumber                       string            `xml:"ItemPartNumber,omitempty"`
	Label                                string            `xml:"Label,omitempty"`
	Languages                            []*Language       `xml:"Languages>Language,omitempty"`
	LegalDisclaimer                      string            `xml:"LegalDisclaimer,omitempty"`
	ListPrice                            *Price            `xml:"ListPrice,omitempty"`
	MagazineType                         string            `xml:"MagazineType,omitempty"`
	Manufacturer                         string            `xml:"Manufacturer,omitempty"`
	ManufacturerMaximumAge               *DecimalWithUnits `xml:"ManufacturerMaximumAge,omitempty"`
	ManufacturerMinimumAge               *DecimalWithUnits `xml:"ManufacturerMinimumAge,omitempty"`
	ManufacturerPartsWarrantyDescription string            `xml:"ManufacturerPartsWarrantyDescription,omitempty"`
	MediaType                            string            `xml:"MediaType,omitempty"`
	Model                                string            `xml:"Model,omitempty"`
	ModelYear                            *uint64           `xml:"ModelYear,omitempty"`
	MPN                                  string            `xml:"MPN,omitempty"`
	NumberOfDiscs                        *uint64           `xml:"NumberOfDiscs,omitempty"`
	NumberOfIssues                       *uint64           `xml:"NumberOfIssues,omitempty"`
	NumberOfItems                        *uint64           `xml:"NumberOfItems,omitempty"`
	NumberOfPages                        *uint64           `xml:"NumberOfPages,omitempty"`
	NumberOfTracks                       *uint64           `xml:"NumberOfTracks,omitempty"`
	OperatingSystem                      string            `xml:"OperatingSystem,omitempty"`
	PackageDimensions                    struct {
		Height *DecimalWithUnits `xml:"Height,omitempty"`
		Length *DecimalWithUnits `xml:"Length,omitempty"`
		Weight *DecimalWithUnits `xml:"Weight,omitempty"`
		Width  *DecimalWithUnits `xml:"Width,omitempty"`
	} `xml:"PackageDimensions,omitempty"`
	PackageQuantity        *uint64                      `xml:"PackageQuantity,omitempty"`
	PartNumber             string                       `xml:"PartNumber,omitempty"`
	PictureFormat          []string                     `xml:"PictureFormat,omitempty"`
	Platform               []string                     `xml:"Platform,omitempty"`
	ProductGroup           string                       `xml:"ProductGroup,omitempty"`
	ProductTypeName        string                       `xml:"ProductTypeName,omitempty"`
	ProductTypeSubcategory string                       `xml:"ProductTypeSubcategory,omitempty"`
	PublicationDate        string                       `xml:"PublicationDate,omitempty"`
	Publisher              string                       `xml:"Publisher,omitempty"`
	RegionCode             string                       `xml:"RegionCode,omitempty"`
	ReleaseDate            string                       `xml:"ReleaseDate,omitempty"`
	SeasonSequence         string                       `xml:"SeasonSequence,omitempty"`
	RunningTime            *DecimalWithUnits            `xml:"RunningTime,omitempty"`
	SeikodoProductCode     string                       `xml:"SeikodoProductCode,omitempty"`
	Size                   string                       `xml:"Size,omitempty"`
	SKU                    string                       `xml:"SKU,omitempty"`
	Studio                 string                       `xml:"Studio,omitempty"`
	SubscriptionLength     *NonNegativeIntegerWithUnits `xml:"SubscriptionLength,omitempty"`
	Title                  string                       `xml:"Title,omitempty"`
	TrackSequence          string                       `xml:"TrackSequence,omitempty"`
	TradeInValue           *Price                       `xml:"TradeInValue,omitempty"`
	UPC                    string                       `xml:"UPC,omitempty"`
	UPCList                []string                     `xml:"UPCList>UPCList,omitempty"`
	Warranty               string                       `xml:"Warranty,omitempty"`
	WEEETaxValue           *Price                       `xml:"WEEETaxValue,omitempty"`
}

// ItemCreator is creator of an item.
type ItemCreator struct {
	Value string `xml:",chardata"`
	Role  string `xml:"Role,attr,omitempty"`
}

// Language is a language
type Language struct {
	Name        string `xml:"Name,omitempty"`
	Type        string `xml:"Type,omitempty"`
	AudioFormat string `xml:"AudioFormat,omitempty"`
}
