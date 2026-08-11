package multiple_schemas

import (
	"encoding/xml"
)

// Element1Wrapper represents the Element1 element
type Element1Wrapper struct {
	XMLName xml.Name `xml:"Element1"`
	Value   string   `xml:",chardata"`
}
