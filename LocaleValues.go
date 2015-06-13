package agozon

// Locale Information
type LocaleSearchIndex struct {
	BrowseNode           int64    `json:"BrowseNode"`
	SortValues           []string `json:"SortValues"`
	ItemSearchParameters []string `json:"ItemSearchParameters"`
}