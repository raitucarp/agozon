package agozon

import "strings"

var AudienceRatingList = map[string][]string{
	"CA": []string{"G", "PG", "PG-13", "R", "NC-17", "NR", "Unrated", "Family Viewing"},
	"DE": []string{"6", "12", "16"},
	"FR": []string{"PG", "12", "16", "18"},
	"US": []string{"G", "PG", "PG-13", "R", "NC-17", "NR", "Unrated"},
}

// Default service name from Amazon.
const ServiceName = "AWSECommerceService"

// Version of AWSECommerceService.
// Agozon use the latest version.
const Version = "2013-08-01"

// Valid Value for Condition
const (
	ConditionNew         = "New"
	ConditionUsed        = "Used"
	ConditionCollectible = "Collectible"
	ConditionRefurbished = "Refurbished"
	ConditionAll         = "All"
)

const (
	ResponseGroupAccessories       = "Accessories"
	ResponseGroupAlternateVersions = "AlternateVersions"
	ResponseGroupBrowseNodeInfo    = "BrowseNodeInfo"
	ResponseGroupBrowseNodes       = "BrowseNodes"
	ResponseGroupCart              = "Cart"
	ResponseGroupCartNewReleases   = "CartNewReleases"
	ResponseGroupCartTopSellers    = "CartTopSellers"
	ResponseGroupCartSimilarities  = "CartSimilarities"
	ResponseGroupEditorialReview   = "EditorialReview"
	ResponseGroupImages            = "Images"
	ResponseGroupItemAttributes    = "ItemAttributes"
	ResponseGroupItemIds           = "ItemIds"
	ResponseGroupLarge             = "Large"
	ResponseGroupMedium            = "Medium"
	ResponseGroupMostGifted        = "MostGifted"
	ResponseGroupMostWishedFor     = "MostWishedFor"
	ResponseGroupNewReleases       = "NewReleases"
	ResponseGroupOfferFull         = "OfferFull"
	ResponseGroupOfferListings     = "OfferListings"
	ResponseGroupOffers            = "Offers"
	ResponseGroupOfferSummary      = "OfferSummary"
	ResponseGroupPromotionSummary  = "PromotionSummary"
	ResponseGroupRelatedItems      = "RelatedItems"
	ResponseGroupRequest           = "Request"
	ResponseGroupReviews           = "Reviews"
	ResponseGroupSalesRank         = "SalesRank"
	ResponseGroupSearchBins        = "SearchBins"
	ResponseGroupSimilarities      = "Similarities"
	ResponseGroupSmall             = "Small"
	ResponseGroupTopSellers        = "TopSellers"
	ResponseGroupTracks            = "Tracks"
	ResponseGroupVariations        = "Variations"
	ResponseGroupVariationImages   = "VariationImages"
	ResponseGroupVariationMatrix   = "VariationMatrix"
	ResponseGroupVariationOffers   = "VariationOffers"
	ResponseGroupVariationSummary  = "VariationSummary"
)

// valid value for responseGroup
var ValidResponseGroup = map[string][]string{
	"ItemSearch": []string{
		ResponseGroupAccessories, ResponseGroupBrowseNodes,
		ResponseGroupEditorialReview, ResponseGroupItemAttributes,
		ResponseGroupItemIds, ResponseGroupLarge,
		ResponseGroupMedium, ResponseGroupOfferFull,
		ResponseGroupOffers, ResponseGroupOfferSummary, ResponseGroupReviews,
		ResponseGroupRelatedItems, ResponseGroupSearchBins, ResponseGroupSimilarities,
		ResponseGroupSmall, ResponseGroupTracks, ResponseGroupVariations, ResponseGroupVariationSummary,
	},
	/*"Accessories", "AlternateVersions", "BrowseNodeInfo", "BrowseNodes", "Cart", "CartNewReleases",
	"CartTopSellers", "CartSimilarities", "EditorialReview", "Images", "ItemAttributes", "ItemIds",
	"Large", "Medium", "MostGifted", "MostWishedFor", "NewReleases", "OfferFull",
	"OfferListings", "Offers", "OfferSummary", "PromotionSummary", "RelatedItems",
	"Request", "Reviews", "SalesRank", "SearchBins", "Similarities", "Small", "TopSellers",
	"Tracks", "Variations", "VariationImages", "VariationMatrix", "VariationOffers", "VariationSummary",*/
}

type Argument struct {
	Name  string `xml:",attr"`
	Value string `xml:",attr"`
}
type Arguments []Argument

type Error struct {
	Code    string `xml:",omitempty"  json:",omitempty"`
	Message string `xml:",omitempty"  json:",omitempty"`
}

func (e *Error) Error() string {
	return e.Code + " >> " + e.Message
}

type Errors []Error

func (e *Errors) Error() string {
	s := "error found:\n"
	m := []string{}
	for _, error := range *e {
		m = append(m, error.Error())
	}
	return s + strings.Join(m, "\n")
}

type Header struct {
	Name  string `xml:",attr,omitempty"`
	Value string `xml:",attr,omitempty"`
}
type HTTPHeaders []Header

type OperationRequest struct {
	*HTTPHeaders          `xml:"HTTPHeaders>Header,omitempty" json:",omitempty"`
	RequestId             string `xml:",omitempty" json:",omitempty"`
	*Arguments            `xml:"Arguments>Argument,omitempty" json:",omitempty"`
	Errors                *Errors `xml:"Errors>Error,omitempty" json:",omitempty"`
	RequestProcessingTime float64 `xml:",omitempty" json:",omitempty"`
}

type CorrectedQuery struct {
	Keywords string `xml:",attr,omitempty"`
	Message  string `xml:",attr,omitempty"`
}

type Bin struct {
	BinName      string `xml:",omitempty" json:",omitempty"`
	BinItemCount uint   `xml:",omitempty" json:",omitempty"`
	BinParameter []struct {
		Name  string `xml:",omitempty" json:",omitempty"`
		Value string `xml:",omitempty" json:",omitempty"`
	} `xml:",omitempty" json:",omitempty"`
}

type SearchBinSet struct {
	NarrowBy string `xml:",attr"`
	Bin      []Bin  `xml:",omitempty" json:",omitempty"`
}

type SearchBinSets []SearchBinSet

type SearchIndex struct {
	ASIN           []string       `xml:"ASIN,omitempty" json:",omitempty"`
	CorrectedQuery CorrectedQuery `xml:"CorrectedQuery,omitempty" json:",omitempty"`
	IndexName      string         `xml:",omitempty" json:",omitempty"`
	Pages          uint           `xml:",omitempty" json:",omitempty"`
	RelevanceRank  uint           `xml:",omitempty" json:",omitempty"`
	Results        uint           `xml:",omitempty" json:",omitempty"`
}
type SearchResultsMap struct {
	SearchIndex []SearchIndex `xml:",omitempty" json:",omitempty"`
}

type BrowseNodeLookupRequest struct {
	BrowseNodeId  []string `xml:"BrowseNodeId,omitempty"`
	ResponseGroup []string `xml:"ResponseGroup,omitempty"`
}

type CartItem struct {
	ASIN           string      `xml:",omitempty" json:",omitempty"`
	AssociateTag   string      `xml:",omitempty" json:",omitempty"`
	ListItemId     string      `xml:",omitempty" json:",omitempty"`
	OfferListingId string      `xml:",omitempty" json:",omitempty"`
	Quantity       string      `xml:",omitempty" json:",omitempty"`
	MetaData       *[]MetaData `xml:",omitempty" json:",omitempty"`
}
type CartItems []CartItem

type CartCommon struct {
	CartId        string   `xml:",omitempty" json:",omitempty"`
	HMAC          string   `xml:",omitempty" json:",omitempty"`
	MergeCart     string   `xml:",omitempty" json:",omitempty"`
	ResponseGroup []string `xml:",omitempty" json:",omitempty"`
}

/* Operation Request */
type CartAddRequest struct {
	*CartCommon `xml:",omitempty" json:",omitempty"`
	Items       CartItems `xml:",omitempty" json:",omitempty"`
}

type CartClearRequest struct {
	*CartCommon
}

type CartCreateRequest struct {
	*CartCommon
	Items CartItems
}

type CartGetRequest struct {
	*CartCommon
}

type CartModifyRequest struct {
	*CartCommon
	Items CartItems
}

type MetaData struct {
	Key   string
	Value string
}

type ItemLookupRequest struct {
	Condition             string `xml:",omitempty" json:",omitempty"`
	IdType                string `xml:",omitempty" json:",omitempty"`
	IncludeReviewsSummary string
	ItemId                []string
	MerchantId            string
	RelatedItemPage       string
	RelationshipType      []string
	ResponseGroup         string
	SearchIndex           string
	TruncateReviewsAt     string
	VariationPage         string
}

type SimilarityLookupRequest struct {
	Condition      string
	ItemId         []string
	MerchantId     string
	ResponseGroup  []string
	SimilarityType string
}

type Accessory struct {
	ASIN  string `xml:",omitempty" json:",omitempty"`
	Title string `xml:",omitempty" json:",omitempty"`
}
type Accessories []Accessory

type AlternateVersion struct {
	ASIN    string `xml:",omitempty" json:",omitempty"`
	Binding string `xml:",omitempty" json:",omitempty"`
	Title   string `xml:",omitempty" json:",omitempty"`
}
type AlternateVersions []AlternateVersion

type Property struct {
	Name  string `xml:",omitempty" json:",omitempty"`
	Value string `xml:",omitempty" json:",omitempty"`
}
type Properties []Property

type NewRelease struct {
	ASIN  string `xml:",omitempty" json:",omitempty"`
	Title string `xml:",omitempty" json:",omitempty"`
}
type NewReleases []NewRelease

type TopItem struct {
	ASIN          string   `xml:",omitempty" json:",omitempty"`
	Actor         []string `xml:",omitempty" json:",omitempty"`
	Artist        []string `xml:",omitempty" json:",omitempty"`
	Author        []string `xml:",omitempty" json:",omitempty"`
	DetailPageURL string   `xml:",omitempty" json:",omitempty"`
	ProductGroup  string   `xml:",omitempty" json:",omitempty"`
	Title         string   `xml:",omitempty" json:",omitempty"`
}
type TopItemSet []TopItem

type TopSeller struct {
	ASIN  string `xml:",omitempty" json:",omitempty"`
	Title string `xml:",omitempty" json:",omitempty"`
}
type TopSellers []TopSeller

type BrowseNode struct {
	BrowseNodeId   string
	Name           string
	IsCategoryRoot bool          `xml:",omitempty" json:",omitempty"`
	Properties     *Properties   `xml:"Properties>Property,omitempty" json:",omitempty"`
	Children       *[]BrowseNode `xml:"Children>BrowseNode,omitempty" json:",omitempty"`
	Ancestors      *[]BrowseNode `xml:"Ancestors>BrowseNode,omitempty" json:",omitempty"`
	NewReleases    *NewReleases  `xml:"NewReleases>NewRelease,omitempty" json:",omitempty"`
	TopItemSet     *TopItemSet   `xml:"TopItemSet>TopItem,omitempty" json:",omitempty"`
	TopSellers     *TopSellers   `xml:"TopSellers>TopSeller,omitempty" json:",omitempty"`
}
type BrowseNodes []BrowseNode

type CollectionItem struct {
	ASIN  string `xml:",omitempty" json:",omitempty"`
	Title string `xml:",omitempty" json:",omitempty"`
}
type Collection struct {
	CollectionItem   []CollectionItem
	CollectionParent struct {
		ASIN  string `xml:",omitempty" json:",omitempty"`
		Title string `xml:",omitempty" json:",omitempty"`
	}
	CollectionSummary struct {
		HighestListPrice Price `xml:",omitempty" json:",omitempty"`
		HighestSalePrice Price `xml:",omitempty" json:",omitempty"`
		LowestListPrice  Price `xml:",omitempty" json:",omitempty"`
		LowestSalePrice  Price `xml:",omitempty" json:",omitempty"`
	}
}
type Collections []Collection

type Price struct {
	Amount         string `json:",omitempty`
	CurrencyCode   string `json:",omitempty`
	FormattedPrice string `json:",omitempty`
}

type CustomerReviews struct {
	IFrameURL  string `xml:",omitempty" json:",omitempty"`
	HasReviews string `xml:",omitempty" json:",omitempty"`
}

type EditorialReview struct {
	Source           string `xml:",omitempty" json:",omitempty"`
	Content          string `xml:",omitempty" json:",omitempty"`
	IsLinkSuppressed bool
}
type EditorialReviews []EditorialReview

type DecimalWithUnits struct {
	Units string `xml:"Units,attr,omitempty" json:",omitempty"`
	Value string `xml:",chardata" json:",omitempty"`
}
type Image struct {
	IsVerified bool              `xml:",omitempty" json:",omitempty"`
	URL        string            `xml:",omitempty" json:",omitempty"`
	Height     *DecimalWithUnits `xml:",omitempty" json:",omitempty"`
	Width      *DecimalWithUnits `xml:",omitempty" json:",omitempty"`
}

type ImageSet struct {
	SwatchImage    *Image `xml:",omitempty" json:",omitempty"`
	SmallImage     *Image `xml:",omitempty" json:",omitempty"`
	ThumbnailImage *Image `xml:",omitempty" json:",omitempty"`
	TinyImage      *Image `xml:",omitempty" json:",omitempty"`
	MediumImage    *Image `xml:",omitempty" json:",omitempty"`
	LargeImage     *Image `xml:",omitempty" json:",omitempty"`
	HiResImage     *Image `xml:",omitempty" json:",omitempty"`
	Category       string `xml:"Category,attr" json:",omitempty"`
}

type ImageSets []ImageSet

type ItemLink struct {
	Description string `xml:",omitempty" json:",omitempty"`
	URL         string `xml:",omitempty" json:",omitempty"`
}
type ItemLinks []ItemLink

type OfferSummary struct {
	LowestNewPrice         *Price `xml:",omitempty" json:",omitempty"`
	LowestUsedPrice        *Price `xml:",omitempty" json:",omitempty"`
	LowestCollectiblePrice *Price `xml:",omitempty" json:",omitempty"`
	LowestRefurbishedPrice *Price `xml:",omitempty" json:",omitempty"`
	TotalNew               string `xml:",omitempty" json:",omitempty"`
	TotalUsed              string `xml:",omitempty" json:",omitempty"`
	TotalCollectible       string `xml:",omitempty" json:",omitempty"`
	TotalRefurbished       string `xml:",omitempty" json:",omitempty"`
}

type VariationSummary struct {
	HighestPrice     Price `xml:",omitempty" json:",omitempty"`
	HighestSalePrice Price `xml:",omitempty" json:",omitempty"`
	LowestPrice      Price `xml:",omitempty" json:",omitempty"`
	LowestSalePrice  Price `xml:",omitempty" json:",omitempty"`
}

type VariationDimension string
type VariationDimensions []VariationDimension
type Variations struct {
	Item                []Item              `xml:",omitempty" json:",omitempty"`
	TotalVariationPages string              `xml:",omitempty" json:",omitempty"`
	TotalVariations     string              `xml:",omitempty" json:",omitempty"`
	VariationDimensions VariationDimensions `xml:"VariationDimensions>VariationDimension" json:",omitempty"`
}
type VariationAttribute struct {
	Name  string `xml:",omitempty" json:",omitempty"`
	Value string `xml:",omitempty" json:",omitempty"`
}
type VariationAttributes []VariationAttribute

type Track struct {
	Number string `xml:"Number,attr"`
}
type Disc struct {
	Track  []Track
	Number string `xml:"Number,attr"`
}
type Tracks []Disc

type Subject string
type Subjects []Subject

type SimilarProduct struct {
	ASIN  string `xml:",omitempty" json:",omitempty"`
	Title string `xml:",omitempty" json:",omitempty"`
}
type SimilarProducts []SimilarProduct

type Merchant struct {
	Name string `xml:",omitempty" json:",omitempty"`
}

type NonNegativeIntegerWithUnits struct {
	Units string `xml:"Units,attr"`
}
type RentalListing struct {
	Disclaimer          string                      `xml:",omitempty" json:",omitempty"`
	IsFulfilledByAmazon bool                        `xml:",omitempty" json:",omitempty"`
	Period              NonNegativeIntegerWithUnits `xml:",omitempty" json:",omitempty"`
	Price               Price                       `xml:",omitempty" json:",omitempty"`
}
type RentalOffer struct {
	Merchant      Merchant        `xml:",omitempty" json:",omitempty"`
	RentalListing []RentalListing `xml:",omitempty" json:",omitempty"`
}

type RentalOffers []RentalOffer

type RelatedItem struct {
	Item Item `xml:",omitempty" json:",omitempty"`
}
type RelatedItems struct {
	RelatedItem          []RelatedItem
	RelatedItemCount     string `xml:",omitempty" json:",omitempty"`
	RelatedItemPage      string `xml:",omitempty" json:",omitempty"`
	RelatedItemPageCount string `xml:",omitempty" json:",omitempty"`
	Relationship         string `xml:",omitempty" json:",omitempty"`
	RelationshipType     string `xml:",omitempty" json:",omitempty"`
}

type OfferAttributes struct {
	Condition string `xml:",omitempty" json:",omitempty"`
}

type LoyaltyPoints struct {
	Points                 string `xml:",omitempty" json:",omitempty"`
	TypicalRedemptionValue Price  `xml:",omitempty" json:",omitempty"`
}

type OfferListing struct {
	OfferListingId         string `xml:",omitempty" json:",omitempty"`
	PricePerUnit           string `xml:",omitempty" json:",omitempty"`
	Price                  *Price `xml:",omitempty" json:",omitempty"`
	SalePrice              *Price `xml:",omitempty" json:",omitempty"`
	AmountSaved            *Price `xml:",omitempty" json:",omitempty"`
	PercentageSaved        string `xml:",omitempty" json:",omitempty"`
	Availability           string `xml:",omitempty" json:",omitempty"`
	AvailabilityAttributes struct {
		AvailabilityType string `xml:",omitempty" json:",omitempty"`
		IsPreorder       bool   `xml:",omitempty" json:",omitempty"`
		MinimumHours     int    `xml:",omitempty" json:",omitempty"`
		MaximumHours     int    `xml:",omitempty" json:",omitempty"`
	}
	IsEligibleForSuperSaverShipping    bool `xml:",omitempty" json:",omitempty"`
	IsEligibleForPrimeFreeDigitalVideo bool `xml:",omitempty" json:",omitempty"`
	IsEligibleForPrime                 bool `xml:",omitempty" json:",omitempty"`
}

type Promotion struct {
	Summary struct {
		BenefitDescription                string `xml:",omitempty" json:",omitempty"`
		Category                          string `xml:",omitempty" json:",omitempty"`
		EligibilityRequirementDescription string `xml:",omitempty" json:",omitempty"`
		EndDate                           string `xml:",omitempty" json:",omitempty"`
		Message                           string `xml:",omitempty" json:",omitempty"`
		PromotionId                       string `xml:",omitempty" json:",omitempty"`
		StartDate                         string `xml:",omitempty" json:",omitempty"`
		TermsAndConditions                string `xml:",omitempty" json:",omitempty"`
	}
}
type Promotions []Promotion

type Offer struct {
	Merchant        *Merchant        `xml:",omitempty" json:",omitempty"`
	OfferAttributes *OfferAttributes `xml:",omitempty" json:",omitempty"`
	OfferListing    *OfferListing    `xml:",omitempty" json:",omitempty"`
	LoyaltyPoints   *LoyaltyPoints   `xml:",omitempty" json:",omitempty"`
	Promotions      *Promotions      `xml:",omitempty" json:",omitempty"`
}
type Offers struct {
	TotalOffers     int
	TotalOfferPages int
	MoreOffersUrl   string   `xml:",omitempty" json:",omitempty"`
	Offer           *[]Offer `xml:",omitempty" json:",omitempty"`
}

type CatalogNumberListElement string
type Creator struct {
	Role string `xml:"Role,attr"`
	Text string
}

type EANListElement string

type Language struct {
	AudioFormat string `xml:",omitempty" json:",omitempty"`
	Name        string `xml:",omitempty" json:",omitempty"`
	Type        string `xml:",omitempty" json:",omitempty"`
}

type Languages []Language

type UPCListElement string

type ItemAttributes struct {
	Actor                 []string                    `xml:",omitempty" json:",omitempty"`
	Artist                []string                    `xml:",omitempty" json:",omitempty"`
	AspectRatio           string                      `xml:",omitempty" json:",omitempty"`
	AudienceRating        string                      `xml:",omitempty" json:",omitempty"`
	AudioFormat           []string                    `xml:",omitempty" json:",omitempty"`
	Author                []string                    `xml:",omitempty" json:",omitempty"`
	Binding               string                      `xml:",omitempty" json:",omitempty"`
	Brand                 string                      `xml:",omitempty" json:",omitempty"`
	CatalogNumberList     *[]CatalogNumberListElement `xml:"CatalogNumberList>CatalogNumberListElement,omitempty" json:",omitempty"`
	Category              []string                    `xml:",omitempty" json:",omitempty"`
	CEROAgeRating         string                      `xml:",omitempty" json:",omitempty"`
	ClothingSize          string                      `xml:",omitempty" json:",omitempty"`
	Color                 string                      `xml:",omitempty" json:",omitempty"`
	Creator               []Creator                   `xml:",omitempty" json:",omitempty"`
	Department            string                      `xml:",omitempty" json:",omitempty"`
	Director              string                      `xml:",omitempty" json:",omitempty"`
	EAN                   string                      `xml:",omitempty" json:",omitempty"`
	EANList               []EANListElement            `xml:"EANList>EANListElement,omitempty" json:",omitempty"`
	Edition               string                      `xml:",omitempty" json:",omitempty"`
	EISBN                 []string                    `xml:",omitempty" json:",omitempty"`
	EnergyEfficiencyClass string                      `xml:",omitempty" json:",omitempty"`
	EpisodeSequence       string                      `xml:",omitempty" json:",omitempty"`
	ESRBAgeRating         string                      `xml:",omitempty" json:",omitempty"`
	Feature               []string                    `xml:",omitempty" json:",omitempty"`
	Format                []string                    `xml:",omitempty" json:",omitempty"`
	Genre                 string                      `xml:",omitempty" json:",omitempty"`
	HardwarePlatform      string                      `xml:",omitempty" json:",omitempty"`
	HazardousMaterialType string                      `xml:",omitempty" json:",omitempty"`
	ISBN                  string                      `xml:",omitempty" json:",omitempty"`
	IsAdultProduct        bool                        `xml:",omitempty" json:",omitempty"`
	IsAutographed         bool                        `xml:",omitempty" json:",omitempty"`
	IsEligibleForTradeIn  bool                        `xml:",omitempty" json:",omitempty"`
	IsMemorabilia         bool                        `xml:",omitempty" json:",omitempty"`
	IssuesPerYear         string                      `xml:",omitempty" json:",omitempty"`
	ItemDimensions        struct {
		Height *DecimalWithUnits `xml:",omitempty" json:",omitempty"`
		Length *DecimalWithUnits `xml:",omitempty" json:",omitempty"`
		Weight *DecimalWithUnits `xml:",omitempty" json:",omitempty"`
		Width  *DecimalWithUnits `xml:",omitempty" json:",omitempty"`
	} `xml:",omitempty" json:",omitempty"`
	ItemPartNumber                       string            `xml:",omitempty" json:",omitempty"`
	Label                                string            `xml:",omitempty" json:",omitempty"`
	Languages                            *Languages        `xml:"Languages>Language,omitempty" json:",omitempty"`
	LegalDisclaimer                      string            `xml:",omitempty" json:",omitempty"`
	ListPrice                            *Price            `xml:",omitempty" json:",omitempty"`
	MPN                                  string            `xml:",omitempty" json:",omitempty"`
	MagazineType                         string            `xml:",omitempty" json:",omitempty"`
	Manufacturer                         string            `xml:",omitempty" json:",omitempty"`
	ManufacturerMaximumAge               *DecimalWithUnits `xml:",omitempty" json:",omitempty"`
	ManufacturerMinimumAge               *DecimalWithUnits `xml:",omitempty" json:",omitempty"`
	ManufacturerPartsWarrantyDescription string            `xml:",omitempty" json:",omitempty"`
	MediaType                            string            `xml:",omitempty" json:",omitempty"`
	Model                                string            `xml:",omitempty" json:",omitempty"`
	ModelYear                            string            `xml:",omitempty" json:",omitempty"`
	NumberOfDiscs                        int               `xml:",omitempty" json:",omitempty"`
	NumberOfIssues                       int               `xml:",omitempty" json:",omitempty"`
	NumberOfItems                        int               `xml:",omitempty" json:",omitempty"`
	NumberOfPages                        int               `xml:",omitempty" json:",omitempty"`
	NumberOfTracks                       int               `xml:",omitempty" json:",omitempty"`
	OperatingSystem                      string            `xml:",omitempty" json:",omitempty"`
	PackageDimensions                    struct {
		Height *DecimalWithUnits `xml:",omitempty" json:",omitempty"`
		Length *DecimalWithUnits `xml:",omitempty" json:",omitempty"`
		Weight *DecimalWithUnits `xml:",omitempty" json:",omitempty"`
		Width  *DecimalWithUnits `xml:",omitempty" json:",omitempty"`
	} `xml:",omitempty" json:",omitempty"`
	PackageQuantity        string                       `xml:",omitempty" json:",omitempty"`
	PartNumber             string                       `xml:",omitempty" json:",omitempty"`
	PictureFormat          []string                     `xml:",omitempty" json:",omitempty"`
	Platform               []string                     `xml:",omitempty" json:",omitempty"`
	ProductGroup           string                       `xml:",omitempty" json:",omitempty"`
	ProductTypeName        string                       `xml:",omitempty" json:",omitempty"`
	ProductTypeSubcategory string                       `xml:",omitempty" json:",omitempty"`
	PublicationDate        string                       `xml:",omitempty" json:",omitempty"`
	Publisher              string                       `xml:",omitempty" json:",omitempty"`
	RegionCode             string                       `xml:",omitempty" json:",omitempty"`
	ReleaseDate            string                       `xml:",omitempty" json:",omitempty"`
	SeasonSequence         string                       `xml:",omitempty" json:",omitempty"`
	RunningTime            *DecimalWithUnits            `xml:",omitempty" json:",omitempty"`
	SeikodoProductCode     string                       `xml:",omitempty" json:",omitempty"`
	Size                   string                       `xml:",omitempty" json:",omitempty"`
	SKU                    string                       `xml:",omitempty" json:",omitempty"`
	Studio                 string                       `xml:",omitempty" json:",omitempty"`
	SubscriptionLength     *NonNegativeIntegerWithUnits `xml:",omitempty" json:",omitempty"`
	Title                  string                       `xml:",omitempty" json:",omitempty"`
	TrackSequence          string                       `xml:",omitempty" json:",omitempty"`
	TradeInValue           *Price                       `xml:",omitempty" json:",omitempty"`
	UPC                    string                       `xml:",omitempty" json:",omitempty"`
	UPCList                *[]UPCListElement            `xml:"UPCList>UPCListElement,omitempty" json:",omitempty"`
	Warranty               string                       `xml:",omitempty" json:",omitempty"`
	WEEETaxValue           *Price                       `xml:",omitempty" json:",omitempty"`
}

type Item struct {
	ASIN                string              `xml:",omitempty" json:",omitempty"`
	ParentASIN          string              `xml:",omitempty" json:",omitempty"`
	Errors              *Errors             `xml:",omitempty" json:",omitempty"`
	DetailPageURL       string              `xml:",omitempty" json:",omitempty"`
	ItemLinks           *ItemLinks          `xml:"ItemLinks>ItemLink,omitempty" json:",omitempty"`
	SalesRank           string              `xml:",omitempty" json:",omitempty"`
	SmallImage          *Image              `xml:",omitempty" json:",omitempty"`
	MediumImage         *Image              `xml:",omitempty" json:",omitempty"`
	LargeImage          *Image              `xml:",omitempty" json:",omitempty"`
	ImageSets           ImageSets           `xml:"ImageSets>ImageSet,omitempty" json:",omitempty"`
	ItemAttributes      *ItemAttributes     `xml:",omitempty" json:",omitempty"`
	VariationAttributes VariationAttributes `xml:",omitempty" json:",omitempty"`
	RelatedItems        *RelatedItems       `xml:",omitempty" json:",omitempty"`
	Collections         *Collections        `xml:",omitempty" json:",omitempty"`
	Subjects            *Subjects           `xml:"Subjects>Subject,omitempty" json:",omitempty"`
	OfferSummary        *OfferSummary       `xml:",omitempty" json:",omitempty"`
	Offers              *Offers             `xml:",omitempty" json:",omitempty"`
	RentalOffers        *RentalOffers       `xml:",omitempty" json:",omitempty"`
	VariationSummary    *VariationSummary   `xml:",omitempty" json:",omitempty"`
	Variations          *Variations         `xml:",omitempty" json:",omitempty"`
	CustomerReviews     *CustomerReviews    `xml:",omitempty" json:",omitempty"`
	EditorialReviews    *EditorialReviews   `xml:"EditorialReviews>EditorialReview,omitempty" json:",omitempty"`
	SimilarProducts     *SimilarProducts    `xml:"SimilarProducts>SimilarProduct,omitempty" json:",omitempty"`
	Accessories         *Accessories        `xml:"Accessories>Accessory,omitempty" json:",omitempty"`
	Tracks              *Tracks             `xml:"Tracks>Disc,omitempty" json:",omitempty"`
	BrowseNodes         *BrowseNodes        `xml:"BrowseNodes>BrowseNode,omitempty" json:",omitempty"`
	AlternateVersions   *AlternateVersions  `xml:"AlternateVersions>AlternateVersion,omitempty" json:",omitempty"`
}

type Items struct {
	Request              *Request          `xml:",omitempty" json:",omitempty"`
	CorrectedQuery       *CorrectedQuery   `xml:",omitempty" json:",omitempty"`
	Qid                  string            `xml:",omitempty" json:",omitempty"`
	EngineQuery          string            `xml:",omitempty" json:",omitempty"`
	TotalResults         uint              `xml:",omitempty" json:",omitempty"`
	TotalPages           uint              `xml:",omitempty" json:",omitempty"`
	MoreSearchResultsUrl string            `xml:",omitempty" json:",omitempty"`
	SearchResultsMap     *SearchResultsMap `xml:",omitempty" json:",omitempty"`
	Item                 []Item            `xml:",omitempty" json:",omitempty"`
	SearchBinSets        *SearchBinSets    `xml:"SearchBinSets>SearchBinSet" json:",omitempty"`
}
