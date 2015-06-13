package agozon

const (
	// Age: high to low
	SortAgeMin = "-age-min"
	// Album: A to Z
	SortAlbumRankAsc = "albumrank"
	// Album: Z to A
	SortAlbumRankDesc = "-albumrank"
	// Alphabetical: A to Z
	SortAmazonRank = "amzrank"
	// Artist: A to Z
	SortArtistAlbumRankAsc = "artistalbumrank"
	// Artist: Z to A
	SortArtistAlbumRankDesc = "-artistalbumrank"
	// Artist name: A to Z
	SortArtistRank = "artistrank"
	// Most to least available
	SortAvailability = "availability"
	// Publication date: old to new
	SortDate = "-date"
	// Publication date: new to old
	SortDateRankDesc = "daterank"
	// Publication date: old to new
	SortDateRankAsc = "-daterank"
	// Publication date: new to old
	SortDateDescRank = "date-desc-rank"
	// Quickest to slowest selling products.
	SortEditionSalesVelocity = "-edition-sales-velocity"

	// Price: high to low
	SortInversePrice = "inverseprice"
	//Price: high to low
	SortInverse_Price = "inverse-price"
	// Price: high to low
	SortInverse_PriceRank = "inverse-pricerank"

	// Launch date: newer to older
	SortLaunchDate = "launchdate"
	// Launch date: newer to older
	SortLaunch_Date = "launch-date"
	// Launch date: older to newer
	SortLaunch_DateDesc = "-launch-date"

	// Age: low to high
	SortMFGAgeMinAsc = "mfg-age-min"
	// Age: high to low
	SortMFGAgeMinDesc = "-mfg-age-min"
	// Original release date: earliest to latest
	SortOriginalReleaseDateAsc = "orig-rel-date"
	// Original release date: latest to earliest
	SortOriginalReleaseDateDesc = "-orig-rel-date"

	// Bestseller ranking taking into consideration projected sales.
	// The lower the value, the better the sales.
	SortPaidSalesRank = "paidsalesrank"

	// Discount: high to low
	SortDiscountDesc = "pct-off"
	// Discount: low to high
	SortDiscountAsc = "-pct-off"

	// Featured items
	SortFeaturedItems = "pmrank"

	// Items ranked by popularity
	SortPopularityRank  = "popularityrank"
	SortPopularity_Rank = "popularity-rank"

	// Price: low to high
	SortPriceAsc = "price"
	// Price: high to low
	SortPriceDesc = "-price"

	// Price: low to high
	SortPriceAscRank = "price-asc-rank"
	// Price: high to low
	SortPriceDescRank = "price-desc-rank"
	// Price: low to high
	SortPriceNewBinAsc = "price-new-bin"
	// Price: high to low
	SortPriceNewBinDesc = "-price-new-bin"

	// Price: low to high
	SortPriceRankAsc = "pricerank"
	// Price: high to low
	SortPriceRankDesc = "-pricerank"

	// Bestseller ranking taking into consideration projected sales.
	// The lower the value, the better the sales.
	SortPSRank = "psrank"

	// Publication date: newest to oldest
	SortPubDateAsc = "pubdate"
	// Publication date: oldest to most recent
	SortPubDateDesc = "-pubdate"

	// Publication date: newest to oldest
	SortPublicationDateAsc  = "publicationdate"
	SortPublication_DateAsc = "publication_date"

	// Publication date: oldest to most recent
	SortPublicationDateDesc  = "-publicationdate"
	SortPublication_DateDesc = "-publication_date"

	// Release date: older to newer
	SortReleaseDateAsc  = "releasedate"
	SortRelease_DateAsc = "release-date"

	// Release date: newer to older
	SortReleaseDateDesc  = "-releasedate"
	SortRelease_DateDesc = "-release-date"

	// Items ranked according to the following criteria: how often the keyword appears in the description,
	// where the keyword appears (the ranking is higher when keywords are found in titles and—if there are
	// multiple keywords—how closely they occur in descriptions),
	// and how often customers purchased the products they found using the keyword.
	SortRelevance     = "relevance"
	SortRelevanceRank = "relevancerank"

	//
	SortRelevanceFSRank = "relevance-fs-rank"
	// Highest to lowest ratings in customer reviews.
	SortReviewRankAsc  = "reviewrank"
	SOrtReview_RankAsc = "review-rank"

	// Review rank: high to low
	SortReviewRank_AuthorityAsc = "reviewrank_authority"
	// Review rank: low to high
	SortReviewRank_AuthorityDesc = "-reviewrank_authority"

	// Track length: high to low
	SortRuntimeAsc = "runtime"
	// Track length: low to high
	SortRuntimeDesc = "-runtime"

	// On sale
	SortSaleFlag = "sale-flag"

	// Bestselling
	SortSalesRank = "salesrank"
	// Most popular
	SortSongTitleRank = "songtitlerank"
	// Bestselling
	SortSubslotSalesRank = "subslot-salesrank"

	// Alphabetical: A to Z
	SortTitleRankAsc = "titlerank"
	// Alphabetical: Z to A
	SortTitleRankDesc = "-titlerank"

	//
	SortUnitSales = "-unit-sales"

	// Date added
	SortUploadDateRank = "uploaddaterank"

	// Release date: newer to older
	SortVideoReleaseDateDesc = "-video-release-date"
	//
	SortXSRelevanceRank = "xsrelevancerank"
)
