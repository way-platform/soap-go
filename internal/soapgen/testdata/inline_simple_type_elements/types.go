package inline_simple_type_elements

import (
	"encoding/xml"
)

// Inline enumeration types

// EnabledStatus_EnabledStatus represents an inline enumeration type
type EnabledStatus_EnabledStatus string

// EnabledStatus_EnabledStatus enumeration values
const (
	EnabledStatus_EnabledStatusENABLED  EnabledStatus_EnabledStatus = "ENABLED"
	EnabledStatus_EnabledStatusDISABLED EnabledStatus_EnabledStatus = "DISABLED"
)

// String returns the string representation of EnabledStatus_EnabledStatus
func (e EnabledStatus_EnabledStatus) String() string {
	return string(e)
}

// IsValid returns true if the EnabledStatus_EnabledStatus value is valid
func (e EnabledStatus_EnabledStatus) IsValid() bool {
	switch e {
	case EnabledStatus_EnabledStatusENABLED, EnabledStatus_EnabledStatusDISABLED:
		return true
	default:
		return false
	}
}

// Complex types

// ConfigurationType represents the ConfigurationType complex type
type ConfigurationType struct {
	Name          string        `xml:"name"`
	EnabledStatus EnabledStatus `xml:"EnabledStatus"`
	Priority      Priority      `xml:"Priority"`
	ProductCode   *ProductCode  `xml:"ProductCode,omitempty"`
}

// EnabledStatus represents the EnabledStatus element
type EnabledStatus struct {
	XMLName xml.Name                    `xml:"EnabledStatus"`
	Value   EnabledStatus_EnabledStatus `xml:",chardata"`
}

// Priority represents the Priority element
type Priority struct {
	XMLName xml.Name `xml:"Priority"`
	Value   int32    `xml:",chardata"`
}

// ProductCode represents the ProductCode element
type ProductCode struct {
	XMLName xml.Name `xml:"ProductCode"`
	Value   string   `xml:",chardata"`
}

// Configuration represents the Configuration element
type Configuration struct {
	XMLName xml.Name          `xml:"Configuration"`
	Value   ConfigurationType `xml:",chardata"`
}
