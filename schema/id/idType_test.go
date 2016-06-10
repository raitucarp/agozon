package id

import (
	"encoding/xml"
	"strings"
	"testing"
)

func TestIdType(t *testing.T) {
	validVal := []Type{ASIN, UPC, SKU, EAN, ISBN}

	for _, v := range validVal {
		idtype, err := xml.Marshal(v)
		if err != nil {
			t.Error(err)
		}

		err = xml.Unmarshal(idtype, &v)
		if err != nil {
			t.Error(err)
		}

		if v.Valid() == false {
			t.Error("Value is not valid")
		}
	}
}

func TestInvalidIdTypeMarshal(t *testing.T) {
	var typ Type = 13
	idType, err := xml.Marshal(typ)
	if err != nil {
		t.Error(err)
	}

	err = xml.Unmarshal(idType, &typ)
	if err != nil {
		t.Error(err)
	}

	if typ.String() != "" {
		t.Error("Out of range value should be empty string")
	}

	if typ.Valid() {
		t.Error("Should not valid")
	}
}

func TestFailedIdTypeUnmarshal(t *testing.T) {
	asin := ASIN
	idType, err := xml.Marshal(ASIN)
	if err != nil {
		t.Error(err)
	}

	var typ Type
	err = xml.Unmarshal(idType, &typ)
	if err != nil {
		t.Error(err)
	}

	//s := strings.Replace(string(rating), "AudienceRating", "Audience", -1)
	s := strings.Replace(string(idType), "</IdType", "</Galangal", -1)

	err = xml.Unmarshal([]byte(s), &asin)
	if err == nil {
		t.Error("Should error, not returning nil")
	}

	s = strings.Replace(string(idType), "IdType", "Bababa", -1)

	t.Log(s)
	err = xml.Unmarshal([]byte(s), &asin)
	if err == nil {
		t.Error("Should error, not returning nil")
	}
}
