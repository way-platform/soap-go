package attributes

import (
	"encoding/xml"
)

// WithAttributesWrapper represents the WithAttributes element
type WithAttributesWrapper struct {
	XMLName xml.Name `xml:"WithAttributes"`
	Content string   `xml:"content"`
	Id      string   `xml:"id,attr"`
	Version *int32   `xml:"version,attr,omitempty"`
}
