package simple_element_with_string_type

import (
	"encoding/xml"
)

// TestElement represents the TestElement element
type TestElement struct {
	XMLName xml.Name `xml:"http://example.com/test TestElement"`
	Value   string   `xml:",chardata"`
}
