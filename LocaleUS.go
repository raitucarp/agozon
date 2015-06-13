package agozon

// Locale Information for the US Marketplace
var LocaleUS = struct {
	All, Apparel, Appliances, ArtsAndCrafts,
	Automotive, Baby, Beauty, Blended,
	Books, Classical, Collectibles, DVD,
	DigitalMusic, Electronics, Fashion, FashionBaby,
	FashionBoys, FashionGirls, FashionMen, FashionWomen,
	GiftCards, GourmetFood, Grocery, HealthPersonalCare,
	HomeGarden, Industrial, Jewelry, KindleStore,
	Kitchen, LawnAndGarden, Luggage, MP3Downloads,
	Magazines, Miscellaneous, MobileApps, Music,
	MusicTracks, MusicalInstruments, OfficeProducts, OutdoorLiving,
	PCHardware, PetSupplies, Photo, Shoes, Software,
	SportingGoods, Tools, Toys, UnboxVideo,
	VHS, Video, VideoGames, Watches,
	Wireless LocaleSearchIndex
}{
	All: &LocaleSearchIndex{
		BrowseNode: 0,
		SortValues: []string{},
		ItemSearchParameters: []string{
			"Availability",
			"ItemPage",
			"Keywords",
			"MaximumPrice",
			"MerchantId",
			"MinimumPrice",
		},
	},

	Apparel: &LocaleSearchIndex{
		BrowseNode: 1036592,
		SortValues: []string{
			SortLaunch_DateDesc,
			SortInversePrice,
			SortPriceRankAsc,
			SortRelevanceRank,
			SortSaleFlag,
			SortSalesRank,
		},
		ItemSearchParameters: []string{
			"Author",
			"Availability",
			"Brand",
			"ItemPage",
			"Keywords",
			"Manufacturer",
			"MaximumPrice",
			"MerchantId",
			"MinPercentageOff",
			"MinimumPrice",
			"Sort",
			"Title",
		},
	},

	Appliances: &LocaleSearchIndex{
		BrowseNode: 2619526011,
		SortValues: []string{
			SortPriceDesc,
			SortFeaturedItems,
			SortPriceAsc,
			SortRelevanceRank,
			SortReviewRankAsc,
			SortReviewRank_AuthorityAsc,
			SortSalesRank,
		},
		ItemSearchParameters: []string{
			"Availability",
			"Brand",
			"ItemPage",
			"Keywords",
			"Manufacturer",
			"MaximumPrice",
			"MerchantId",
			"MinPercentageOff",
			"MinimumPrice",
			"Sort",
			"Title",
		},
	},

	ArtsAndCrafts: &LocaleSearchIndex{
		BrowseNode: 2617942011,
		SortValues: []string{
			SortPriceDesc,
			SortFeaturedItems,
			SortPriceAsc,
			SortRelevanceRank,
			SortReviewRankAsc,
			SortReviewRank_AuthorityAsc,
			SortSalesRank,
		},
		ItemSearchParameters: []string{
			"Availability",
			"Brand",
			"ItemPage",
			"Keywords",
			"Manufacturer",
			"MaximumPrice",
			"MerchantId",
			"MinPercentageOff",
			"MinimumPrice",
			"Sort",
			"Title",
		},
	},

	Automotive: &LocaleSearchIndex{
		BrowseNode: 15690151,
		SortValues: []string{
			SortPriceDesc,
			SortTitleRankDesc,
			SortPriceAsc,
			SortRelevanceRank,
			SortSalesRank,
			SortTitleRankAsc,
		},
		ItemSearchParameters: []string{
			"Availability",
			"Brand",
			"ItemPage",
			"Keywords",
			"Manufacturer",
			"MaximumPrice",
			"MerchantId",
			"MinPercentageOff",
			"MinimumPrice",
			"Sort",
			"Title",
		},
	},

	Baby: &LocaleSearchIndex{
		BrowseNode: 165797011,
		SortValues: []string{
			SortPriceDesc,
			SortPriceAsc,
			SortPSRank,
			SortSalesRank,
			SortTitleRankAsc,
		},
		ItemSearchParameters: []string{
			"Author",
			"Availability",
			"Brand",
			"ItemPage",
			"Keywords",
			"Manufacturer",
			"MaximumPrice",
			"MerchantId",
			"MinPercentageOff",
			"MinimumPrice",
			"Sort",
			"Title",
		},
	},

	Beauty: &LocaleSearchIndex{
		BrowseNode: 11055981,
		SortValues: []string{
			SortLaunch_DateDesc,
			SortPriceDesc,
			SortFeaturedItems,
			SortPriceAsc,
			SortSaleFlag,
			SortSalesRank,
		},
		ItemSearchParameters: []string{
			"Author",
			"Availability",
			"Brand",
			"ItemPage",
			"Keywords",
			"Manufacturer",
			"MaximumPrice",
			"MerchantId",
			"MinPercentageOff",
			"MinimumPrice",
			"Sort",
			"Title",
		},
	},

	Blended: &LocaleSearchIndex{
		BrowseNode: 0,
		SortValues: []string{},
		ItemSearchParameters: []string{
			"Availability",
			"ItemPage",
			"Keywords",
		},
	},

	Books: &LocaleSearchIndex{
		BrowseNode: 1000,
		SortValues: []string{
			SortPriceDesc,
			SortPublication_DateDesc,
			SortTitleRankDesc,
			SortUnitSales,
			SortDateRankDesc,
			SortInverse_PriceRank,
			SortPriceAsc,
			SortPriceRankAsc,
			SortRelevanceRank,
			SortReviewRankAsc,
			SortSalesRank,
			SortTitleRankAsc,
		},
		ItemSearchParameters: []string{
			"Author",
			"Availability",
			"ItemPage",
			"Keywords",
			"MaximumPrice",
			"MerchantId",
			"MinPercentageOff",
			"MinimumPrice",
			"Power",
			"Publisher",
			"Sort",
			"Title",
		},
	},

	Classical: &LocaleSearchIndex{
		BrowseNode: 301668,
		SortValues: []string{
			SortOriginalReleaseDateDesc,
			SortPriceDesc,
			SortReleaseDateDesc,
			SortTitleRankDesc,
			SortOriginalReleaseDateAsc,
			SortPriceAsc,
			SortPSRank,
			SortReleaseDateAsc,
			SortSalesRank,
			SortTitleRankAsc,
		},
		ItemSearchParameters: []string{
			"Artist",
			"Author",
			"Availability",
			"Composer",
			"Conductor",
			"ItemPage",
			"Keywords",
			"MaximumPrice",
			"MerchantId",
			"MinPercentageOff",
			"MinimumPrice",
			"Orchestra",
			"Sort",
			"Title",
		},
	},

	Collectibles: &LocaleSearchIndex{
		BrowseNode: 4991426011,
		SortValues: []string{
			SortPriceDesc,
			SortPriceAsc,
			SortRelevanceRank,
			SortReviewRankAsc,
			SortReviewRank_AuthorityAsc,
			SortSalesRank,
		},
		ItemSearchParameters: []string{
			"Artist",
			"Author",
			"Availability",
			"Composer",
			"Conductor",
			"ItemPage",
			"Keywords",
			"MaximumPrice",
			"MerchantId",
			"MinPercentageOff",
			"MinimumPrice",
			"Orchestra",
			"Sort",
			"Title",
		},
	},

	DVD: &LocaleSearchIndex{
		BrowseNode: 2625374011,
		SortValues: []string{
			SortPriceDesc,
			SortReleaseDateDesc,
			SortVideoReleaseDateDesc,
			SortPriceAsc,
			SortRelevanceRank,
			SortSalesRank,
			SortTitleRankAsc,
		},
		ItemSearchParameters: []string{
			"Actor",
			"AudienceRating",
			"Availability",
			"Director",
			"ItemPage",
			"Keywords",
			"MaximumPrice",
			"MerchantId",
			"MinPercentageOff",
			"MinimumPrice",
			"Publisher",
			"Sort",
			"Title",
		},
	},

	DigitalMusic: &LocaleSearchIndex{
		BrowseNode: 624868011,
		SortValues: []string{
			SortSongTitleRank,
			SortUploadDateRank,
		},
		ItemSearchParameters: []string{
			"Artist",
			"Author",
			"Availability",
			"ItemPage",
			"Keywords",
			"MaximumPrice",
			"MerchantId",
			"MinPercentageOff",
			"MinimumPrice",
			"Sort",
			"Title",
		},
	},

	Electronics: &LocaleSearchIndex{
		BrowseNode: 493964,
		SortValues: []string{
			SortPriceDesc,
			SortFeaturedItems,
			SortPriceAsc,
			SortReviewRankAsc,
			SortSalesRank,
			SortTitleRankAsc,
		},
		ItemSearchParameters: []string{
			"Author",
			"Availability",
			"Brand",
			"ItemPage",
			"Keywords",
			"Manufacturer",
			"MaximumPrice",
			"MerchantId",
			"MinPercentageOff",
			"MinimumPrice",
			"Sort",
			"Title",
		},
	},

	Fashion: &LocaleSearchIndex{
		BrowseNode: 0,
		SortValues: []string{
			SortPriceDesc,
			SortLaunch_Date,
			SortPopularity_Rank,
			SortPriceAsc,
			SortRelevanceRank,
			SortReviewRankAsc,
		},
		ItemSearchParameters: []string{
			"Availability",
			"Brand",
			"ItemPage",
			"Keywords",
			"Manufacturer",
			"MaximumPrice",
			"MerchantId",
			"MinPercentageOff",
			"MinimumPrice",
			"Sort",
			"Title",
		},
	},

	FashionBaby: &LocaleSearchIndex{
		BrowseNode: 0,
		SortValues: []string{
			SortPriceDesc,
			SortLaunch_Date,
			SortPopularity_Rank,
			SortPriceAsc,
			SortRelevanceRank,
			SortReviewRankAsc,
		},
		ItemSearchParameters: []string{
			"Availability",
			"Brand",
			"ItemPage",
			"Keywords",
			"Manufacturer",
			"MaximumPrice",
			"MerchantId",
			"MinPercentageOff",
			"MinimumPrice",
			"Sort",
			"Title",
		},
	},

	FashionBoys: &LocaleSearchIndex{
		BrowseNode: 0,
		SortValues: []string{
			SortPriceDesc,
			SortLaunch_Date,
			SortPopularity_Rank,
			SortPriceAsc,
			SortRelevanceRank,
			SortReviewRankAsc,
		},
		ItemSearchParameters: []string{
			"Availability",
			"Brand",
			"ItemPage",
			"Keywords",
			"Manufacturer",
			"MaximumPrice",
			"MerchantId",
			"MinPercentageOff",
			"MinimumPrice",
			"Sort",
			"Title",
		},
	},

	FashionGirls: &LocaleSearchIndex{
		BrowseNode: 0,
		SortValues: []string{
			SortPriceDesc,
			SortLaunch_Date,
			SortPopularity_Rank,
			SortPriceAsc,
			SortRelevanceRank,
			SortReviewRankAsc,
		},
		ItemSearchParameters: []string{
			"Availability",
			"Brand",
			"ItemPage",
			"Keywords",
			"Manufacturer",
			"MaximumPrice",
			"MerchantId",
			"MinPercentageOff",
			"MinimumPrice",
			"Sort",
			"Title",
		},
	},

	FashionMen: &LocaleSearchIndex{
		BrowseNode: 0,
		SortValues: []string{
			SortPriceDesc,
			SortLaunch_Date,
			SortPopularity_Rank,
			SortPriceAsc,
			SortRelevanceRank,
			SortReviewRankAsc,
		},
		ItemSearchParameters: []string{
			"Availability",
			"Brand",
			"ItemPage",
			"Keywords",
			"Manufacturer",
			"MaximumPrice",
			"MerchantId",
			"MinPercentageOff",
			"MinimumPrice",
			"Sort",
			"Title",
		},
	},

	FashionWomen: &LocaleSearchIndex{
		BrowseNode: 0,
		SortValues: []string{
			SortPriceDesc,
			SortLaunch_Date,
			SortPopularity_Rank,
			SortPriceAsc,
			SortRelevanceRank,
			SortReviewRankAsc,
		},
		ItemSearchParameters: []string{
			"Availability",
			"Brand",
			"ItemPage",
			"Keywords",
			"Manufacturer",
			"MaximumPrice",
			"MerchantId",
			"MinPercentageOff",
			"MinimumPrice",
			"Sort",
			"Title",
		},
	},

	GiftCards: &LocaleSearchIndex{
		BrowseNode: 2864120011,
		SortValues: []string{
			SortPriceDesc,
			SortPriceAsc,
			SortRelevanceRank,
			SortReviewRankAsc,
			SortReviewRank_AuthorityAsc,
			SortSalesRank,
		},
		ItemSearchParameters: []string{
			"Author",
			"Availability",
			"Brand",
			"ItemPage",
			"Keywords",
			"MaximumPrice",
			"MerchantId",
			"MinimumPrice",
			"Neighborhood",
			"Sort",
			"Title",
		},
	},

	GourmetFood: &LocaleSearchIndex{
		BrowseNode: 16310211,
		SortValues: []string{
			SortInversePrice,
			SortLaunch_Date,
			SortPriceRankAsc,
			SortRelevanceRank,
			SortSaleFlag,
			SortSalesRank,
		},
		ItemSearchParameters: []string{
			"Author",
			"Availability",
			"Brand",
			"ItemPage",
			"Keywords",
			"Manufacturer",
			"MaximumPrice",
			"MerchantId",
			"MinPercentageOff",
			"MinimumPrice",
			"Sort",
			"Title",
		},
	},

	Grocery: &LocaleSearchIndex{
		BrowseNode: 16310211,
		SortValues: []string{
			SortInversePrice,
			SortLaunch_Date,
			SortPriceRankAsc,
			SortRelevanceRank,
			SortSaleFlag,
			SortSalesRank,
		},
		ItemSearchParameters: []string{
			"Author",
			"Availability",
			"Brand",
			"ItemPage",
			"Keywords",
			"Manufacturer",
			"MaximumPrice",
			"MerchantId",
			"MinPercentageOff",
			"MinimumPrice",
			"Sort",
			"Title",
		},
	},

	HealthPersonalCare: &LocaleSearchIndex{
		BrowseNode: 3760931,
		SortValues: []string{
			SortInversePrice,
			SortLaunch_Date,
			SortFeaturedItems,
			SortPriceRankAsc,
			SortSaleFlag,
			SortSalesRank,
		},
		ItemSearchParameters: []string{
			"Author",
			"Availability",
			"Brand",
			"ItemPage",
			"Keywords",
			"Manufacturer",
			"MaximumPrice",
			"MerchantId",
			"MinPercentageOff",
			"MinimumPrice",
			"Sort",
			"Title",
		},
	},

	HomeGarden: &LocaleSearchIndex{
		BrowseNode: 1063498,
		SortValues: []string{
			SortPriceDesc,
			SortTitleRankDesc,
			SortPriceAsc,
			SortSalesRank,
			SortTitleRankAsc,
		},
		ItemSearchParameters: []string{
			"Author",
			"Availability",
			"Brand",
			"ItemPage",
			"Keywords",
			"Manufacturer",
			"MaximumPrice",
			"MerchantId",
			"MinPercentageOff",
			"MinimumPrice",
			"Neighborhood",
			"Sort",
			"Title",
		},
	},

	Industrial: &LocaleSearchIndex{
		BrowseNode: 16310161,
		SortValues: []string{
			SortPriceDesc,
			SortTitleRankDesc,
			SortFeaturedItems,
			SortPriceAsc,
			SortSalesRank,
			SortTitleRankAsc,
		},
		ItemSearchParameters: []string{
			"Availability",
			"ItemPage",
			"Keywords",
			"Manufacturer",
			"MaximumPrice",
			"MerchantId",
			"MinPercentageOff",
			"MinimumPrice",
			"Sort",
			"Title",
		},
	},

	Jewelry: &LocaleSearchIndex{
		BrowseNode: 2516784011,
		SortValues: []string{
			SortInversePrice,
			SortLaunch_Date,
			SortFeaturedItems,
			SortPriceRankAsc,
			SortSalesRank,
		},
		ItemSearchParameters: []string{
			"Author",
			"Availability",
			"Brand",
			"ItemPage",
			"Keywords",
			"Manufacturer",
			"MaximumPrice",
			"MerchantId",
			"MinPercentageOff",
			"MinimumPrice",
			"Sort",
			"Title",
		},
	},

	KindleStore: &LocaleSearchIndex{
		BrowseNode: 133141011,
		SortValues: []string{
			SortEditionSalesVelocity,
			SortPriceDesc,
			SortDateRankDesc,
			SortPriceAsc,
			SortRelevanceRank,
			SortReviewRankAsc,
			SortSalesRank,
		},
		ItemSearchParameters: []string{
			"Author",
			"Availability",
			"ItemPage",
			"Keywords",
			"MaximumPrice",
			"MerchantId",
			"MinPercentageOff",
			"MinimumPrice",
			"Publisher",
			"Sort",
			"Title",
		},
	},

	Kitchen: &LocaleSearchIndex{
		BrowseNode: 284507,
		SortValues: []string{
			SortPriceDesc,
			SortTitleRankDesc,
			SortFeaturedItems,
			SortPriceAsc,
			SortSalesRank,
			SortTitleRankAsc,
		},
		ItemSearchParameters: []string{
			"Author",
			"Availability",
			"Brand",
			"ItemPage",
			"Keywords",
			"Manufacturer",
			"MaximumPrice",
			"MerchantId",
			"MinPercentageOff",
			"MinimumPrice",
			"Sort",
			"Title",
		},
	},

	LawnAndGarden: &LocaleSearchIndex{
		BrowseNode: 3238155011,
		SortValues: []string{
			SortPriceDesc,
			SortPriceAsc,
			SortRelevanceRank,
			SortReviewRankAsc,
			SortReviewRank_AuthorityAsc,
			SortSalesRank,
		},
		ItemSearchParameters: []string{
			"Author",
			"Availability",
			"Brand",
			"ItemPage",
			"Keywords",
			"Manufacturer",
			"MaximumPrice",
			"MerchantId",
			"MinPercentageOff",
			"MinimumPrice",
			"Sort",
			"Title",
		},
	},

	Luggage: &LocaleSearchIndex{
		BrowseNode: 0,
		SortValues: []string{
			SortPriceDesc,
			SortLaunch_Date,
			SortPopularity_Rank,
			SortPriceAsc,
			SortRelevanceRank,
			SortReviewRankAsc,
		},
		ItemSearchParameters: []string{
			"Availability",
			"Brand",
			"Condition",
			"ItemPage",
			"Keywords",
			"Manufacturer",
			"MaximumPrice",
			"MerchantId",
			"MinPercentageOff",
			"MinimumPrice",
			"Sort",
			"Title",
		},
	},

	MP3Downloads: &LocaleSearchIndex{
		BrowseNode: 624868011,
		SortValues: []string{
			SortPriceDesc,
			SortReleaseDateDesc,
			SortPriceAsc,
			SortRelevanceRank,
			SortSalesRank,
		},
		ItemSearchParameters: []string{
			"Author",
			"Availability",
			"ItemPage",
			"Keywords",
			"MaximumPrice",
			"MerchantId",
			"MinPercentageOff",
			"MinimumPrice",
			"Sort",
			"Title",
		},
	},

	Magazines: &LocaleSearchIndex{
		BrowseNode: 599872,
		SortValues: []string{
			SortPriceDesc,
			SortPublication_DateDesc,
			SortTitleRankDesc,
			SortUnitSales,
			SortDateRankDesc,
			SortPriceAsc,
			SortReviewRankAsc,
			SortSubslotSalesRank,
			SortTitleRankAsc,
		},
		ItemSearchParameters: []string{
			"Author",
			"Availability",
			"ItemPage",
			"Keywords",
			"MaximumPrice",
			"MerchantId",
			"MinPercentageOff",
			"MinimumPrice",
			"Publisher",
			"Sort",
			"Title",
		},
	},

	Miscellaneous: &LocaleSearchIndex{
		BrowseNode: 10304191,
		SortValues: []string{
			SortPriceDesc,
			SortTitleRankDesc,
			SortFeaturedItems,
			SortPriceAsc,
			SortSalesRank,
			SortTitleRankAsc,
		},
		ItemSearchParameters: []string{
			"Author",
			"Availability",
			"Brand",
			"ItemPage",
			"Keywords",
			"Manufacturer",
			"MaximumPrice",
			"MerchantId",
			"MinPercentageOff",
			"MinimumPrice",
			"Sort",
			"Title",
		},
	},

	MobileApps: &LocaleSearchIndex{
		BrowseNode: 2350150011,
		SortValues: []string{
			SortPriceDesc,
			SortFeaturedItems,
			SortPriceAsc,
			SortRelevanceRank,
			SortReviewRankAsc,
			SortReviewRank_AuthorityAsc,
		},
		ItemSearchParameters: []string{
			"Author",
			"Availability",
			"Brand",
			"ItemPage",
			"Keywords",
			"Manufacturer",
			"MaximumPrice",
			"MerchantId",
			"MinPercentageOff",
			"MinimumPrice",
			"Sort",
			"Title",
		},
	},

	Music: &LocaleSearchIndex{
		BrowseNode: 301668,
		SortValues: []string{
			SortOriginalReleaseDateDesc,
			SortPriceDesc,
			SortReleaseDateDesc,
			SortTitleRankDesc,
			SortArtistRank,
			SortOriginalReleaseDateAsc,
			SortPriceAsc,
			SortPSRank,
			SortRelease_DateAsc,
			SortReleaseDateAsc,
			SortRelevanceRank,
			SortSalesRank,
			SortTitleRankAsc,
		},
		ItemSearchParameters: []string{
			"Artist",
			"Availability",
			"ItemPage",
			"Keywords",
			"MaximumPrice",
			"MerchantId",
			"MinPercentageOff",
			"MinimumPrice",
			"Sort",
			"Title",
		},
	},

	MusicTracks: &LocaleSearchIndex{
		BrowseNode: 301668,
		SortValues: []string{
			SortTitleRankDesc,
			SortTitleRankAsc,
		},
		ItemSearchParameters: []string{
			"Author",
			"Availability",
			"ItemPage",
			"Keywords",
			"MaximumPrice",
			"MerchantId",
			"MinPercentageOff",
			"MinimumPrice",
			"Sort",
		},
	},

	MusicalInstruments: &LocaleSearchIndex{
		BrowseNode: 11965861,
		SortValues: []string{
			SortLaunch_DateDesc,
			SortPriceDesc,
			SortFeaturedItems,
			SortPriceAsc,
			SortSaleFlag,
			SortSalesRank,
		},
		ItemSearchParameters: []string{
			"Author",
			"Availability",
			"Brand",
			"ItemPage",
			"Keywords",
			"Manufacturer",
			"MaximumPrice",
			"MerchantId",
			"MinPercentageOff",
			"MinimumPrice",
			"Sort",
			"Title",
		},
	},

	OfficeProducts: &LocaleSearchIndex{
		BrowseNode: 1084128,
		SortValues: []string{
			SortPriceDesc,
			SortFeaturedItems,
			SortPriceAsc,
			SortReviewRankAsc,
			SortSalesRank,
			SortTitleRankAsc,
		},
		ItemSearchParameters: []string{
			"Author",
			"Availability",
			"Brand",
			"ItemPage",
			"Keywords",
			"Manufacturer",
			"MaximumPrice",
			"MerchantId",
			"MinPercentageOff",
			"MinimumPrice",
			"Sort",
			"Title",
		},
	},

	OutdoorLiving: &LocaleSearchIndex{
		BrowseNode: 2972638011,
		SortValues: []string{
			SortPriceDesc,
			SortTitleRankDesc,
			SortPriceAsc,
			SortPSRank,
			SortSalesRank,
			SortTitleRankAsc,
		},
		ItemSearchParameters: []string{
			"Author",
			"Availability",
			"Brand",
			"ItemPage",
			"Keywords",
			"Manufacturer",
			"MaximumPrice",
			"MerchantId",
			"MinPercentageOff",
			"MinimumPrice",
			"Sort",
			"Title",
		},
	},

	PCHardware: &LocaleSearchIndex{
		BrowseNode: 541966,
		SortValues: []string{
			SortPriceDesc,
			SortPriceAsc,
			SortPSRank,
			SortSalesRank,
			SortTitleRankAsc,
		},
		ItemSearchParameters: []string{
			"Author",
			"Availability",
			"Brand",
			"ItemPage",
			"Keywords",
			"Manufacturer",
			"MaximumPrice",
			"MerchantId",
			"MinPercentageOff",
			"MinimumPrice",
			"Sort",
			"Title",
		},
	},

	PetSupplies: &LocaleSearchIndex{
		BrowseNode: 2619534011,
		SortValues: []string{
			SortPriceDesc,
			SortTitleRankDesc,
			SortPriceAsc,
			SortRelevance,
			SortRelevanceRank,
			SortReviewRankAsc,
			SortReviewRank_AuthorityAsc,
			SortSalesRank,
			SortTitleRankAsc,
		},
		ItemSearchParameters: []string{
			"Author",
			"Availability",
			"Brand",
			"ItemPage",
			"Keywords",
			"Manufacturer",
			"MaximumPrice",
			"MerchantId",
			"MinPercentageOff",
			"MinimumPrice",
			"Sort",
			"Title",
		},
	},

	Photo: &LocaleSearchIndex{
		BrowseNode: 502394,
		SortValues: []string{
			SortPriceDesc,
			SortTitleRankDesc,
			SortFeaturedItems,
			SortPriceAsc,
			SortSalesRank,
			SortTitleRankAsc,
		},
		ItemSearchParameters: []string{
			"Author",
			"Availability",
			"Brand",
			"ItemPage",
			"Keywords",
			"Manufacturer",
			"MaximumPrice",
			"MerchantId",
			"MinPercentageOff",
			"MinimumPrice",
			"Sort",
			"Title",
		},
	},

	Shoes: &LocaleSearchIndex{
		BrowseNode: 672124011,
		SortValues: []string{
			SortLaunch_DateDesc,
			SortPriceDesc,
			SortFeaturedItems,
			SortPriceAsc,
			SortRelevanceRank,
			SortReviewRankAsc,
			SortReviewRank_AuthorityAsc,
		},
		ItemSearchParameters: []string{
			"Author",
			"Availability",
			"Brand",
			"ItemPage",
			"Keywords",
			"Manufacturer",
			"MaximumPrice",
			"MerchantId",
			"MinPercentageOff",
			"MinimumPrice",
			"Sort",
			"Title",
		},
	},

	Software: &LocaleSearchIndex{
		BrowseNode: 409488,
		SortValues: []string{
			SortPriceDesc,
			SortFeaturedItems,
			SortPriceAsc,
			SortSalesRank,
			SortTitleRankAsc,
		},
		ItemSearchParameters: []string{
			"Author",
			"Availability",
			"Brand",
			"ItemPage",
			"Keywords",
			"Manufacturer",
			"MaximumPrice",
			"MerchantId",
			"MinPercentageOff",
			"MinimumPrice",
			"Sort",
			"Title",
		},
	},

	SportingGoods: &LocaleSearchIndex{
		BrowseNode: 3375301,
		SortValues: []string{
			SortPriceDesc,
			SortInversePrice,
			SortLaunch_Date,
			SortPriceAsc,
			SortPriceRankAsc,
			SortRelevanceFSRank,
			SortRelevanceRank,
			SortReviewRank_AuthorityAsc,
			SortSaleFlag,
			SortSalesRank,
		},
		ItemSearchParameters: []string{
			"Author",
			"Availability",
			"Brand",
			"ItemPage",
			"Keywords",
			"Manufacturer",
			"MaximumPrice",
			"MerchantId",
			"MinPercentageOff",
			"MinimumPrice",
			"Sort",
			"Title",
		},
	},

	Tools: &LocaleSearchIndex{
		BrowseNode: 468240,
		SortValues: []string{
			SortPriceDesc,
			SortTitleRankDesc,
			SortFeaturedItems,
			SortPriceAsc,
			SortSalesRank,
			SortTitleRankAsc,
		},
		ItemSearchParameters: []string{
			"Author",
			"Availability",
			"Brand",
			"ItemPage",
			"Keywords",
			"Manufacturer",
			"MaximumPrice",
			"MerchantId",
			"MinPercentageOff",
			"MinimumPrice",
			"Sort",
			"Title",
		},
	},

	Toys: &LocaleSearchIndex{
		BrowseNode: 165795011,
		SortValues: []string{
			SortAgeMin,
			SortPriceDesc,
			SortFeaturedItems,
			SortPriceAsc,
			SortSalesRank,
			SortTitleRankAsc,
		},
		ItemSearchParameters: []string{
			"Availability",
			"ItemPage",
			"Keywords",
			"MaximumPrice",
			"MerchantId",
			"MinPercentageOff",
			"MinimumPrice",
			"Sort",
			"Title",
		},
	},

	UnboxVideo: &LocaleSearchIndex{
		BrowseNode: 2858778011,
		SortValues: []string{
			SortLaunch_DateDesc,
			SortPriceDesc,
			SortVideoReleaseDateDesc,
			SortPriceAsc,
			SortRelevanceRank,
			SortSalesRank,
			SortTitleRankAsc,
		},
		ItemSearchParameters: []string{
			"Actor",
			"AudienceRating",
			"Availability",
			"Director",
			"ItemPage",
			"Keywords",
			"MaximumPrice",
			"MerchantId",
			"MinPercentageOff",
			"MinimumPrice",
			"Publisher",
			"Sort",
			"Title",
		},
	},

	VHS: &LocaleSearchIndex{
		BrowseNode: 2625374011,
		SortValues: []string{
			SortPriceDesc,
			SortReleaseDateDesc,
			SortVideoReleaseDateDesc,
			SortPriceAsc,
			SortRelevanceRank,
			SortSalesRank,
			SortTitleRankAsc,
		},
		ItemSearchParameters: []string{
			"Actor",
			"AudienceRating",
			"Availability",
			"Director",
			"ItemPage",
			"Keywords",
			"MaximumPrice",
			"MerchantId",
			"MinPercentageOff",
			"MinimumPrice",
			"Publisher",
			"Sort",
			"Title",
		},
	},

	Video: &LocaleSearchIndex{
		BrowseNode: 404276,
		SortValues: []string{
			SortPriceDesc,
			SortReleaseDateDesc,
			SortVideoReleaseDateDesc,
			SortPriceAsc,
			SortRelevanceRank,
			SortSalesRank,
			SortTitleRankAsc,
		},
		ItemSearchParameters: []string{
			"Actor",
			"AudienceRating",
			"Availability",
			"Director",
			"ItemPage",
			"Keywords",
			"MaximumPrice",
			"MerchantId",
			"MinPercentageOff",
			"MinimumPrice",
			"Publisher",
			"Sort",
			"Title",
		},
	},

	VideoGames: &LocaleSearchIndex{
		BrowseNode: 11846801,
		SortValues: []string{
			SortPriceDesc,
			SortFeaturedItems,
			SortPriceAsc,
			SortSalesRank,
			SortTitleRankAsc,
		},
		ItemSearchParameters: []string{
			"Author",
			"Availability",
			"Brand",
			"ItemPage",
			"Keywords",
			"Manufacturer",
			"MaximumPrice",
			"MerchantId",
			"MinPercentageOff",
			"MinimumPrice",
			"Sort",
			"Title",
		},
	},

	Watches: &LocaleSearchIndex{
		BrowseNode: 378516011,
		SortValues: []string{
			SortPriceDesc,
			SortPriceAsc,
			SortRelevanceRank,
			SortReviewRankAsc,
			SortSalesRank,
		},
		ItemSearchParameters: []string{
			"Author",
			"Availability",
			"Brand",
			"ItemPage",
			"Keywords",
			"Manufacturer",
			"MaximumPrice",
			"MerchantId",
			"MinPercentageOff",
			"MinimumPrice",
			"Sort",
			"Title",
		},
	},

	Wireless: &LocaleSearchIndex{
		BrowseNode: 2335753011,
		SortValues: []string{
			SortTitleRankDesc,
			SortDateRankDesc,
			SortInverse_PriceRank,
			SortPriceRankAsc,
			SortReviewRankAsc,
			SortSalesRank,
			SortTitleRankAsc,
		},
		ItemSearchParameters: []string{
			"Author",
			"Availability",
			"ItemPage",
			"Keywords",
			"MaximumPrice",
			"MerchantId",
			"MinPercentageOff",
			"MinimumPrice",
			"Sort",
			"Title",
		},
	},
}
