package agozon

type LocaleSearchIndex struct {
	BrowseNode           int64    `json:"BrowseNode"`
	SortValues           []string `json:"SortValues"`
	ItemSearchParameters []string `json:"ItemSearchParameters"`
}

var y = LocaleUS

/*

var LocaleBR = map[string]LocaleSearchIndex{}
var LocaleCA = map[string]LocaleSearchIndex{}
var LocaleCN = map[string]LocaleSearchIndex{}
var LocaleDE = map[string]LocaleSearchIndex{}
var LocaleES = map[string]LocaleSearchIndex{}
var LocaleFR = map[string]LocaleSearchIndex{}
var LocaleIN = map[string]LocaleSearchIndex{}
var LocaleIT = map[string]LocaleSearchIndex{}
var LocaleJP = map[string]LocaleSearchIndex{}
var LocaleUK = map[string]LocaleSearchIndex{}

var LocaleReference = map[string]map[string]LocaleSearchIndex{
	"BR": LocaleBR,
	"CA": LocaleCA,
	"CN": LocaleCN,
	"DE": LocaleDE,
	"ES": LocaleES,
	"FR": LocaleFR,
	"IN": LocaleIN,
	"IT": LocaleIT,
	"JP": LocaleJP,
	"UK": LocaleUK,
	"US": LocaleUS,
}
*/
//var Marketplaces = []Marketplace{}
