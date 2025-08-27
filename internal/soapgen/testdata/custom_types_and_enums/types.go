package custom_types_and_enums

import (
	"encoding/xml"
)

// Enumeration types

// StatusType represents an enumeration type
type StatusType string

// StatusType enumeration values
const (
	StatusTypeActive   StatusType = "active"
	StatusTypeInactive StatusType = "inactive"
	StatusTypePending  StatusType = "pending"
)

// String returns the string representation of StatusType
func (e StatusType) String() string {
	return string(e)
}

// IsValid returns true if the StatusType value is valid
func (e StatusType) IsValid() bool {
	switch e {
	case StatusTypeActive, StatusTypeInactive, StatusTypePending:
		return true
	default:
		return false
	}
}

// Complex types

// UserInfoType represents the UserInfoType complex type
type UserInfoType struct {
	UserId int64      `xml:"userId"`
	Status StatusType `xml:"status"`
	Email  string     `xml:"email"`
}

// User represents the User element
type User struct {
	XMLName xml.Name     `xml:"http://example.com/test User"`
	Value   UserInfoType `xml:",chardata"`
}

// StatusCheck represents the StatusCheck element
type StatusCheck struct {
	XMLName       xml.Name   `xml:"http://example.com/test StatusCheck"`
	CurrentStatus StatusType `xml:"currentStatus"`
	TargetStatus  StatusType `xml:"targetStatus"`
}
