package simple_element_with_string_type

// TestElement represents the TestElement element
type TestElement struct {
	Value string `xml:",chardata"`
}
