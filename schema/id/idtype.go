package id

import (
	"encoding/xml"
	"errors"
	"reflect"
)

type Type int

// valid values of types
var idTypes = []string{"", "ASIN", "UPC", "SKU", "EAN", "ISBN"}

const (
	ASIN Type = 1 + iota
	UPC
	SKU
	EAN
	ISBN
)

// MarshalXML marshal IdTypes into xml with custom xml structure
func (t Type) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "IdType"
	e.EncodeToken(start)
	e.EncodeToken(xml.CharData(t.String()))
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

// UnmarshalXML unmarshal IdTypes to be exactly good structure.
func (t *Type) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if start.Name.Local != "IdType" {
		return errors.New("XML Tag should IdType")
	}

	// create IdType
	var IdType string

	// decode to IdType
	if err := d.DecodeElement(&IdType, &start); err != nil {
		return err
	}

	// create index
	var index int

	// find index in IdTypes
	for k, v := range idTypes {
		if v == IdType {
			index = k
		}
	}

	// set idType
	*t = Type(index)
	return nil
}

// String is a stringer for IdTypes
func (t Type) String() (s string) {
	index := reflect.ValueOf(t).Int()
	l := len(idTypes)
	if index > int64(l) {
		return
	}
	return idTypes[index]
}

// Valid is validate Type
func (t *Type) Valid() bool {
	if t.String() == "" {
		return false
	}
	return true
}
