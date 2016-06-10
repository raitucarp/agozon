package schema

// DecimalWithUnits measure a unit with a decimal
type DecimalWithUnits struct {
	Value string `xml:",chardata"`
	Units string `xml:"Units,attr,omitempty"`
}
