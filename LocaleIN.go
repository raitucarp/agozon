package agozon

var LocaleIN = struct {
	All, Apparel, Automotive, Baby, Beauty,
	Books, DVD, Electronics, GiftCards, Grocery, HealthPersonalCare,
	HomeGarden, Jewelry, KindleStore, Luggage, Music, MusicalInstruments,
	OfficeProducts, PCHardware, PetSupplies, Shoes, SportingGoods, Toys, VideoGames, Watches LocaleSearchIndex
}{
	Jewelry: &LocaleSearchIndex{
		BrowseNode:           1951049031,
		SortValues:           []string{"-price", "popularity-rank", "price", "relevancerank", "reviewrank", "reviewrank_authority"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Books: &LocaleSearchIndex{
		BrowseNode:           976390031,
		SortValues:           []string{"-price", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Author", "Availability", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Power", "Publisher", "Sort", "Title"},
	},
	Automotive: &LocaleSearchIndex{
		BrowseNode:           0,
		SortValues:           []string{"-price", "-release-date", "popularityrank", "price", "relevancerank", "reviewrank", "salesrank"},
		ItemSearchParameters: []string{"Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Baby: &LocaleSearchIndex{
		BrowseNode:           1571275031,
		SortValues:           []string{"-price", "date-desc-rank", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Beauty: &LocaleSearchIndex{
		BrowseNode:           1355017031,
		SortValues:           []string{"-price", "date-desc-rank", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	GiftCards: &LocaleSearchIndex{
		BrowseNode:           0,
		SortValues:           []string{"-price", "-reviewrank_authority", "date-desc-rank", "price", "relevancerank", "reviewrank", "salesrank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice"},
	},
	HomeGarden: &LocaleSearchIndex{
		BrowseNode:           2454176031,
		SortValues:           []string{"-price", "date-desc-rank", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Luggage: &LocaleSearchIndex{
		BrowseNode:           2454170031,
		SortValues:           []string{"-price", "date-desc-rank", "popularity-rank", "price", "relevancerank", "reviewrank", "reviewrank_authority"},
		ItemSearchParameters: []string{"Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	OfficeProducts: &LocaleSearchIndex{
		BrowseNode:           0,
		SortValues:           []string{"-price", "-release-date", "popularityrank", "price", "relevancerank", "reviewrank"},
		ItemSearchParameters: []string{"Availability", "Brand", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	All: &LocaleSearchIndex{
		BrowseNode:           0,
		SortValues:           []string{},
		ItemSearchParameters: []string{"Availability", "ItemPage", "Keywords", "MerchantId"},
	},
	Toys: &LocaleSearchIndex{
		BrowseNode:           1350381031,
		SortValues:           []string{"-price", "date-desc-rank", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	SportingGoods: &LocaleSearchIndex{
		BrowseNode:           0,
		SortValues:           []string{"-price", "date-desc-rank", "popularity-rank", "price", "relevancerank", "reviewrank_authority"},
		ItemSearchParameters: []string{"Author", "Availability", "ItemPage", "Keywords", "MerchantId"},
	},
	Grocery: &LocaleSearchIndex{
		BrowseNode:           0,
		SortValues:           []string{"-price", "-release-date", "popularityrank", "price", "relevancerank", "reviewrank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Music: &LocaleSearchIndex{
		BrowseNode:           976446031,
		SortValues:           []string{"-price", "-releasedate", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	PetSupplies: &LocaleSearchIndex{
		BrowseNode:           0,
		SortValues:           []string{"-price", "-titlerank", "price", "relevance", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Author", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	VideoGames: &LocaleSearchIndex{
		BrowseNode:           976461031,
		SortValues:           []string{"-price", "-releasedate", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Watches: &LocaleSearchIndex{
		BrowseNode:           1350388031,
		SortValues:           []string{"-price", "popularity-rank", "price", "relevancerank", "reviewrank", "reviewrank_authority"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	DVD: &LocaleSearchIndex{
		BrowseNode:           976417031,
		SortValues:           []string{"-price", "daterank", "inverse-pricerank", "price", "releasedate", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Actor", "AudienceRating", "Availability", "Director", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Publisher", "Sort", "Title"},
	},
	Apparel: &LocaleSearchIndex{
		BrowseNode:           0,
		SortValues:           []string{"-price", "date-desc-rank", "price", "relevancerank", "reviewrank_authority"},
		ItemSearchParameters: []string{"Author", "Availability", "ItemPage", "Keywords", "MerchantId"},
	},
	Electronics: &LocaleSearchIndex{
		BrowseNode:           976420031,
		SortValues:           []string{"-price", "date-desc-rank", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Author", "Availability", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	HealthPersonalCare: &LocaleSearchIndex{
		BrowseNode:           0,
		SortValues:           []string{"-price", "date-desc-rank", "price", "relevancerank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Author", "Availability", "ItemPage", "Keywords", "MerchantId", "Sort", "Title"},
	},
	KindleStore: &LocaleSearchIndex{
		BrowseNode:           0,
		SortValues:           []string{"-price", "-release-date", "popularityrank", "price", "relevancerank", "reviewrank", "salesrank"},
		ItemSearchParameters: []string{"Author", "Availability", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Publisher", "Sort", "Title"},
	},
	MusicalInstruments: &LocaleSearchIndex{
		BrowseNode:           0,
		SortValues:           []string{"-price", "-release-date", "popularityrank", "price", "relevancerank", "reviewrank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	PCHardware: &LocaleSearchIndex{
		BrowseNode:           976393031,
		SortValues:           []string{"-price", "date-desc-rank", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Shoes: &LocaleSearchIndex{
		BrowseNode:           0,
		SortValues:           []string{"-price", "date-desc-rank", "price", "relevancerank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
}
