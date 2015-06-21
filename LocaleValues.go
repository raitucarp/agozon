package agozon

// Locale Information
type LocaleSearchIndex struct {
	BrowseNode           int64    `json:"BrowseNode"`
	SortValues           []string `json:"SortValues"`
	ItemSearchParameters []string `json:"ItemSearchParameters"`
}

var LocaleInformation = []map[string]LocaleSearchIndex{
	LocaleBRMap, LocaleCAMap, LocaleCNMap, LocaleDEMap, LocaleESMap, LocaleFRMap, LocaleINMap,
	LocaleITMap, LocaleJPMap, LocaleUKMap, LocaleUSMap,
}
