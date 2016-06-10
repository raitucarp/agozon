package audience

import (
	"encoding/xml"
	"strings"
	//"strings"
	"testing"
)

func TestAudienceRating(t *testing.T) {
	// test all test
	valid := []Rating{G, PG, PG13, R, NC17, NR, Unrated, Six, Twelve, Sixteen, Eighteen, FamilyViewing}

	for _, v := range valid {
		rating, err := xml.Marshal(v)
		if err != nil {
			t.Error(err)
		}

		err = xml.Unmarshal(rating, &v)
		if err != nil {
			t.Error(err)
		}

		if v.Valid() == false {
			t.Error("Value is not valid")
		}
	}
}

func TestInvalidAudienceRatingValue(t *testing.T) {
	var r Rating = 13
	rating, err := xml.Marshal(r)
	if err != nil {
		t.Error(err)
	}

	err = xml.Unmarshal(rating, &r)
	if err != nil {
		t.Error(err)
	}

	if r.String() != "" {
		t.Error("Out of range value should be empty string")
	}

	if r.Valid() {
		t.Error("Should not valid")
	}
}

func TestFailedAudienceRatingUnmarshal(t *testing.T) {
	f := FamilyViewing
	rating, err := xml.Marshal(FamilyViewing)
	if err != nil {
		t.Error(err)
	}

	var v Rating
	err = xml.Unmarshal(rating, &v)
	if err != nil {
		t.Error(err)
	}

	//s := strings.Replace(string(rating), "AudienceRating", "Audience", -1)
	s := strings.Replace(string(rating), "</Audi", "</Deodor", -1)

	err = xml.Unmarshal([]byte(s), &f)
	if err == nil {
		t.Error("Should error, not returning nil")
	}

	s = strings.Replace(string(rating), "AudienceRating", "Bababa", -1)

	err = xml.Unmarshal([]byte(s), &f)
	if err == nil {
		t.Error("Should error, not returning nil")
	}
}
