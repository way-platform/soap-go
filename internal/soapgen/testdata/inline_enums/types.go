package inline_enums

import (
	"encoding/xml"
)

// RawXML captures raw XML content for untyped elements.
type RawXML []byte

// Inline enumeration types

// GetServerPropertiesRequest_Priority represents an inline enumeration type
type GetServerPropertiesRequest_Priority string

// GetServerPropertiesRequest_Priority enumeration values
const (
	GetServerPropertiesRequest_PriorityHIGH   GetServerPropertiesRequest_Priority = "HIGH"
	GetServerPropertiesRequest_PriorityMEDIUM GetServerPropertiesRequest_Priority = "MEDIUM"
	GetServerPropertiesRequest_PriorityLOW    GetServerPropertiesRequest_Priority = "LOW"
)

// String returns the string representation of GetServerPropertiesRequest_Priority
func (e GetServerPropertiesRequest_Priority) String() string {
	return string(e)
}

// IsValid returns true if the GetServerPropertiesRequest_Priority value is valid
func (e GetServerPropertiesRequest_Priority) IsValid() bool {
	switch e {
	case GetServerPropertiesRequest_PriorityHIGH, GetServerPropertiesRequest_PriorityMEDIUM, GetServerPropertiesRequest_PriorityLOW:
		return true
	default:
		return false
	}
}

// GetServerPropertiesRequest_Status represents an inline enumeration type
type GetServerPropertiesRequest_Status string

// GetServerPropertiesRequest_Status enumeration values
const (
	GetServerPropertiesRequest_StatusACTIVE   GetServerPropertiesRequest_Status = "ACTIVE"
	GetServerPropertiesRequest_StatusINACTIVE GetServerPropertiesRequest_Status = "INACTIVE"
)

// String returns the string representation of GetServerPropertiesRequest_Status
func (e GetServerPropertiesRequest_Status) String() string {
	return string(e)
}

// IsValid returns true if the GetServerPropertiesRequest_Status value is valid
func (e GetServerPropertiesRequest_Status) IsValid() bool {
	switch e {
	case GetServerPropertiesRequest_StatusACTIVE, GetServerPropertiesRequest_StatusINACTIVE:
		return true
	default:
		return false
	}
}

// GetServerPropertiesRequest_Version represents an inline enumeration type
type GetServerPropertiesRequest_Version string

// GetServerPropertiesRequest_Version enumeration values
const (
	GetServerPropertiesRequest_Version10 GetServerPropertiesRequest_Version = "1.0"
	GetServerPropertiesRequest_Version20 GetServerPropertiesRequest_Version = "2.0"
)

// String returns the string representation of GetServerPropertiesRequest_Version
func (e GetServerPropertiesRequest_Version) String() string {
	return string(e)
}

// IsValid returns true if the GetServerPropertiesRequest_Version value is valid
func (e GetServerPropertiesRequest_Version) IsValid() bool {
	switch e {
	case GetServerPropertiesRequest_Version10, GetServerPropertiesRequest_Version20:
		return true
	default:
		return false
	}
}

// GetServerPropertiesRequestWrapper represents the GetServerPropertiesRequest element
type GetServerPropertiesRequestWrapper struct {
	XMLName     xml.Name                            `xml:"http://example.com/inlineenums GetServerPropertiesRequest"`
	Priority    GetServerPropertiesRequest_Priority `xml:"Priority"`
	Urgency     GetServerPropertiesRequest_Priority `xml:"Urgency"`
	Status      GetServerPropertiesRequest_Status   `xml:"Status"`
	Description string                              `xml:"Description"`
	Version     *GetServerPropertiesRequest_Version `xml:"Version,attr,omitempty"`
}

// GetServerPropertiesResponseWrapper represents the GetServerPropertiesResponse element
type GetServerPropertiesResponseWrapper struct {
	XMLName xml.Name `xml:"http://example.com/inlineenums GetServerPropertiesResponse"`
	Success bool     `xml:"Success"`
}
