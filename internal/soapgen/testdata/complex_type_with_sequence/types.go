package complex_type_with_sequence

import (
	"encoding/xml"
)

// Person represents the Person element
type Person struct {
	XMLName xml.Name `xml:"http://example.com/test Person"`
	Name    string   `xml:"name"`
	Age     int32    `xml:"age"`
}
