package agozon

var LocaleUK = struct {
	All, Apparel, Appliances, Automotive, Baby, Beauty,
	Blended, Books, Classical, DVD, Electronics, Grocery, HealthPersonalCare,
	HomeGarden, HomeImprovement, Jewelry, KindleStore, Kitchen, Lighting, Luggage, MP3Downloads,
	MobileApps, Music, MusicTracks, MusicalInstruments, OfficeProducts, OutdoorLiving, Outlet, PCHardware,
	PetSupplies, Shoes, Software, SoftwareVideoGames, SportingGoods, Tools, Toys,
	UnboxVideo, VHS, Video, VideoGames, Watches LocaleSearchIndex
}{
	All: &LocaleSearchIndex{
		BrowseNode:           0,
		SortValues:           []string{},
		ItemSearchParameters: []string{"Availability", "ItemPage", "Keywords", "MerchantId"},
	},
	MusicalInstruments: &LocaleSearchIndex{
		BrowseNode:           340837031,
		SortValues:           []string{"-price", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	MusicTracks: &LocaleSearchIndex{
		BrowseNode:           520920,
		SortValues:           []string{"-titlerank", "titlerank"},
		ItemSearchParameters: []string{"Author", "Availability", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort"},
	},
	Outlet: &LocaleSearchIndex{
		BrowseNode:           245408031,
		SortValues:           []string{},
		ItemSearchParameters: []string{"Availability", "ItemPage", "Keywords", "MerchantId", "MinPercentageOff"},
	},
	UnboxVideo: &LocaleSearchIndex{
		BrowseNode:           0,
		SortValues:           []string{"date-desc-rank", "popularity-rank", "price-asc-rank", "price-desc-rank", "relevancerank", "review-rank"},
		ItemSearchParameters: []string{"Actor", "AudienceRating", "Availability", "Director", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Publisher", "Sort", "Title"},
	},
	VHS: &LocaleSearchIndex{
		BrowseNode:           125556011,
		SortValues:           []string{"-price", "-titlerank", "daterank", "inverse-pricerank", "price", "releasedate", "reviewrank", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Actor", "Availability", "Director", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Publisher", "Sort", "Title"},
	},
	Kitchen: &LocaleSearchIndex{
		BrowseNode:           11052591,
		SortValues:           []string{"-price", "-titlerank", "daterank", "price", "reviewrank", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Lighting: &LocaleSearchIndex{
		BrowseNode:           213078031,
		SortValues:           []string{"-price", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Availability", "Brand", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	PCHardware: &LocaleSearchIndex{
		BrowseNode:           340832031,
		SortValues:           []string{"-price", "launch_date", "price", "psrank", "reviewrank", "reviewrank_authority", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	SoftwareVideoGames: &LocaleSearchIndex{
		BrowseNode:           1025616,
		SortValues:           []string{"-titlerank", "daterank", "inverse-pricerank", "price", "reviewrank", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Author", "Availability", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Music: &LocaleSearchIndex{
		BrowseNode:           520920,
		SortValues:           []string{"-price", "-releasedate", "-titlerank", "inverse-pricerank", "price", "releasedate", "reviewrank", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Artist", "Availability", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	SportingGoods: &LocaleSearchIndex{
		BrowseNode:           319530011,
		SortValues:           []string{"-price", "-titlerank", "price", "reviewrank", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Availability", "ItemPage", "Keywords", "MerchantId", "MinPercentageOff", "Sort", "Title"},
	},

	Baby: &LocaleSearchIndex{
		BrowseNode:           60032031,
		SortValues:           []string{"-price", "price", "relevancerank", "salesrank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Classical: &LocaleSearchIndex{
		BrowseNode:           505510,
		SortValues:           []string{"-price", "-titlerank", "inverse-pricerank", "price", "reviewrank", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Artist", "Availability", "Composer", "Conductor", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Orchestra", "Sort", "Title"},
	},
	Electronics: &LocaleSearchIndex{
		BrowseNode:           560800,
		SortValues:           []string{"-titlerank", "daterank", "inverse-pricerank", "price", "reviewrank", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Author", "Availability", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	HomeGarden: &LocaleSearchIndex{
		BrowseNode:           11052591,
		SortValues:           []string{"-price", "-titlerank", "daterank", "price", "reviewrank", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	HomeImprovement: &LocaleSearchIndex{
		BrowseNode:           79904031,
		SortValues:           []string{"-price", "price", "reviewrank", "salesrank"},
		ItemSearchParameters: []string{"Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	VideoGames: &LocaleSearchIndex{
		BrowseNode:           1025616,
		SortValues:           []string{"-titlerank", "daterank", "inverse-pricerank", "price", "reviewrank", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	OfficeProducts: &LocaleSearchIndex{
		BrowseNode:           560800,
		SortValues:           []string{"-price", "price", "relevancerank", "reviewrank", "salesrank"},
		ItemSearchParameters: []string{"Availability", "Brand", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Apparel: &LocaleSearchIndex{
		BrowseNode:           83451031,
		SortValues:           []string{"-launch-date", "-price", "price", "relevancerank", "reviewrank", "salesrank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Blended: &LocaleSearchIndex{
		BrowseNode:           0,
		SortValues:           []string{},
		ItemSearchParameters: []string{"Availability", "ItemPage", "Keywords"},
	},
	Grocery: &LocaleSearchIndex{
		BrowseNode:           344155031,
		SortValues:           []string{"-price", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	KindleStore: &LocaleSearchIndex{
		BrowseNode:           341677031,
		SortValues:           []string{"-edition-sales-velocity", "-price", "daterank", "price", "relevancerank", "reviewrank", "salesrank"},
		ItemSearchParameters: []string{"Author", "Availability", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Publisher", "Sort", "Title"},
	},
	MP3Downloads: &LocaleSearchIndex{
		BrowseNode:           77925031,
		SortValues:           []string{"-price", "-releasedate", "price", "relevancerank", "reviewrank", "salesrank"},
		ItemSearchParameters: []string{"Author", "Availability", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	MobileApps: &LocaleSearchIndex{
		BrowseNode:           1661658031,
		SortValues:           []string{"-price", "pmrank", "price", "relevancerank", "reviewrank", "reviewrank_authority"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Luggage: &LocaleSearchIndex{
		BrowseNode:           2454167031,
		SortValues:           []string{"-price", "date-desc-rank", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Shoes: &LocaleSearchIndex{
		BrowseNode:           362350011,
		SortValues:           []string{"-launch-date", "-price", "pmrank", "price", "reviewrank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Appliances: &LocaleSearchIndex{
		BrowseNode:           908799031,
		SortValues:           []string{"-price", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Beauty: &LocaleSearchIndex{
		BrowseNode:           66280031,
		SortValues:           []string{"-price", "price", "relevancerank", "reviewrank", "salesrank"},
		ItemSearchParameters: []string{"Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Books: &LocaleSearchIndex{
		BrowseNode:           1025612,
		SortValues:           []string{"-price", "-publication_date", "-titlerank", "-unit-sales", "daterank", "inverse-pricerank", "price", "pricerank", "pubdate", "publication_date", "reviewrank", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Author", "Availability", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Power", "Publisher", "Sort", "Title"},
	},
	DVD: &LocaleSearchIndex{
		BrowseNode:           573406,
		SortValues:           []string{"-price", "-titlerank", "daterank", "inverse-pricerank", "price", "releasedate", "reviewrank", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Actor", "AudienceRating", "Availability", "Director", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Publisher", "Sort", "Title"},
	},
	HealthPersonalCare: &LocaleSearchIndex{
		BrowseNode:           66280031,
		SortValues:           []string{"-price", "-titlerank", "daterank", "price", "releasedate", "reviewrank", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Jewelry: &LocaleSearchIndex{
		BrowseNode:           193717031,
		SortValues:           []string{"-launch-date", "-price", "price", "relevancerank", "reviewrank", "salesrank"},
		ItemSearchParameters: []string{"Availability", "ItemPage", "Keywords", "MerchantId", "MinPercentageOff", "Sort", "Title"},
	},
	Tools: &LocaleSearchIndex{
		BrowseNode:           11052591,
		SortValues:           []string{"-price", "-titlerank", "daterank", "price", "reviewrank", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Actor", "Artist", "AudienceRating", "Author", "Availability", "Brand", "Composer", "Conductor", "Director", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Neighborhood", "Orchestra", "Power", "Publisher", "ReleaseDate", "Sort", "Title"},
	},
	Video: &LocaleSearchIndex{
		BrowseNode:           283926,
		SortValues:           []string{"-price", "-titlerank", "daterank", "inverse-pricerank", "price", "releasedate", "reviewrank", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Actor", "AudienceRating", "Availability", "Director", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Publisher", "Sort", "Title"},
	},
	Watches: &LocaleSearchIndex{
		BrowseNode:           328229011,
		SortValues:           []string{"-launch-date", "-price", "-titlerank", "price", "relevancerank", "reviewrank", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Availability", "ItemPage", "Keywords", "MerchantId", "MinPercentageOff", "Sort", "Title"},
	},
	Automotive: &LocaleSearchIndex{
		BrowseNode:           248878031,
		SortValues:           []string{"-price", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	OutdoorLiving: &LocaleSearchIndex{
		BrowseNode:           11052591,
		SortValues:           []string{"-price", "-titlerank", "daterank", "price", "reviewrank", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Software: &LocaleSearchIndex{
		BrowseNode:           1025614,
		SortValues:           []string{"-titlerank", "daterank", "inverse-pricerank", "price", "reviewrank", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	PetSupplies: &LocaleSearchIndex{
		BrowseNode:           340841031,
		SortValues:           []string{"-price", "-price-new-bin", "price", "price-new-bin", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Toys: &LocaleSearchIndex{
		BrowseNode:           712832,
		SortValues:           []string{"-mfg-age-min", "-price", "mfg-age-min", "price", "salesrank"},
		ItemSearchParameters: []string{"Availability", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"}}}
