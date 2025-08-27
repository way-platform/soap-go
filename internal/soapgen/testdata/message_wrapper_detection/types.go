package message_wrapper_detection

import (
	"encoding/xml"
)

// LoginRequest represents the LoginRequest element
type LoginRequest struct {
	XMLName  xml.Name `xml:"http://example.com/test LoginRequest"`
	Username string   `xml:"username"`
	Password string   `xml:"password"`
}

// LoginResponse represents the LoginResponse element
type LoginResponse struct {
	XMLName   xml.Name `xml:"http://example.com/test LoginResponse"`
	SessionId string   `xml:"sessionId"`
	Success   bool     `xml:"success"`
}

// UserData represents the UserData element
type UserData struct {
	XMLName  xml.Name `xml:"http://example.com/test UserData"`
	UserId   int64    `xml:"userId"`
	UserData string   `xml:"userData"`
}
