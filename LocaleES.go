package agozon

var LocaleESMap = map[string]LocaleSearchIndex{
	"All": LocaleES.All, "Apparel": LocaleES.Apparel, "Automotive": LocaleES.Automotive, "Baby": LocaleES.Baby, "Books": LocaleES.Books,
	"DVD": LocaleES.DVD, "Electronics": LocaleES.Electronics, "ForeignBooks": LocaleES.ForeignBooks, "GiftCards": LocaleES.GiftCards, "HealthPersonalCare": LocaleES.HealthPersonalCare,
	"Jewelry": LocaleES.Jewelry, "KindleStore": LocaleES.KindleStore, "Kitchen": LocaleES.Kitchen, "Lighting": LocaleES.Lighting,
	"Luggage": LocaleES.Luggage, "MP3Downloads": LocaleES.MP3Downloads, "MobileApps": LocaleES.MobileApps, "Music": LocaleES.Music,
	"MusicalInstruments": LocaleES.MusicalInstruments, "OfficeProducts": LocaleES.OfficeProducts, "Shoes": LocaleES.Shoes,
	"Software": LocaleES.Software, "SportingGoods": LocaleES.SportingGoods, "Tools": LocaleES.Tools, "Toys": LocaleES.Toys, "VideoGames": LocaleES.VideoGames, "Watches": LocaleES.Watches,
}

var LocaleES = struct {
	All, Apparel, Automotive, Baby, Books,
	DVD, Electronics, ForeignBooks, GiftCards, HealthPersonalCare,
	Jewelry, KindleStore, Kitchen, Lighting,
	Luggage, MP3Downloads, MobileApps, Music,
	MusicalInstruments, OfficeProducts, Shoes,
	Software, SportingGoods, Tools, Toys, VideoGames, Watches LocaleSearchIndex
}{
	Kitchen: &LocaleSearchIndex{
		BrowseNode:           599392031,
		SortValues:           []string{"-price", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Books: &LocaleSearchIndex{
		BrowseNode:           599365031,
		SortValues:           []string{"-price", "-pubdate", "-publication_date", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Author", "Availability", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Power", "Publisher", "Sort", "Title"},
	},
	GiftCards: &LocaleSearchIndex{
		BrowseNode:           0,
		SortValues:           []string{"-price", "date-desc-rank", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Artist", "Availability", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice"},
	},
	Software: &LocaleSearchIndex{
		BrowseNode:           599377031,
		SortValues:           []string{"-price", "-releasedate", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	All: &LocaleSearchIndex{
		BrowseNode:           0,
		SortValues:           []string{},
		ItemSearchParameters: []string{"Availability", "ItemPage", "Keywords", "MerchantId"},
	},
	Automotive: &LocaleSearchIndex{
		BrowseNode:           1951052031,
		SortValues:           []string{"-price", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	MobileApps: &LocaleSearchIndex{
		BrowseNode:           1661651031,
		SortValues:           []string{"-price", "pmrank", "price", "relevancerank", "reviewrank", "reviewrank_authority"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	OfficeProducts: &LocaleSearchIndex{
		BrowseNode:           0,
		SortValues:           []string{"-price", "-release-date", "popularityrank", "price", "relevancerank", "reviewrank"},
		ItemSearchParameters: []string{"Availability", "Brand", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	VideoGames: &LocaleSearchIndex{
		BrowseNode:           599383031,
		SortValues:           []string{"-price", "-releasedate", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	"": &LocaleSearchIndex{
		BrowseNode:           0,
		SortValues:           []string{},
		ItemSearchParameters: []string{},
	},
	HealthPersonalCare: &LocaleSearchIndex{
		BrowseNode:           0,
		SortValues:           []string{"-price", "date-desc-rank", "price", "relevancerank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "Director", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	KindleStore: &LocaleSearchIndex{
		BrowseNode:           0,
		SortValues:           []string{"-edition-sales-velocity", "-price", "daterank", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Author", "Availability", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Publisher", "Sort", "Title"},
	},
	Tools: &LocaleSearchIndex{
		BrowseNode:           2454134031,
		SortValues:           []string{"-price", "date-desc-rank", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Actor", "Artist", "AudienceRating", "Author", "Availability", "Brand", "Composer", "Conductor", "Director", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Neighborhood", "Orchestra", "Power", "Publisher", "ReleaseDate", "Sort", "Title"},
	},
	MP3Downloads: &LocaleSearchIndex{
		BrowseNode:           1748201031,
		SortValues:           []string{"-albumrank", "-artistalbumrank", "-price", "-releasedate", "-runtime", "-titlerank", "albumrank", "artistalbumrank", "price", "relevancerank", "reviewrank", "reviewrank_authority", "runtime", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Availability", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Watches: &LocaleSearchIndex{
		BrowseNode:           599389031,
		SortValues:           []string{"-price", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Availability", "ItemPage", "Keywords", "MerchantId", "MinPercentageOff", "Sort", "Title"},
	},
	SportingGoods: &LocaleSearchIndex{
		BrowseNode:           2665403031,
		SortValues:           []string{"-price", "date-desc-rank", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Availability", "ItemPage", "Keywords", "MerchantId", "MinPercentageOff", "Sort", "Title"},
	},
	Baby: &LocaleSearchIndex{
		BrowseNode:           1703496031,
		SortValues:           []string{"-price", "price", "relevancerank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Jewelry: &LocaleSearchIndex{
		BrowseNode:           2454127031,
		SortValues:           []string{"-price", "date-desc-rank", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Availability", "ItemPage", "Keywords", "MerchantId", "MinPercentageOff", "Sort", "Title"},
	},
	Electronics: &LocaleSearchIndex{
		BrowseNode:           667050031,
		SortValues:           []string{"-price", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Availability", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	ForeignBooks: &LocaleSearchIndex{
		BrowseNode:           599368031,
		SortValues:           []string{"-price", "-pubdate", "-publication_date", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Author", "Availability", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Power", "Publisher", "Sort", "Title"},
	},
	MusicalInstruments: &LocaleSearchIndex{
		BrowseNode:           0,
		SortValues:           []string{"-price", "-release-date", "popularityrank", "price", "relevancerank", "reviewrank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Apparel: &LocaleSearchIndex{
		BrowseNode:           0,
		SortValues:           []string{"-price", "price", "relevancerank", "reviewrank", "salesrank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	DVD: &LocaleSearchIndex{
		BrowseNode:           599380031,
		SortValues:           []string{"-price", "-releasedate", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Actor", "AudienceRating", "Availability", "Director", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Publisher", "Sort", "Title"},
	},
	Music: &LocaleSearchIndex{
		BrowseNode:           599374031,
		SortValues:           []string{"-price", "-releasedate", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Artist", "Availability", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Shoes: &LocaleSearchIndex{
		BrowseNode:           1571263031,
		SortValues:           []string{"-launch-date", "-price", "popularity-rank", "price", "relevancerank", "reviewrank", "reviewrank_authority"},
		ItemSearchParameters: []string{"Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Toys: &LocaleSearchIndex{
		BrowseNode:           599386031,
		SortValues:           []string{"-price", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Availability", "ItemPage", "Keywords", "MerchantId", "MinPercentageOff", "Sort", "Title"},
	},
	Lighting: &LocaleSearchIndex{
		BrowseNode:           0,
		SortValues:           []string{"-price", "-release-date", "popularityrank", "price", "relevancerank", "reviewrank"},
		ItemSearchParameters: []string{"Availability", "Brand", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Luggage: &LocaleSearchIndex{
		BrowseNode:           2454130031,
		SortValues:           []string{"-price", "date-desc-rank", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"}}}
