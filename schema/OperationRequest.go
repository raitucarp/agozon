package schema

import (
	"encoding/xml"
)

// OperationRequest is response for every operation request.
type OperationRequest struct {
	XMLName               xml.Name     `xml:"OperationRequest"`
	HTTPHeaders           *HTTPHeaders `xml:"HTTPHeaders,omitempty"`
	RequestID             string       `xml:"RequestId,omitempty"`
	Arguments             *Arguments   `xml:"Arguments,omitempty"`
	Errors                *Errors      `xml:"Errors,omitempty"`
	RequestProcessingTime float32      `xml:"RequestProcessingTime,omitempty"`
}
