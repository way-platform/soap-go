package duration_type

import (
	"encoding/xml"
)

// WithDurationWrapper represents the WithDuration element
type WithDurationWrapper struct {
	XMLName  xml.Name `xml:"WithDuration"`
	Required string   `xml:"required"`
	Optional *string  `xml:"optional,omitempty"`
}
