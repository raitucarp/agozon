package agozon

var LocaleCNMap = map[string]LocaleSearchIndex{
	"All": LocaleCN.All, "Apparel": LocaleCN.Apparel, "Appliances": LocaleCN.Appliances,
	"Automotive": LocaleCN.Automotive, "Baby": LocaleCN.Baby, "Beauty": LocaleCN.Beauty, "Books": LocaleCN.Books,
	"Electronics": LocaleCN.Electronics, "Grocery": LocaleCN.Grocery, "HealthPersonalCare": LocaleCN.HealthPersonalCare,
	"Home": LocaleCN.Home, "HomeImprovement": LocaleCN.HomeImprovement, "Jewelry": LocaleCN.Jewelry, "KindleStore": LocaleCN.KindleStore, "Music": LocaleCN.Music,
	"MusicalInstruments": LocaleCN.MusicalInstruments, "OfficeProducts": LocaleCN.OfficeProducts, "PetSupplies": LocaleCN.PetSupplies, "Photo": LocaleCN.Photo,
	"Shoes": LocaleCN.Shoes, "Software": LocaleCN.Software, "SportingGoods": LocaleCN.SportingGoods, "Toys": LocaleCN.Toys,
	"Video": LocaleCN.Video, "VideoGames": LocaleCN.VideoGames, "Watches": LocaleCN.Watches,
}

var LocaleCN = struct {
	All, Apparel, Appliances,
	Automotive, Baby, Beauty, Books,
	Electronics, Grocery, HealthPersonalCare,
	Home, HomeImprovement, Jewelry, KindleStore, Music,
	MusicalInstruments, OfficeProducts, PetSupplies, Photo,
	Shoes, Software, SportingGoods, Toys,
	Video, VideoGames, Watches LocaleSearchIndex
}{
	HomeImprovement: LocaleSearchIndex{
		BrowseNode:           1952921051,
		SortValues:           []string{"-price", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Toys: LocaleSearchIndex{
		BrowseNode:           647071051,
		SortValues:           []string{"-price", "-release-date", "-releasedate", "-titlerank", "price", "release-date", "releasedate", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Availability", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Automotive: LocaleSearchIndex{
		BrowseNode:           1947900051,
		SortValues:           []string{"-launch-date", "-pct-off", "-price", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Electronics: LocaleSearchIndex{
		BrowseNode:           2016117051,
		SortValues:           []string{"-price", "-release-date", "-releasedate", "-titlerank", "price", "release-date", "releasedate", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Author", "Availability", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	MusicalInstruments: LocaleSearchIndex{
		BrowseNode:           2127219051,
		SortValues:           []string{"-price", "date-desc-rank", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	OfficeProducts: LocaleSearchIndex{
		BrowseNode:           2127222051,
		SortValues:           []string{"-price", "price", "relevancerank", "reviewrank", "salesrank"},
		ItemSearchParameters: []string{"Availability", "Brand", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	VideoGames: LocaleSearchIndex{
		BrowseNode:           897416051,
		SortValues:           []string{"-price", "-release-date", "-releasedate", "-titlerank", "price", "release-date", "releasedate", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "ReleaseDate", "Sort", "Title"},
	},
	Apparel: LocaleSearchIndex{
		BrowseNode:           2016157051,
		SortValues:           []string{"-launch-date", "-pct-off", "-price", "price", "relevancerank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Baby: LocaleSearchIndex{
		BrowseNode:           42693071,
		SortValues:           []string{"-price", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	HealthPersonalCare: LocaleSearchIndex{
		BrowseNode:           852804051,
		SortValues:           []string{"-price", "-release-date", "-releasedate", "-titlerank", "price", "release-date", "releasedate", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Music: LocaleSearchIndex{
		BrowseNode:           754387051,
		SortValues:           []string{"-orig-rel-date", "-price", "-pricerank", "-releasedate", "-titlerank", "orig-rel-date", "price", "pricerank", "releasedate", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Artist", "Availability", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "ReleaseDate", "Sort", "Title"},
	},
	Software: LocaleSearchIndex{
		BrowseNode:           863873051,
		SortValues:           []string{"-price", "-release-date", "-releasedate", "-titlerank", "price", "release-date", "releasedate", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Grocery: LocaleSearchIndex{
		BrowseNode:           2127216051,
		SortValues:           []string{"-price", "price", "relevancerank", "reviewrank", "salesrank"},
		ItemSearchParameters: []string{"Availability", "Brand", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Jewelry: LocaleSearchIndex{
		BrowseNode:           816483051,
		SortValues:           []string{"-price", "price", "relevancerank", "reviewrank", "salesrank"},
		ItemSearchParameters: []string{"Availability", "ItemPage", "Keywords", "MerchantId", "MinPercentageOff", "Sort", "Title"},
	},
	PetSupplies: LocaleSearchIndex{
		BrowseNode:           118864071,
		SortValues:           []string{"-launch-date", "-pct-off", "-price", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	All: LocaleSearchIndex{
		BrowseNode:           0,
		SortValues:           []string{"\u00a0"},
		ItemSearchParameters: []string{"Availability", "ItemPage", "Keywords", "MerchantId"},
	},
	Shoes: LocaleSearchIndex{
		BrowseNode:           2029190051,
		SortValues:           []string{"-launch-date", "-price", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Home: LocaleSearchIndex{
		BrowseNode:           2016127051,
		SortValues:           []string{"-price", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Neighborhood", "Sort", "Title"},
	},
	Watches: LocaleSearchIndex{
		BrowseNode:           1953165051,
		SortValues:           []string{"-price", "-titlerank", "price", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Availability", "ItemPage", "Keywords", "MerchantId", "MinPercentageOff", "Sort", "Title"},
	},
	Photo: LocaleSearchIndex{
		BrowseNode:           755653051,
		SortValues:           []string{"-launch-date", "-pct-off", "-price", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Video: LocaleSearchIndex{
		BrowseNode:           2016137051,
		SortValues:           []string{"-orig-rel-date", "-price", "-pricerank", "-releasedate", "-titlerank", "orig-rel-date", "price", "pricerank", "releasedate", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Actor", "Availability", "Director", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Publisher", "Sort", "Title"},
	},
	Beauty: LocaleSearchIndex{
		BrowseNode:           746777051,
		SortValues:           []string{"-price", "price", "relevancerank", "reviewrank", "salesrank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Books: LocaleSearchIndex{
		BrowseNode:           658391051,
		SortValues:           []string{"-price", "-publication_date", "-titlerank", "-unit-sales", "daterank", "inverse-pricerank", "price", "pricerank", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Author", "Availability", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Power", "Publisher", "Sort", "Title"},
	},
	KindleStore: LocaleSearchIndex{
		BrowseNode:           116088071,
		SortValues:           []string{"-price", "daterank", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Author", "Availability", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Publisher", "Sort", "Title"},
	},
	SportingGoods: LocaleSearchIndex{
		BrowseNode:           836313051,
		SortValues:           []string{"-price", "-release-date", "-titlerank", "price", "release-date", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Availability", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Appliances: LocaleSearchIndex{
		BrowseNode:           80208071,
		SortValues:           []string{"-launch-date", "-pct-off", "-price", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"}}}
