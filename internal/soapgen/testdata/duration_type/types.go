package duration_type

import (
	"encoding/xml"
)

// WithDuration represents the WithDuration element
type WithDuration struct {
	XMLName  xml.Name `xml:"WithDuration"`
	Required string   `xml:"required"`
	Optional *string  `xml:"optional,omitempty"`
}
