package schema

// Version is schema version of AWSECommerce
const Version = "2013-08-01"

// MetaData for common MetaData struct
type Metadata struct {
	Key   string `xml:"Key,omitempty"`
	Value string `xml:"Value,omitempty"`
}
