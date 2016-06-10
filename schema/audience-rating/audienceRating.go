package audience

import (
	"encoding/xml"
	"errors"
	"reflect"
)

type Rating int

// enumeration of audience rating
var validValues = []string{"", "G", "PG", "PG-13", "R", "NC-17", "NR", "Unrated", "6", "12", "16", "18", "FamilyViewing"}

const (
	G Rating = iota + 1
	PG
	PG13
	R
	NC17
	NR
	Unrated
	Six
	Twelve
	Sixteen
	Eighteen
	FamilyViewing
)

// MarshalXML marshal audience rating into xml with custom xml structure
func (r Rating) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "AudienceRating"
	e.EncodeToken(start)
	e.EncodeToken(xml.CharData(r.String()))
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

// UnmarshalXML unmarshal audience rating type to be exactly good structure.
func (r *Rating) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	if start.Name.Local != "AudienceRating" {
		return errors.New("The XML tag name should AudienceRating!")
	}

	// create audience rating
	var rating string

	// decode to audience rating
	err := d.DecodeElement(&rating, &start)
	if err != nil {
		return err
	}

	// create index
	var index int

	// find index in validValues
	for k, v := range validValues {
		if v == rating {
			index = k
		}
	}

	// set condition
	*r = Rating(index)
	return nil
}

// String is a stringer for Condition type
func (r Rating) String() (s string) {
	index := reflect.ValueOf(r).Int()

	l := len(validValues)
	if index >= int64(l) {
		return
	}
	return validValues[index]
}

// Valid check whether value is valid or not
func (r *Rating) Valid() bool {
	if r.String() == "" {
		return false
	}
	return true
}
