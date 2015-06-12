package agozon



type LocaleSearchIndex struct {
	Browsenode int `json:"BrowseNode"`
	Sortvalues []string `json:"SortValues"`
	ItemSearchParameters []string `json:"ItemSearchParameters"`
}

// Locale Information for the US Marketplace
var USMarketplaceSearchIndex = map[string]LocaleSearchIndex{
	"All": {
		Browsenode: 0,
		Sortvalues: []string{},
		ItemSearchParameters: []string{},
	},
}

var BRMarketplaceSearchIndex = map[string]LocaleSearchIndex{}
var CAMarketplaceSearchIndex = map[string]LocaleSearchIndex{}
var CNMarketplaceSearchIndex = map[string]LocaleSearchIndex{}
var DEMarketplaceSearchIndex = map[string]LocaleSearchIndex{}
var ESMarketplaceSearchIndex = map[string]LocaleSearchIndex{}
var FRMarketplaceSearchIndex = map[string]LocaleSearchIndex{}
var INMarketPlaceSearchIndex = map[string]LocaleSearchIndex{}
var ITMarketplaceSearchIndex = map[string]LocaleSearchIndex{}
var JPMarketplaceSearchIndex = map[string]LocaleSearchIndex{}
var UKMarketplaceSearchIndex = map[string]LocaleSearchIndex{}

var MarketPlacesSearchIndexes = map[string]map[string]LocaleSearchIndex{
	"BR": BRMarketplaceSearchIndex,
	"CA": CAMarketplaceSearchIndex,
	"CN": CNMarketplaceSearchIndex,
	"DE": DEMarketplaceSearchIndex,
	"ES": ESMarketplaceSearchIndex,
	"FR": FRMarketplaceSearchIndex,
	"IN": INMarketPlaceSearchIndex,
	"IT": ITMarketplaceSearchIndex,
	"JP": JPMarketplaceSearchIndex,
	"UK": UKMarketplaceSearchIndex,
	"US": USMarketplaceSearchIndex,

}

//var Marketplaces = []Marketplace{}
