package trailing_spaces_in_element_names

import (
	"encoding/xml"
)

// Complex types

// ResponseType represents the ResponseType complex type
type ResponseType struct {
	Status      string `xml:"status"`
	Priority    int32  `xml:"priority"`
	Data        string `xml:"data"`
	NormalField string `xml:"normalField"`
}

// TestResponse represents the TestResponse  element
type TestResponse struct {
	XMLName     xml.Name `xml:"TestResponse"`
	Status      string   `xml:"status"`
	Priority    int32    `xml:"priority"`
	Data        string   `xml:"data"`
	NormalField string   `xml:"normalField"`
}
