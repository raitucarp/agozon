package agozon

// Locale Information
type LocaleSearchIndex struct {
	BrowseNode           int64    `json:"BrowseNode"`
	SortValues           []string `json:"SortValues"`
	ItemSearchParameters []string `json:"ItemSearchParameters"`
}

var LocaleInformation = map[string]map[string]LocaleSearchIndex{
	"BR": LocaleBRMap,
	"CA": LocaleCAMap,
	"CN": LocaleCNMap,
	"DE": LocaleDEMap,
	"ES": LocaleESMap,
	"FR": LocaleFRMap,
	"IN": LocaleINMap,
	"IT": LocaleITMap,
	"JP": LocaleJPMap,
	"UK": LocaleUKMap,
	"US": LocaleUSMap,
}
