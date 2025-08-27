package multiple_schemas

import (
	"encoding/xml"
)

// Element1 represents the Element1 element
type Element1 struct {
	XMLName xml.Name `xml:"http://example.com/test1 Element1"`
	Value   string   `xml:",chardata"`
}
