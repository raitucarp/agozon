package agozon

import "net/url"

var PowerSearchList = []string{
	"after", "ASIN", "author",
	"author-begins", "author-exact",
	"binding", "during", "EISBN",
	"ISBN", "keywords", "keywords-begin",
	"language", "pubdate",
	"publisher", "subject",
	"subject-begins", "subject-words-begin",
	"title", "title-begins",
	"title-words-begin",
}

type PowerSearch url.Values