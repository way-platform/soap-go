package context_aware_naming

import (
	"encoding/xml"
)

// Complex types

// ProcessRequestResponseType represents the ProcessRequestResponseType complex type
type ProcessRequestResponseType struct {
	ProcessResult string             `xml:"processResult"`
	ProcessData   ProcessRequestType `xml:"processData"`
}

// ProcessRequestType represents the ProcessRequestType complex type
type ProcessRequestType struct {
	ProcessID string `xml:"processID"`
	Priority  int32  `xml:"priority"`
	Timeout   int32  `xml:"timeout"`
}

// SystemInfoType represents the SystemInfoType complex type
type SystemInfoType struct {
	SystemID string `xml:"systemID"`
	Version  string `xml:"version"`
	Status   string `xml:"status"`
}

// UserDataResponseType represents the UserDataResponseType complex type
type UserDataResponseType struct {
	Result   string       `xml:"result"`
	UserData UserDataType `xml:"userData"`
}

// UserDataType represents the UserDataType complex type
type UserDataType struct {
	UserID   string `xml:"userID"`
	UserName string `xml:"userName"`
	Email    string `xml:"email"`
}

// UserDataWrapper represents the UserData element
type UserDataWrapper struct {
	XMLName  xml.Name `xml:"http://example.com/context-naming UserData"`
	UserID   string   `xml:"userID"`
	UserName string   `xml:"userName"`
	Email    string   `xml:"email"`
}

// ProcessRequestWrapper represents the ProcessRequest element
type ProcessRequestWrapper struct {
	XMLName   xml.Name `xml:"http://example.com/context-naming ProcessRequest"`
	ProcessID string   `xml:"processID"`
	Priority  int32    `xml:"priority"`
	Timeout   int32    `xml:"timeout"`
}

// SystemInfoWrapper represents the SystemInfo element
type SystemInfoWrapper struct {
	XMLName  xml.Name `xml:"http://example.com/context-naming SystemInfo"`
	SystemID string   `xml:"systemID"`
	Version  string   `xml:"version"`
	Status   string   `xml:"status"`
}
