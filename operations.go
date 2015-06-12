package agozon

type Operations interface {
	BrowseNodeLookup()
	CartAdd()
	CartClear()
	CartCreate()
	CartGet()
	CartModify()
	ItemLookup()
	ItemSearch()
	SimilarityLookup()
}
