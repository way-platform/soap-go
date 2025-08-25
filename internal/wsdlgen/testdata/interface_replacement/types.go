package interface_replacement

// MixedContent represents the MixedContent element
type MixedContent struct {
	KnownField      string   `xml:"knownField"`
	UnknownField    string   `xml:"unknownField"`
	UnknownArray    []string `xml:"unknownArray"`
	OptionalUnknown *string  `xml:"optionalUnknown"`
}
