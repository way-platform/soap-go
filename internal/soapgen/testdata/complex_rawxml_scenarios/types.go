package complex_rawxml_scenarios

import (
	"encoding/xml"
	"time"
)

// RawXML captures raw XML content for untyped elements.
type RawXML []byte

// NestedDynamicDocument_NestedDocument represents an inline complex type
type NestedDynamicDocument_NestedDocument struct {
	InnerElement string `xml:"innerElement"`
	Content      RawXML `xml:",innerxml"`
}

// Inline complex types

// Complex types

// FlexibleDocumentType represents the FlexibleDocumentType complex type
type FlexibleDocumentType struct {
	DocumentID   string `xml:"documentID"`
	Version      string `xml:"version"`
	OtherContent RawXML `xml:",innerxml"`
}

// MultiAnyType represents the MultiAnyType complex type
type MultiAnyType struct {
	Section1               string `xml:"section1"`
	Section2               string `xml:"section2"`
	LocalContent           RawXML `xml:",innerxml"`
	TargetNamespaceContent RawXML `xml:",innerxml"`
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

// FlexibleDocumentWrapper represents the FlexibleDocument element
type FlexibleDocumentWrapper struct {
	XMLName xml.Name             `xml:"http://example.com/rawxml-scenarios FlexibleDocument"`
	Value   FlexibleDocumentType `xml:",chardata"`
}

// DynamicContentWrapper represents the DynamicContent element
type DynamicContentWrapper struct {
	XMLName xml.Name `xml:"http://example.com/rawxml-scenarios DynamicContent"`
	Header  string   `xml:"header"`
	Content RawXML   `xml:",innerxml"`
}

// MixedDocumentWrapper represents the MixedDocument element
type MixedDocumentWrapper struct {
	XMLName            xml.Name `xml:"http://example.com/rawxml-scenarios MixedDocument"`
	KnownElement       string   `xml:"knownElement"`
	UnknownTypeElement *string  `xml:"unknownTypeElement,omitempty"`
	Content            RawXML   `xml:",innerxml"`
}

// PerformanceReportWrapper represents the PerformanceReport element
type PerformanceReportWrapper struct {
	XMLName xml.Name            `xml:"http://example.com/rawxml-scenarios PerformanceReport"`
	Value   PerformanceDataType `xml:",chardata"`
}

// UntypedElementWrapper represents the UntypedElement element
type UntypedElementWrapper struct {
	XMLName xml.Name `xml:"http://example.com/rawxml-scenarios UntypedElement"`
	Value   string   `xml:",chardata"`
}

// NestedDynamicDocument represents the NestedDynamicDocument element
type NestedDynamicDocument struct {
	XMLName        xml.Name `xml:"NestedDynamicDocument"`
	OuterElement   string   `xml:"outerElement"`
	NestedDocument RawXML   `xml:",innerxml"`
}

// ValidElement represents the ValidElement element
type ValidElement struct {
	XMLName xml.Name  `xml:"ValidElement"`
	Value   ValidType `xml:",chardata"`
}
