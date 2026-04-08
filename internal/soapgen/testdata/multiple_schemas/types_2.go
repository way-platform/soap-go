package multiple_schemas

import (
	"encoding/xml"
)

// Element2Wrapper represents the Element2 element
type Element2Wrapper struct {
	XMLName xml.Name `xml:"Element2"`
	Value   int32    `xml:",chardata"`
}
