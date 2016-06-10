package schema

import (
	"encoding/xml"
)

// Error is Product Advertising API errors that provide information about syntactical errors in your requests,
// as well as errors that occur during the execution of your request; for example,
// a search for products returns no results.
//
type Error struct {
	// Errors are composed of two elements:
	//
	// The error code is a unique string that identifies the error.
	Code string `xml:"Code,omitempty"`

	// The error message is a human-readable description of the error to help you debug the issue.
	Message string `xml:"Message,omitempty"`
}

// Errors may appear at different levels in your response.
// Their location reflects at what stage in the execution of the request the error was generated and the type of error.
// Errors in syntax that prevent requests from being executed will appear as children of the response's root element.
// An error associated with a particular item in the response will be a child of the Item element.
// See the sample requests for examples of each of these situations.
//
// http://docs.aws.amazon.com/AWSECommerceService/latest/DG/ErrorCodesAndMessages.html
type Errors struct {
	XMLName xml.Name `xml:"Errors"`
	Error   []*Error `xml:"Error,omitempty"`
}
