package complex_rawxml_scenarios

import (
	"encoding/xml"
)

// RawXML captures raw XML content for untyped elements.
type RawXML []byte

// Complex types

// FlexibleDocumentType represents the FlexibleDocumentType complex type
type FlexibleDocumentType struct {
	DocumentID string `xml:"documentID"`
	Version    string `xml:"version"`
}

// MultiAnyType represents the MultiAnyType complex type
type MultiAnyType struct {
	Section1 string `xml:"section1"`
	Section2 string `xml:"section2"`
}

// PerformanceDataType represents the PerformanceDataType complex type
type PerformanceDataType struct {
	Timestamp  time.Time `xml:"timestamp"`
	Metrics    *string   `xml:"metrics,omitempty"`
	CustomData *string   `xml:"customData,omitempty"`
}

// ValidType represents the ValidType complex type
type ValidType struct {
	ValidElement string `xml:"validElement"`
}

// NestedDynamicDocument_NestedDocument represents an inline complex type
type NestedDynamicDocument_NestedDocument struct {
	InnerElement string `xml:"innerElement"`
}

// Inline complex types

// FlexibleDocument represents the FlexibleDocument element
type FlexibleDocument struct {
	XMLName xml.Name             `xml:"FlexibleDocument"`
	Value   FlexibleDocumentType `xml:",chardata"`
}

// DynamicContent represents the DynamicContent element
type DynamicContent struct {
	XMLName xml.Name `xml:"DynamicContent"`
	Header  string   `xml:"header"`
	Content RawXML   `xml:",innerxml"`
}

// MixedDocument represents the MixedDocument element
type MixedDocument struct {
	XMLName            xml.Name `xml:"MixedDocument"`
	KnownElement       string   `xml:"knownElement"`
	UnknownTypeElement *string  `xml:"unknownTypeElement,omitempty"`
	Content            RawXML   `xml:",innerxml"`
}

// PerformanceReport represents the PerformanceReport element
type PerformanceReport struct {
	XMLName xml.Name            `xml:"PerformanceReport"`
	Value   PerformanceDataType `xml:",chardata"`
}

// UntypedElement represents the UntypedElement element
type UntypedElement struct {
	XMLName xml.Name `xml:"UntypedElement"`
	Value   string   `xml:",chardata"`
}

// NestedDynamicDocument represents the NestedDynamicDocument element
type NestedDynamicDocument struct {
	XMLName        xml.Name                             `xml:"NestedDynamicDocument"`
	OuterElement   string                               `xml:"outerElement"`
	NestedDocument NestedDynamicDocument_NestedDocument `xml:"nestedDocument"`
}

// ValidElement represents the ValidElement element
type ValidElement struct {
	XMLName xml.Name  `xml:"ValidElement"`
	Value   ValidType `xml:",chardata"`
}
