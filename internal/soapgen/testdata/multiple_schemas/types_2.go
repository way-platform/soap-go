package multiple_schemas

import (
	"encoding/xml"
)

// Element2 represents the Element2 element
type Element2 struct {
	XMLName xml.Name `xml:"Element2"`
	Value   int32    `xml:",chardata"`
}
