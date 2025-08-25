package custom_types_and_enums

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
	Value UserInfoType `xml:",chardata"`
}

// StatusCheck represents the StatusCheck element
type StatusCheck struct {
	CurrentStatus string `xml:"currentStatus"`
	TargetStatus  string `xml:"targetStatus"`
}
