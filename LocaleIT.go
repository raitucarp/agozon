package agozon

var LocaleITMap = map[string]LocaleSearchIndex{
	"All": LocaleIT.All, "Apparel": LocaleIT.Apparel, "Automotive": LocaleIT.Automotive, "Baby": LocaleIT.Baby,
	"Books": LocaleIT.Books, "DVD": LocaleIT.DVD, "Electronics": LocaleIT.Electronics, "ForeignBooks": LocaleIT.ForeignBooks, "Garden": LocaleIT.Garden,
	"GiftCards": LocaleIT.GiftCards, "Jewelry": LocaleIT.Jewelry, "KindleStore": LocaleIT.KindleStore, "Kitchen": LocaleIT.Kitchen, "Lighting": LocaleIT.Lighting,
	"Luggage": LocaleIT.Luggage, "MP3Downloads": LocaleIT.MP3Downloads, "MobileApps": LocaleIT.MobileApps, "Music": LocaleIT.Music, "MusicalInstruments": LocaleIT.MusicalInstruments, "OfficeProducts": LocaleIT.OfficeProducts, "PCHardware": LocaleIT.PCHardware,
	"Shoes": LocaleIT.Shoes, "Software": LocaleIT.Software, "SportingGoods": LocaleIT.SportingGoods, "Tools": LocaleIT.Tools, "Toys": LocaleIT.Toys, "VideoGames": LocaleIT.VideoGames, "Watches": LocaleIT.Watches,
}

var LocaleIT = struct {
	All, Apparel, Automotive, Baby,
	Books, DVD, Electronics, ForeignBooks, Garden,
	GiftCards, Jewelry, KindleStore, Kitchen, Lighting,
	Luggage, MP3Downloads, MobileApps, Music, MusicalInstruments, OfficeProducts, PCHardware,
	Shoes, Software, SportingGoods, Tools, Toys, VideoGames, Watches LocaleSearchIndex
}{
	Baby: LocaleSearchIndex{
		BrowseNode:           1571287031,
		SortValues:           []string{"-price", "price", "relevancerank", "reviewrank", "salesrank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Lighting: LocaleSearchIndex{
		BrowseNode:           1571293031,
		SortValues:           []string{"-price", "price", "relevancerank", "reviewrank", "salesrank"},
		ItemSearchParameters: []string{"Availability", "Brand", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Watches: LocaleSearchIndex{
		BrowseNode:           524010031,
		SortValues:           []string{"-price", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Availability", "ItemPage", "Keywords", "MerchantId", "MinPercentageOff", "Sort", "Title"},
	},
	All: LocaleSearchIndex{
		BrowseNode:           0,
		SortValues:           []string{},
		ItemSearchParameters: []string{"Availability", "ItemPage", "Keywords", "MerchantId"},
	},
	Books: LocaleSearchIndex{
		BrowseNode:           411664031,
		SortValues:           []string{"-price", "-pubdate", "-publication_date", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Author", "Availability", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Power", "Publisher", "Sort", "Title"},
	},
	Shoes: LocaleSearchIndex{
		BrowseNode:           524007031,
		SortValues:           []string{"-launch-date", "-price", "inverse-pricerank", "price", "pricerank", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Electronics: LocaleSearchIndex{
		BrowseNode:           412610031,
		SortValues:           []string{"-price", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Availability", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Kitchen: LocaleSearchIndex{
		BrowseNode:           524016031,
		SortValues:           []string{"-price", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	MobileApps: LocaleSearchIndex{
		BrowseNode:           1661661031,
		SortValues:           []string{"-price", "pmrank", "price", "relevancerank", "reviewrank", "reviewrank_authority"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Toys: LocaleSearchIndex{
		BrowseNode:           523998031,
		SortValues:           []string{"-price", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Availability", "ItemPage", "Keywords", "MerchantId", "MinPercentageOff", "Sort", "Title"},
	},
	Apparel: LocaleSearchIndex{
		BrowseNode:           0,
		SortValues:           []string{"-price", "price", "relevancerank", "reviewrank", "salesrank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Garden: LocaleSearchIndex{
		BrowseNode:           635017031,
		SortValues:           []string{"-price", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Neighborhood", "Sort", "Title"},
	},
	GiftCards: LocaleSearchIndex{
		BrowseNode:           0,
		SortValues:           []string{"-price", "date-desc-rank", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Artist", "Availability", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice"},
	},
	OfficeProducts: LocaleSearchIndex{
		BrowseNode:           0,
		SortValues:           []string{"-price", "date-desc-rank", "price", "relevancerank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Availability", "Brand", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	PCHardware: LocaleSearchIndex{
		BrowseNode:           425917031,
		SortValues:           []string{"-price", "date-desc-rank", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	DVD: LocaleSearchIndex{
		BrowseNode:           412607031,
		SortValues:           []string{"-price", "-releasedate", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Actor", "AudienceRating", "Availability", "Director", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Publisher", "Sort", "Title"},
	},
	ForeignBooks: LocaleSearchIndex{
		BrowseNode:           433843031,
		SortValues:           []string{"-price", "-pubdate", "-publication_date", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Author", "Availability", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Power", "Publisher", "Sort", "Title"},
	},
	SportingGoods: LocaleSearchIndex{
		BrowseNode:           524013031,
		SortValues:           []string{"-price", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Availability", "ItemPage", "Keywords", "MerchantId", "MinPercentageOff", "Sort", "Title"},
	},
	Automotive: LocaleSearchIndex{
		BrowseNode:           1571281031,
		SortValues:           []string{"-price", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Jewelry: LocaleSearchIndex{
		BrowseNode:           2454164031,
		SortValues:           []string{"-price", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Availability", "ItemPage", "Keywords", "MerchantId", "MinPercentageOff", "Sort", "Title"},
	},
	Music: LocaleSearchIndex{
		BrowseNode:           412601031,
		SortValues:           []string{"-price", "-releasedate", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Artist", "Availability", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	MusicalInstruments: LocaleSearchIndex{
		BrowseNode:           0,
		SortValues:           []string{"-price", "-release-date", "popularityrank", "price", "relevancerank", "reviewrank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	VideoGames: LocaleSearchIndex{
		BrowseNode:           412604031,
		SortValues:           []string{"-price", "-releasedate", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Luggage: LocaleSearchIndex{
		BrowseNode:           2454149031,
		SortValues:           []string{"-price", "date-desc-rank", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	KindleStore: LocaleSearchIndex{
		BrowseNode:           1331141031,
		SortValues:           []string{"-edition-sales-velocity", "-price", "daterank", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Author", "Availability", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Publisher", "Sort", "Title"},
	},
	MP3Downloads: LocaleSearchIndex{
		BrowseNode:           1748204031,
		SortValues:           []string{"-albumrank", "-artistalbumrank", "-price", "-releasedate", "-runtime", "-titlerank", "albumrank", "artistalbumrank", "price", "relevancerank", "reviewrank", "reviewrank_authority", "runtime", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Availability", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Software: LocaleSearchIndex{
		BrowseNode:           412613031,
		SortValues:           []string{"-price", "-releasedate", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Tools: LocaleSearchIndex{
		BrowseNode:           2454161031,
		SortValues:           []string{"-price", "date-desc-rank", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Actor", "Artist", "AudienceRating", "Author", "Availability", "Brand", "Composer", "Conductor", "Director", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Neighborhood", "Orchestra", "Power", "Publisher", "ReleaseDate", "Sort", "Title"}}}
