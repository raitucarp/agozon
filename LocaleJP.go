package agozon

var LocaleJPMap = map[string]LocaleSearchIndex{
	"All": LocaleJP.All, "Apparel": LocaleJP.Apparel, "Appliances": LocaleJP.Appliances, "Automotive": LocaleJP.Automotive, "Baby": LocaleJP.Baby,
	"Beauty": LocaleJP.Beauty, "Blended": LocaleJP.Blended, "Books": LocaleJP.Books, "Classical": LocaleJP.Classical, "DVD": LocaleJP.DVD, "Electronics": LocaleJP.Electronics,
	"ForeignBooks": LocaleJP.ForeignBooks, "GiftCards": LocaleJP.GiftCards, "Grocery": LocaleJP.Grocery, "HealthPersonalCare": LocaleJP.HealthPersonalCare, "Hobbies": LocaleJP.Hobbies, "HomeImprovement": LocaleJP.HomeImprovement, "Jewelry": LocaleJP.Jewelry, "KindleStore": LocaleJP.KindleStore,
	"Kitchen": LocaleJP.Kitchen, "MP3Downloads": LocaleJP.MP3Downloads, "MobileApps": LocaleJP.MobileApps, "Music": LocaleJP.Music, "MusicTracks": LocaleJP.MusicTracks, "MusicalInstruments": LocaleJP.MusicalInstruments, "OfficeProducts": LocaleJP.OfficeProducts, "PCHardware": LocaleJP.PCHardware,
	"PetSupplies": LocaleJP.PetSupplies, "Shoes": LocaleJP.Shoes, "Software": LocaleJP.Software, "SportingGoods": LocaleJP.SportingGoods, "Toys": LocaleJP.Toys,
	"VHS": LocaleJP.VHS, "Video": LocaleJP.Video, "VideoDownload": LocaleJP.VideoDownload, "VideoGames": LocaleJP.VideoGames, "Watches": LocaleJP.Watches,
}

var LocaleJP = struct {
	All, Apparel, Appliances, Automotive, Baby,
	Beauty, Blended, Books, Classical, DVD, Electronics,
	ForeignBooks, GiftCards, Grocery, HealthPersonalCare, Hobbies, HomeImprovement, Jewelry, KindleStore,
	Kitchen, MP3Downloads, MobileApps, Music, MusicTracks, MusicalInstruments, OfficeProducts, PCHardware,
	PetSupplies, Shoes, Software, SportingGoods, Toys,
	VHS, Video, VideoDownload, VideoGames, Watches LocaleSearchIndex
}{
	Apparel: LocaleSearchIndex{
		BrowseNode:           361299011,
		SortValues:           []string{"-price", "price", "relevancerank", "salesrank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Beauty: LocaleSearchIndex{
		BrowseNode:           0,
		SortValues:           []string{"-price", "price", "relevancerank", "reviewrank", "salesrank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Hobbies: LocaleSearchIndex{
		BrowseNode:           13331821,
		SortValues:           []string{"-mfg-age-min", "-price", "-release-date", "-releasedate", "-titlerank", "mfg-age-min", "price", "release-date", "releasedate", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Availability", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Jewelry: LocaleSearchIndex{
		BrowseNode:           85896051,
		SortValues:           []string{"-price", "price", "relevancerank", "reviewrank", "salesrank"},
		ItemSearchParameters: []string{"Availability", "ItemPage", "Keywords", "MerchantId", "MinPercentageOff", "Sort", "Title"},
	},
	Kitchen: LocaleSearchIndex{
		BrowseNode:           0,
		SortValues:           []string{"-price", "date-desc-rank", "price", "relevancerank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Author", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "Sort", "Title"},
	},
	MusicTracks: LocaleSearchIndex{
		BrowseNode:           562032,
		SortValues:           []string{"-titlerank", "titlerank"},
		ItemSearchParameters: []string{"Author", "Availability", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort"},
	},
	Software: LocaleSearchIndex{
		BrowseNode:           637630,
		SortValues:           []string{"-price", "-release-date", "-releasedate", "-titlerank", "price", "release-date", "releasedate", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Classical: LocaleSearchIndex{
		BrowseNode:           562032,
		SortValues:           []string{"-orig-rel-date", "-price", "-pricerank", "-releasedate", "-titlerank", "orig-rel-date", "price", "pricerank", "releasedate", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Artist", "Availability", "Composer", "Conductor", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Orchestra", "Sort", "Title"},
	},
	MP3Downloads: LocaleSearchIndex{
		BrowseNode:           2129039051,
		SortValues:           []string{"-albumrank", "-artistalbumrank", "-price", "-price-new-bin", "-runtime", "-titlerank", "albumrank", "artistalbumrank", "price", "price-new-bin", "releasedate", "relevancerank", "reviewrank_authority", "runtime", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Availability", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	PCHardware: LocaleSearchIndex{
		BrowseNode:           0,
		SortValues:           []string{"-price", "-price-new-bin", "price", "price-new-bin", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	SportingGoods: LocaleSearchIndex{
		BrowseNode:           14315361,
		SortValues:           []string{"-price", "-release-date", "-titlerank", "price", "release-date", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Availability", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Appliances: LocaleSearchIndex{
		BrowseNode:           2277725051,
		SortValues:           []string{"-price", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Books: LocaleSearchIndex{
		BrowseNode:           465610,
		SortValues:           []string{"-price", "-publication_date", "-titlerank", "-unit-sales", "daterank", "inverse-pricerank", "price", "pricerank", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Author", "Availability", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Power", "Publisher", "Sort", "Title"},
	},
	GiftCards: LocaleSearchIndex{
		BrowseNode:           0,
		SortValues:           []string{"-price", "date-desc-rank", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice"},
	},
	MobileApps: LocaleSearchIndex{
		BrowseNode:           2381131051,
		SortValues:           []string{"-price", "pmrank", "price", "relevancerank", "reviewrank", "reviewrank_authority"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	OfficeProducts: LocaleSearchIndex{
		BrowseNode:           86732051,
		SortValues:           []string{"-price", "price", "relevancerank", "reviewrank", "salesrank"},
		ItemSearchParameters: []string{"Availability", "Brand", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	VHS: LocaleSearchIndex{
		BrowseNode:           2130989051,
		SortValues:           []string{"-orig-rel-date", "-price", "-pricerank", "-releasedate", "-titlerank", "orig-rel-date", "price", "pricerank", "releasedate", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Actor", "Availability", "Director", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Publisher", "Sort", "Title"},
	},
	VideoGames: LocaleSearchIndex{
		BrowseNode:           637872,
		SortValues:           []string{"-price", "-releasedate", "-titlerank", "price", "release-date", "releasedate", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "ReleaseDate", "Sort", "Title"},
	},
	MusicalInstruments: LocaleSearchIndex{
		BrowseNode:           2123630051,
		SortValues:           []string{"-price", "price", "relevancerank", "reviewrank", "salesrank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	PetSupplies: LocaleSearchIndex{
		BrowseNode:           2127213051,
		SortValues:           []string{"-price", "-price-new-bin", "price", "price-new-bin", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Shoes: LocaleSearchIndex{
		BrowseNode:           2016926051,
		SortValues:           []string{"-launch-date", "-price", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Toys: LocaleSearchIndex{
		BrowseNode:           13331821,
		SortValues:           []string{"-price", "-release-date", "-releasedate", "-titlerank", "price", "release-date", "releasedate", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Availability", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Automotive: LocaleSearchIndex{
		BrowseNode:           2017305051,
		SortValues:           []string{"-price", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Blended: LocaleSearchIndex{
		BrowseNode:           0,
		SortValues:           []string{},
		ItemSearchParameters: []string{"Availability", "ItemPage", "Keywords"},
	},
	Electronics: LocaleSearchIndex{
		BrowseNode:           3210991,
		SortValues:           []string{"-price", "-release-date", "-releasedate", "-titlerank", "price", "release-date", "releasedate", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Author", "Availability", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	HealthPersonalCare: LocaleSearchIndex{
		BrowseNode:           161669011,
		SortValues:           []string{"-price", "-release-date", "-releasedate", "-titlerank", "price", "release-date", "releasedate", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	HomeImprovement: LocaleSearchIndex{
		BrowseNode:           0,
		SortValues:           []string{"-price", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Music: LocaleSearchIndex{
		BrowseNode:           562032,
		SortValues:           []string{"-orig-rel-date", "-price", "-pricerank", "-releasedate", "-titlerank", "orig-rel-date", "price", "pricerank", "releasedate", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Artist", "Availability", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "ReleaseDate", "Sort", "Title"},
	},
	Video: LocaleSearchIndex{
		BrowseNode:           561972,
		SortValues:           []string{"-orig-rel-date", "-price", "-pricerank", "-releasedate", "-titlerank", "orig-rel-date", "price", "pricerank", "releasedate", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Actor", "Availability", "Director", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Publisher", "Sort", "Title"},
	},
	Baby: LocaleSearchIndex{
		BrowseNode:           13331821,
		SortValues:           []string{"-price", "price", "psrank", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	DVD: LocaleSearchIndex{
		BrowseNode:           562002,
		SortValues:           []string{"-orig-rel-date", "-price", "-pricerank", "-releasedate", "-titlerank", "orig-rel-date", "price", "pricerank", "releasedate", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Actor", "Availability", "Director", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Publisher", "ReleaseDate", "Sort", "Title"},
	},
	All: LocaleSearchIndex{
		BrowseNode:           0,
		SortValues:           []string{},
		ItemSearchParameters: []string{"Availability", "ItemPage", "Keywords", "MerchantId"},
	},
	ForeignBooks: LocaleSearchIndex{
		BrowseNode:           388316011,
		SortValues:           []string{"-price", "-publication_date", "-titlerank", "-unit-sales", "daterank", "inverse-pricerank", "price", "pricerank", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Author", "Availability", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Power", "Publisher", "Sort", "Title"},
	},
	VideoDownload: LocaleSearchIndex{
		BrowseNode:           0,
		SortValues:           []string{"date-desc-rank", "popularity-rank", "price-asc-rank", "price-desc-rank", "relevancerank", "review-rank"},
		ItemSearchParameters: []string{"Actor", "Availability", "Director", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Publisher", "Sort", "Title"},
	},
	Grocery: LocaleSearchIndex{
		BrowseNode:           57240051,
		SortValues:           []string{"-price", "price", "relevancerank", "reviewrank", "salesrank"},
		ItemSearchParameters: []string{"Availability", "Brand", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	KindleStore: LocaleSearchIndex{
		BrowseNode:           2250739051,
		SortValues:           []string{"-price", "daterank", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Author", "Availability", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Publisher", "Sort", "Title"},
	},
	Watches: LocaleSearchIndex{
		BrowseNode:           331952011,
		SortValues:           []string{"-price", "-titlerank", "price", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Availability", "ItemPage", "Keywords", "MerchantId", "MinPercentageOff", "Sort", "Title"}}}
