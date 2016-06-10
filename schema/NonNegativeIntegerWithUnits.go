package schema

// NonNegativeInteger is non negative integer
type NonNegativeInteger uint

// NonNegativeIntegerWithUnits is non negative integer schema with units.
type NonNegativeIntegerWithUnits struct {
	Value *NonNegativeInteger `xml:",chardata"`
	Units string              `xml:"Units,attr,omitempty"`
}
