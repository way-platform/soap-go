package simple_element_with_string_type

import (
	"encoding/xml"
)

// TestElementWrapper represents the TestElement element
type TestElementWrapper struct {
	XMLName xml.Name `xml:"TestElement"`
	Value   string   `xml:",chardata"`
}
