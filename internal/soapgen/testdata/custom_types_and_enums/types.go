package custom_types_and_enums

import (
	"encoding/xml"
)

// Enumeration constants

// StatusType enumeration values
const (
	StatusTypeActive   = "active"
	StatusTypeInactive = "inactive"
	StatusTypePending  = "pending"
)

// Complex types

// UserInfoType represents the UserInfoType complex type
type UserInfoType struct {
	UserId int64  `xml:"userId"`
	Status string `xml:"status"`
	Email  string `xml:"email"`
}

// User represents the User element
type User struct {
	XMLName xml.Name     `xml:"http://example.com/test User"`
	Value   UserInfoType `xml:",chardata"`
}

// StatusCheck represents the StatusCheck element
type StatusCheck struct {
	XMLName       xml.Name `xml:"http://example.com/test StatusCheck"`
	CurrentStatus string   `xml:"currentStatus"`
	TargetStatus  string   `xml:"targetStatus"`
}
