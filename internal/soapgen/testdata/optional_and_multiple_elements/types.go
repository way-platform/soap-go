package optional_and_multiple_elements

import (
	"encoding/xml"
)

// Container represents the Container element
type Container struct {
	XMLName  xml.Name `xml:"http://example.com/test Container"`
	Required string   `xml:"required"`
	Optional *string  `xml:"optional,omitempty"`
	Multiple []string `xml:"multiple"`
}
