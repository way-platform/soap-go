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

// FlexibleDocumentWrapper represents the FlexibleDocument element
type FlexibleDocumentWrapper struct {
	XMLName xml.Name             `xml:"FlexibleDocument"`
	Value   FlexibleDocumentType `xml:",chardata"`
}

// DynamicContentWrapper represents the DynamicContent element
type DynamicContentWrapper struct {
	XMLName xml.Name `xml:"DynamicContent"`
	Header  string   `xml:"header"`
	Content RawXML   `xml:",innerxml"`
}

// MixedDocumentWrapper represents the MixedDocument element
type MixedDocumentWrapper struct {
	XMLName            xml.Name `xml:"MixedDocument"`
	KnownElement       string   `xml:"knownElement"`
	UnknownTypeElement *string  `xml:"unknownTypeElement,omitempty"`
	Content            RawXML   `xml:",innerxml"`
}

// PerformanceReportWrapper represents the PerformanceReport element
type PerformanceReportWrapper struct {
	XMLName xml.Name            `xml:"PerformanceReport"`
	Value   PerformanceDataType `xml:",chardata"`
}

// UntypedElementWrapper represents the UntypedElement element
type UntypedElementWrapper struct {
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
