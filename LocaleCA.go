package agozon

// LocalCA Map
var LocaleCAMap = map[string]LocaleSearchIndex{
	"All": LocaleCA.All, "Automotive": LocaleCA.Automotive, "Baby": LocaleCA.Baby, "Beauty": LocaleCA.Beauty, "Blended": LocaleCA.Blended, "Books": LocaleCA.Books,
	"Classical": LocaleCA.Classical, "DVD": LocaleCA.DVD, "Electronics": LocaleCA.Electronics, "ForeignBooks": LocaleCA.ForeignBooks, "Grocery": LocaleCA.Grocery,
	"HealthPersonalCare": LocaleCA.HealthPersonalCare, "Jewelry": LocaleCA.Jewelry, "KindleStore": LocaleCA.KindleStore, "Kitchen": LocaleCA.Kitchen,
	"LawnAndGarden": LocaleCA.LawnAndGarden, "Luggage": LocaleCA.Luggage, "Music": LocaleCA.Music, "MusicalInstruments": LocaleCA.MusicalInstruments,
	"OfficeProducts": LocaleCA.OfficeProducts, "PetSupplies": LocaleCA.PetSupplies, "Software": LocaleCA.Software, "SoftwareVideoGames": LocaleCA.SoftwareVideoGames,
	"SportingGoods": LocaleCA.SportingGoods, "Tools": LocaleCA.Tools, "Toys": LocaleCA.Toys, "VHS": LocaleCA.VHS, "Video": LocaleCA.Video, "VideoGames": LocaleCA.VideoGames,
	"Watches": LocaleCA.Watches,
}

// Local CA
var LocaleCA = struct {
	All, Automotive, Baby, Beauty, Blended, Books,
	Classical, DVD, Electronics, ForeignBooks, Grocery,
	HealthPersonalCare, Jewelry, KindleStore, Kitchen,
	LawnAndGarden, Luggage, Music, MusicalInstruments,
	OfficeProducts, PetSupplies, Software, SoftwareVideoGames,
	SportingGoods, Tools, Toys, VHS, Video, VideoGames, Watches LocaleSearchIndex
}{
	Automotive: &LocaleSearchIndex{
		BrowseNode:           6948389011,
		SortValues:           []string{"-price", "date-desc-rank", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Beauty: &LocaleSearchIndex{
		BrowseNode:           6205125011,
		SortValues:           []string{"-price", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	VideoGames: &LocaleSearchIndex{
		BrowseNode:           110218011,
		SortValues:           []string{"-titlerank", "inverse-pricerank", "pricerank", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Watches: &LocaleSearchIndex{
		BrowseNode:           2235621011,
		SortValues:           []string{"-price", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Availability", "Brand", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Baby: &LocaleSearchIndex{
		BrowseNode:           3561347011,
		SortValues:           []string{"-price", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	ForeignBooks: &LocaleSearchIndex{
		BrowseNode:           927726,
		SortValues:           []string{"daterank", "inverse-pricerank", "pricerank", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Author", "Availability", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Power", "Publisher", "Sort", "Title"},
	},
	HealthPersonalCare: &LocaleSearchIndex{
		BrowseNode:           6205178011,
		SortValues:           []string{"-price", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	LawnAndGarden: &LocaleSearchIndex{
		BrowseNode:           6299024011,
		SortValues:           []string{"-price", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Blended: &LocaleSearchIndex{
		BrowseNode:           0,
		SortValues:           []string{""},
		ItemSearchParameters: []string{"Availability", "ItemPage", "Keywords"},
	},
	Grocery: &LocaleSearchIndex{
		BrowseNode:           6967216011,
		SortValues:           []string{"-price", "date-desc-rank", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	MusicalInstruments: &LocaleSearchIndex{
		BrowseNode:           0,
		SortValues:           []string{"-price", "date-desc-rank", "price", "relevancerank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Music: &LocaleSearchIndex{
		BrowseNode:           962454,
		SortValues:           []string{"-orig-rel-date", "-releasedate", "orig-rel-date", "releasedate", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Artist", "Availability", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Video: &LocaleSearchIndex{
		BrowseNode:           962454,
		SortValues:           []string{"-titlerank", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Actor", "AudienceRating", "Availability", "Director", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Publisher", "Sort", "Title"},
	},
	All: &LocaleSearchIndex{
		BrowseNode:           0,
		SortValues:           []string{""},
		ItemSearchParameters: []string{"Availability", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinimumPrice"},
	},
	Books: &LocaleSearchIndex{
		BrowseNode:           927726,
		SortValues:           []string{"daterank", "inverse-pricerank", "pricerank", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Author", "Availability", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Power", "Publisher", "Sort", "Title"},
	},
	DVD: &LocaleSearchIndex{
		BrowseNode:           14113311,
		SortValues:           []string{"salesrank", "titlerank"},
		ItemSearchParameters: []string{"Actor", "AudienceRating", "Availability", "Director", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Publisher", "Sort", "Title"},
	},
	Electronics: &LocaleSearchIndex{
		BrowseNode:           677211011,
		SortValues:           []string{"-price", "-titlerank", "price", "relevancerank", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Author", "Availability", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	KindleStore: &LocaleSearchIndex{
		BrowseNode:           2972706011,
		SortValues:           []string{"-price", "daterank", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Author", "Availability", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Publisher", "Sort", "Title"},
	},
	Luggage: &LocaleSearchIndex{
		BrowseNode:           6205506011,
		SortValues:           []string{"-price", "date-desc-rank", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	"": &LocaleSearchIndex{
		BrowseNode:           0,
		SortValues:           []string{},
		ItemSearchParameters: []string{},
	},
	Classical: &LocaleSearchIndex{
		BrowseNode:           962454,
		SortValues:           []string{"-orig-rel-date", "-releasedate", "orig-rel-date", "releasedate", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Artist", "Availability", "Composer", "Conductor", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Orchestra", "Sort", "Title"},
	},
	Jewelry: &LocaleSearchIndex{
		BrowseNode:           0,
		SortValues:           []string{"-price", "-release-date", "popularityrank", "price", "relevancerank", "reviewrank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	OfficeProducts: &LocaleSearchIndex{
		BrowseNode:           6205512011,
		SortValues:           []string{"-price", "date-desc-rank", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Software: &LocaleSearchIndex{
		BrowseNode:           3234171,
		SortValues:           []string{"-daterank", "inverse-pricerank", "pricerank", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Kitchen: &LocaleSearchIndex{
		BrowseNode:           2206276011,
		SortValues:           []string{"-price", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	SportingGoods: &LocaleSearchIndex{
		BrowseNode:           2242990011,
		SortValues:           []string{"-price", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	VHS: &LocaleSearchIndex{
		BrowseNode:           962072,
		SortValues:           []string{"-titlerank", "salesrank"},
		ItemSearchParameters: []string{"Actor", "AudienceRating", "Availability", "Director", "ItemPage", "Keywords", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Publisher", "Sort", "Title"},
	},
	PetSupplies: &LocaleSearchIndex{
		BrowseNode:           6291628011,
		SortValues:           []string{"-price", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	SoftwareVideoGames: &LocaleSearchIndex{
		BrowseNode:           3323751,
		SortValues:           []string{"-daterank", "inverse-pricerank", "pricerank", "salesrank", "titlerank"},
		ItemSearchParameters: []string{"Author", "Availability", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Tools: &LocaleSearchIndex{
		BrowseNode:           3006903011,
		SortValues:           []string{"-price", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Author", "Availability", "Brand", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"},
	},
	Toys: &LocaleSearchIndex{
		BrowseNode:           6205517011,
		SortValues:           []string{"-price", "date-desc-rank", "price", "relevancerank", "reviewrank", "reviewrank_authority", "salesrank"},
		ItemSearchParameters: []string{"Availability", "ItemPage", "Keywords", "Manufacturer", "MaximumPrice", "MerchantId", "MinPercentageOff", "MinimumPrice", "Sort", "Title"}}}
