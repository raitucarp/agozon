package agozon

// LocaleBR map
var LocaleBRMap = map[string]LocaleSearchIndex{
	"All":         LocaleBR.All,
	"Books":       LocaleBR.Books,
	"KindleStore": LocaleBR.KindleStore,
}

// Locale Information for the BR Marketplace
var LocaleBR = struct {
	All, Books, KindleStore LocaleSearchIndex
}{
	All: &LocaleSearchIndex{
		BrowseNode: 0,
		SortValues: []string{},
		ItemSearchParameters: []string{
			"Availability", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinimumPrice",
		},
	},
	Books: &LocaleSearchIndex{
		BrowseNode: 0,
		SortValues: []string{
			"-price", "daterank", "price", "relevancerank", "reviewrank_authority", "salesrank",
		},
		ItemSearchParameters: []string{
			"Author", "Availability", "Condition", "ItemPage",
			"Keywords", "Manufacturer", "MaximumPrice", "MerchantId",
			"MinPercentageOff", "MinimumPrice", "Power", "Publisher",
			"Sort", "Title",
		},
	}, KindleStore: &LocaleSearchIndex{
		BrowseNode: 5308308011,
		SortValues: []string{
			"-price", "daterank", "price", "relevancerank",
			"reviewrank", "reviewrank_authority", "salesrank",
		},
		ItemSearchParameters: []string{
			"Author", "Availability", "ItemPage", "Keywords",
			"MaximumPrice", "MerchantId", "MinPercentageOff",
			"MinimumPrice", "Publisher", "Sort", "Title",
		},
	},
}
