package agozon

var LocaleFR = struct {
	All, Apparel, Appliances, Automotive,
	Baby, Beauty, Blended, Books, Classical, DVD,
	Electronics, ForeignBooks, GiftCards, HealthPersonalCare, HomeImprovement,
	Jewelry, KindleStore, Kitchen, LawnAndGarden, Lighting, Luggage, MP3Downloads,
	MobileApps, Music, MusicTracks, MusicalInstruments, OfficeProducts, PCHardware,
	PetSupplies, Shoes, Software, SoftwareVideoGames, SportingGoods,
	Toys, VHS, Video, VideoGames, Watches LocaleSearchIndex
}{
	MusicalInstruments: &LocaleSearchIndex{
		BrowseNode:           340862031,
		SortValues:           []string{"-price", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	OfficeProducts: &LocaleSearchIndex{
		BrowseNode:           192420031,
		SortValues:           []string{"-price", "price", "relevancerank", "reviewrank", "salesrank"},
		ItemSearchParameters: []string{"Availability", "Brand", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Video: &LocaleSearchIndex{
		BrowseNode:           578608,
		SortValues:           []string{"-price", "-titlerank", "price", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Actor", "AudienceRating", "Availability", "Director", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Publisher", "Sort", "Title"},
	},
	Watches: &LocaleSearchIndex{
		BrowseNode:           60937031,
		SortValues:           []string{"-price", "-titlerank", "price", "relevancerank", "reviewrank", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Availability", "ItemPage", "Keywords", "MerchantId", "MinPercentageOff", "Sort", "Title"},
	},
	Apparel: &LocaleSearchIndex{
		BrowseNode:           340856031,
		SortValues:           []string{"-price", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Jewelry: &LocaleSearchIndex{
		BrowseNode:           193711031,
		SortValues:           []string{"-price", "price", "relevancerank", "reviewrank", "salesrank"},
		ItemSearchParameters: []string{"Availability", "ItemPage", "Keywords", "MerchantId", "MinPercentageOff", "Sort", "Title"},
	},
	Classical: &LocaleSearchIndex{
		BrowseNode:           537366,
		SortValues:           []string{"-price", "-titlerank", "inverse-pricerank", "price", "pricerank", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Artist", "Availability", "Composer", "Conductor", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Orchestra", "Sort", "Title"},
	},
	ForeignBooks: &LocaleSearchIndex{
		BrowseNode:           69633011,
		SortValues:           []string{"-daterank", "-price", "-titlerank", "-unit-sales", "inverse-pricerank", "price", "pricerank", "publication_date", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Author", "Availability", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Power", "Publisher", "Sort", "Title"},
	},
	Music: &LocaleSearchIndex{
		BrowseNode:           537366,
		SortValues:           []string{"-price", "-pricerank", "-releasedate", "-titlerank", "availability", "price", "pricerank", "releasedate", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Artist", "Availability", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	"": &LocaleSearchIndex{
		BrowseNode:           0,
		SortValues:           []string{},
		ItemSearchParameters: []string{},
	},
	All: &LocaleSearchIndex{
		BrowseNode:           0,
		SortValues:           []string{},
		ItemSearchParameters: []string{"Availability", "ItemPage", "Keywords", "MerchantId"},
	},
	MobileApps: &LocaleSearchIndex{
		BrowseNode:           1661655031,
		SortValues:           []string{"-price", "pmrank", "price", "relevancerank", "reviewrank", "reviewrank_authority"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Shoes: &LocaleSearchIndex{
		BrowseNode:           248812031,
		SortValues:           []string{"-price", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Baby: &LocaleSearchIndex{
		BrowseNode:           206618031,
		SortValues:           []string{"-price", "price", "relevancerank", "salesrank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	LawnAndGarden: &LocaleSearchIndex{
		BrowseNode:           0,
		SortValues:           []string{"-price", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Author", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	HomeImprovement: &LocaleSearchIndex{
		BrowseNode:           590749031,
		SortValues:           []string{"-price", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	KindleStore: &LocaleSearchIndex{
		BrowseNode:           0,
		SortValues:           []string{"-edition-sales-velocity", "-price", "daterank", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Author", "Availability", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Publisher", "Sort", "Title"},
	},
	Lighting: &LocaleSearchIndex{
		BrowseNode:           213081031,
		SortValues:           []string{"-price", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Availability", "Brand", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Luggage: &LocaleSearchIndex{
		BrowseNode:           2454146031,
		SortValues:           []string{"-price", "date-desc-rank", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	PCHardware: &LocaleSearchIndex{
		BrowseNode:           0,
		SortValues:           []string{"-price", "launch_date", "price", "psrank", "reviewrank", "reviewrank_authority", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Blended: &LocaleSearchIndex{
		BrowseNode:           0,
		SortValues:           []string{},
		ItemSearchParameters: []string{"Availability", "ItemPage", "Keywords"},
	},
	DVD: &LocaleSearchIndex{
		BrowseNode:           578608,
		SortValues:           []string{"-titlerank", "amzrank", "availability", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Actor", "AudienceRating", "Availability", "Director", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Publisher", "Sort", "Title"},
	},
	Toys: &LocaleSearchIndex{
		BrowseNode:           548014,
		SortValues:           []string{"-price", "-titlerank", "price", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Availability", "ItemPage", "Keywords", "MerchantId", "MinPercentageOff", "Sort", "Title"},
	},
	Books: &LocaleSearchIndex{
		BrowseNode:           468256,
		SortValues:           []string{"-daterank", "-price", "-titlerank", "-unit-sales", "inverse-pricerank", "price", "pricerank", "publication_date", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Author", "Availability", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Power", "Publisher", "Sort", "Title"},
	},
	Electronics: &LocaleSearchIndex{
		BrowseNode:           14011561,
		SortValues:           []string{"-price", "-titlerank", "price", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Availability", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	GiftCards: &LocaleSearchIndex{
		BrowseNode:           0,
		SortValues:           []string{"-price", "-reviewrank_authority", "date-desc-rank", "price", "relevancerank", "reviewrank", "salesrank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice"},
	},
	HealthPersonalCare: &LocaleSearchIndex{
		BrowseNode:           197862031,
		SortValues:           []string{"-price", "price", "relevancerank", "reviewrank", "salesrank"},
		ItemSearchParameters: []string{"Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Kitchen: &LocaleSearchIndex{
		BrowseNode:           57686031,
		SortValues:           []string{"-price", "price", "relevancerank", "salesrank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	MP3Downloads: &LocaleSearchIndex{
		BrowseNode:           206442031,
		SortValues:           []string{"-albumrank", "-artistalbumrank", "-price", "-releasedate", "-runtime", "-titlerank", "albumrank", "artistalbumrank", "price", "relevancerank", "reviewrank", "runtime", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Availability", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	MusicTracks: &LocaleSearchIndex{
		BrowseNode:           301164,
		SortValues:           []string{"-titlerank", "titlerank"},
		ItemSearchParameters: []string{"Author", "Availability", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort"},
	},
	SoftwareVideoGames: &LocaleSearchIndex{
		BrowseNode:           548014,
		SortValues:           []string{"-date", "-pricerank", "-titlerank", "price", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Author", "Availability", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Automotive: &LocaleSearchIndex{
		BrowseNode:           1571266031,
		SortValues:           []string{"-price", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Beauty: &LocaleSearchIndex{
		BrowseNode:           197859031,
		SortValues:           []string{"-price", "price", "relevancerank", "reviewrank", "salesrank"},
		ItemSearchParameters: []string{"Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	SportingGoods: &LocaleSearchIndex{
		BrowseNode:           325615031,
		SortValues:           []string{"-launch-date", "-price", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Availability", "ItemPage", "Keywords", "MerchantId", "MinPercentageOff", "Sort", "Title"},
	},
	VideoGames: &LocaleSearchIndex{
		BrowseNode:           548014,
		SortValues:           []string{"-date", "-price", "-titlerank", "price", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Software: &LocaleSearchIndex{
		BrowseNode:           548012,
		SortValues:           []string{"-date", "-pricerank", "-titlerank", "price", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	VHS: &LocaleSearchIndex{
		BrowseNode:           578610,
		SortValues:           []string{"-titlerank", "amzrank", "availability", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Actor", "AudienceRating", "Availability", "Director", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Publisher", "Sort", "Title"},
	},
	Appliances: &LocaleSearchIndex{
		BrowseNode:           908827031,
		SortValues:           []string{"-price", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	PetSupplies: &LocaleSearchIndex{
		BrowseNode:           1571269031,
		SortValues:           []string{"-price", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"}}}
