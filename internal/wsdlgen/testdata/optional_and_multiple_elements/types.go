package optional_and_multiple_elements

// Container represents the Container element
type Container struct {
	Required string   `xml:"required"`
	Optional *string  `xml:"optional"`
	Multiple []string `xml:"multiple"`
}
