package complex_type_with_sequence

import (
	"encoding/xml"
)

// PersonWrapper represents the Person element
type PersonWrapper struct {
	XMLName xml.Name `xml:"Person"`
	Name    string   `xml:"name"`
	Age     int32    `xml:"age"`
}
