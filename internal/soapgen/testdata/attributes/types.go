package attributes

import (
	"encoding/xml"
)

// WithAttributes represents the WithAttributes element
type WithAttributes struct {
	XMLName xml.Name `xml:"http://example.com/test WithAttributes"`
	Content string   `xml:"content"`
	Id      string   `xml:"id,attr"`
	Version *int32   `xml:"version,attr,omitempty"`
}
