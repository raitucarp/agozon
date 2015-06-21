package agozon

var LocaleDEMap = map[string]LocaleSearchIndex{
	"All": LocaleDE.All, "Apparel": LocaleDE.Apparel, "Appliances": LocaleDE.Appliances, "Automotive": LocaleDE.Automotive, "Baby": LocaleDE.Baby, "Beauty": LocaleDE.Beauty,
	"Blended": LocaleDE.Blended, "Books": LocaleDE.Books, "Classical": LocaleDE.Classical, "DVD": LocaleDE.DVD, "Electronics": LocaleDE.Electronics,
	"ForeignBooks": LocaleDE.ForeignBooks, "GiftCards": LocaleDE.GiftCards, "Grocery": LocaleDE.Grocery, "HealthPersonalCare": LocaleDE.HealthPersonalCare,
	"HomeGarden": LocaleDE.HomeGarden, "HomeImprovement": LocaleDE.HomeImprovement, "Jewelry": LocaleDE.Jewelry, "KindleStore": LocaleDE.KindleStore, "Kitchen": LocaleDE.Kitchen,
	"Lighting": LocaleDE.Lighting, "Luggage": LocaleDE.Luggage, "MP3Downloads": LocaleDE.MP3Downloads, "Magazines": LocaleDE.Magazines, "MobileApps": LocaleDE.MobileApps, "Music": LocaleDE.Music,
	"MusicTracks": LocaleDE.MusicTracks, "MusicalInstruments": LocaleDE.MusicalInstruments, "OfficeProducts": LocaleDE.OfficeProducts, "OutdoorLiving": LocaleDE.OutdoorLiving, "Outlet": LocaleDE.Outlet, "PCHardware": LocaleDE.PCHardware,
	"PetSupplies": LocaleDE.PetSupplies, "Photo": LocaleDE.Photo, "Shoes": LocaleDE.Shoes, "Software": LocaleDE.Software, "SoftwareVideoGames": LocaleDE.SoftwareVideoGames, "SportingGoods": LocaleDE.SportingGoods, "Tools": LocaleDE.Tools, "Toys": LocaleDE.Toys,
	"UnboxVideo": LocaleDE.UnboxVideo, "VHS": LocaleDE.VHS, "Video": LocaleDE.Video, "VideoGames": LocaleDE.VideoGames, "Watches": LocaleDE.Watches,
}

var LocaleDE = struct {
	All, Apparel, Appliances, Automotive, Baby, Beauty,
	Blended, Books, Classical, DVD, Electronics,
	ForeignBooks, GiftCards, Grocery, HealthPersonalCare,
	HomeGarden, HomeImprovement, Jewelry, KindleStore, Kitchen,
	Lighting, Luggage, MP3Downloads, Magazines, MobileApps, Music,
	MusicTracks, MusicalInstruments, OfficeProducts, OutdoorLiving, Outlet, PCHardware,
	PetSupplies, Photo, Shoes, Software, SoftwareVideoGames, SportingGoods, Tools, Toys,
	UnboxVideo, VHS, Video, VideoGames, Watches LocaleSearchIndex
}{
	Electronics: &LocaleSearchIndex{
		BrowseNode:           569604,
		SortValues:           []string{"-price", "-titlerank", "price", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Author", "Availability", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Grocery: &LocaleSearchIndex{
		BrowseNode:           344162031,
		SortValues:           []string{"-price", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Kitchen: &LocaleSearchIndex{
		BrowseNode:           3169011,
		SortValues:           []string{"-price", "-titlerank", "price", "relevancerank", "reviewrank", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Photo: &LocaleSearchIndex{
		BrowseNode:           0,
		SortValues:           []string{"-price", "-titlerank", "price", "relevancerank", "reviewrank", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	All: &LocaleSearchIndex{
		BrowseNode:           0,
		SortValues:           []string{},
		ItemSearchParameters: []string{"Availability", "FutureLaunchDate", "ItemPage", "Keywords", "MerchantId"},
	},
	Automotive: &LocaleSearchIndex{
		BrowseNode:           78193031,
		SortValues:           []string{"-price", "price", "relevancerank", "reviewrank", "salesrank"},
		ItemSearchParameters: []string{"Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Beauty: &LocaleSearchIndex{
		BrowseNode:           64257031,
		SortValues:           []string{"-price", "price", "relevancerank", "reviewrank", "salesrank"},
		ItemSearchParameters: []string{"Availability", "Brand", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	DVD: &LocaleSearchIndex{
		BrowseNode:           547664,
		SortValues:           []string{"-price", "-titlerank", "price", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Actor", "AudienceRating", "Availability", "Director", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Publisher", "Sort", "Title"},
	},
	Toys: &LocaleSearchIndex{
		BrowseNode:           12950661,
		SortValues:           []string{"-date", "-price", "-titlerank", "price", "relevancerank", "reviewrank", "salesrank"},
		ItemSearchParameters: []string{"Availability", "ItemPage", "Keywords", "MerchantId", "MinPercentageOff", "Sort", "Title"},
	},
	Video: &LocaleSearchIndex{
		BrowseNode:           547664,
		SortValues:           []string{"-price", "-titlerank", "price", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Actor", "AudienceRating", "Availability", "Director", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Publisher", "Sort", "Title"},
	},
	VideoGames: &LocaleSearchIndex{
		BrowseNode:           541708,
		SortValues:           []string{"-date", "-price", "-titlerank", "price", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	PCHardware: &LocaleSearchIndex{
		BrowseNode:           569604,
		SortValues:           []string{"-price", "launch_date", "price", "psrank", "reviewrank", "reviewrank_authority", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Apparel: &LocaleSearchIndex{
		BrowseNode:           78689031,
		SortValues:           []string{"-price", "price", "relevancerank", "reviewrank", "salesrank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	KindleStore: &LocaleSearchIndex{
		BrowseNode:           530485031,
		SortValues:           []string{"-edition-sales-velocity", "-price", "daterank", "price", "relevancerank", "reviewrank", "salesrank"},
		ItemSearchParameters: []string{"Author", "Availability", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Publisher", "Sort", "Title"},
	},
	MP3Downloads: &LocaleSearchIndex{
		BrowseNode:           180529031,
		SortValues:           []string{"-albumrank", "-artistalbumrank", "-price", "-releasedate", "-runtime", "-titlerank", "albumrank", "artistalbumrank", "price", "relevancerank", "reviewrank", "runtime", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Availability", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	OutdoorLiving: &LocaleSearchIndex{
		BrowseNode:           10925051,
		SortValues:           []string{"-price", "-titlerank", "price", "relevancerank", "reviewrank", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	SportingGoods: &LocaleSearchIndex{
		BrowseNode:           16435121,
		SortValues:           []string{"-price", "-release-date", "-titlerank", "price", "release-date", "relevancerank", "reviewrank", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Author", "Availability", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Publisher", "Sort", "Title"},
	},
	Tools: &LocaleSearchIndex{
		BrowseNode:           235002011,
		SortValues:           []string{"-price", "price", "relevancerank", "reviewrank", "salesrank"},
		ItemSearchParameters: []string{"Actor", "Artist", "AudienceRating", "Author", "Availability", "Brand", "Composer", "Conductor", "Director", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Neighborhood", "Orchestra", "Power", "Publisher", "ReleaseDate", "Sort", "Title"},
	},
	UnboxVideo: &LocaleSearchIndex{
		BrowseNode:           0,
		SortValues:           []string{"date-desc-rank", "popularity-rank", "price-asc-rank", "price-desc-rank", "relevancerank", "review-rank"},
		ItemSearchParameters: []string{"Actor", "AudienceRating", "Availability", "Director", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Publisher", "Sort", "Title"},
	},
	"": &LocaleSearchIndex{
		BrowseNode:           0,
		SortValues:           []string{},
		ItemSearchParameters: []string{},
	},
	ForeignBooks: &LocaleSearchIndex{
		BrowseNode:           54071011,
		SortValues:           []string{"-price", "-pubdate", "-publication_date", "-titlerank", "-unit-sales", "inverse-pricerank", "price", "pricerank", "reviewrank", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Author", "Availability", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Power", "Publisher", "Sort", "Title"},
	},
	HomeGarden: &LocaleSearchIndex{
		BrowseNode:           10925241,
		SortValues:           []string{"-price", "-titlerank", "price", "relevancerank", "reviewrank", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	OfficeProducts: &LocaleSearchIndex{
		BrowseNode:           192417031,
		SortValues:           []string{"-price", "price", "relevancerank", "reviewrank", "salesrank"},
		ItemSearchParameters: []string{"Availability", "Brand", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Shoes: &LocaleSearchIndex{
		BrowseNode:           362995011,
		SortValues:           []string{"-launch-date", "-price", "price", "relevancerank", "reviewrank", "salesrank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	VHS: &LocaleSearchIndex{
		BrowseNode:           0,
		SortValues:           []string{"-price", "-titlerank", "price", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Actor", "AudienceRating", "Availability", "Director", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Publisher", "Sort", "Title"},
	},
	Watches: &LocaleSearchIndex{
		BrowseNode:           193708031,
		SortValues:           []string{"-price", "-titlerank", "price", "relevancerank", "reviewrank", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Availability", "ItemPage", "Keywords", "MerchantId", "MinPercentageOff", "Sort", "Title"},
	},
	Baby: &LocaleSearchIndex{
		BrowseNode:           357577011,
		SortValues:           []string{"-price", "price", "psrank", "relevancerank", "reviewrank", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Blended: &LocaleSearchIndex{
		BrowseNode:           0,
		SortValues:           []string{},
		ItemSearchParameters: []string{"Availability", "ItemPage", "Keywords"},
	},
	Books: &LocaleSearchIndex{
		BrowseNode:           541686,
		SortValues:           []string{"-price", "-pubdate", "-publication_date", "-titlerank", "-unit-sales", "inverse-pricerank", "price", "pricerank", "reviewrank", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Author", "Availability", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Power", "Publisher", "Sort", "Title"},
	},
	Classical: &LocaleSearchIndex{
		BrowseNode:           542676,
		SortValues:           []string{"-price", "-pubdate", "-publication_date", "-releasedate", "-titlerank", "price", "pubdate", "publication_date", "relevancerank", "reviewrank", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Artist", "Availability", "Composer", "Conductor", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Orchestra", "Sort", "Title"},
	},
	HealthPersonalCare: &LocaleSearchIndex{
		BrowseNode:           64257031,
		SortValues:           []string{"-price", "-titlerank", "price", "relevancerank", "reviewrank", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Magazines: &LocaleSearchIndex{
		BrowseNode:           1161660,
		SortValues:           []string{"-price", "-titlerank", "-unit-sales", "price", "relevancerank", "reviewrank", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Author", "Availability", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Publisher", "Sort", "Title"},
	},
	SoftwareVideoGames: &LocaleSearchIndex{
		BrowseNode:           541708,
		SortValues:           []string{"-date", "-price", "-titlerank", "price", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Author", "Availability", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Jewelry: &LocaleSearchIndex{
		BrowseNode:           327473011,
		SortValues:           []string{"-price", "price", "relevancerank", "reviewrank", "salesrank"},
		ItemSearchParameters: []string{"Availability", "ItemPage", "Keywords", "MerchantId", "MinPercentageOff", "Sort", "Title"},
	},
	Music: &LocaleSearchIndex{
		BrowseNode:           542676,
		SortValues:           []string{"-price", "-pubdate", "-publication_date", "-releasedate", "-titlerank", "price", "pubdate", "publication_date", "releasedate", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Artist", "Availability", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	GiftCards: &LocaleSearchIndex{
		BrowseNode:           0,
		SortValues:           []string{"-price", "date-desc-rank", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Artist", "Availability", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice"},
	},
	MusicTracks: &LocaleSearchIndex{
		BrowseNode:           542676,
		SortValues:           []string{"-titlerank", "titlerank"},
		ItemSearchParameters: []string{"Author", "Availability", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort"},
	},
	PetSupplies: &LocaleSearchIndex{
		BrowseNode:           427727031,
		SortValues:           []string{"-price", "-price-new-bin", "price", "price-new-bin", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	MobileApps: &LocaleSearchIndex{
		BrowseNode:           1661650031,
		SortValues:           []string{"-price", "pmrank", "price", "relevancerank", "reviewrank", "reviewrank_authority"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	MusicalInstruments: &LocaleSearchIndex{
		BrowseNode:           340850031,
		SortValues:           []string{"-price", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Outlet: &LocaleSearchIndex{
		BrowseNode:           245405031,
		SortValues:           []string{},
		ItemSearchParameters: []string{"Availability", "ItemPage", "Keywords", "MerchantId", "MinPercentageOff"},
	},
	Software: &LocaleSearchIndex{
		BrowseNode:           542064,
		SortValues:           []string{"-date", "-price", "-titlerank", "price", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Appliances: &LocaleSearchIndex{
		BrowseNode:           931573031,
		SortValues:           []string{"-price", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	HomeImprovement: &LocaleSearchIndex{
		BrowseNode:           0,
		SortValues:           []string{"-price", "price", "relevancerank", "reviewrank", "salesrank"},
		ItemSearchParameters: []string{"Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Lighting: &LocaleSearchIndex{
		BrowseNode:           213084031,
		SortValues:           []string{"-price", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Availability", "Brand", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Luggage: &LocaleSearchIndex{
		BrowseNode:           2454119031,
		SortValues:           []string{"-price", "date-desc-rank", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"}}}
