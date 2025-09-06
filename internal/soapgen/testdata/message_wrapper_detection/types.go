package message_wrapper_detection

import (
	"encoding/xml"
)

// LoginRequest represents the LoginRequest element
type LoginRequest struct {
	XMLName  xml.Name `xml:"LoginRequest"`
	Username string   `xml:"username"`
	Password string   `xml:"password"`
}

// LoginResponse represents the LoginResponse element
type LoginResponse struct {
	XMLName   xml.Name `xml:"LoginResponse"`
	SessionId string   `xml:"sessionId"`
	Success   bool     `xml:"success"`
}

// UserData represents the UserData element
type UserData struct {
	XMLName  xml.Name `xml:"UserData"`
	UserId   int64    `xml:"userId"`
	UserData string   `xml:"userData"`
}

// Login represents the login element
type Login struct {
	XMLName      xml.Name     `xml:"login"`
	LoginRequest LoginRequest `xml:"LoginRequest"`
}

// LoginResponseWrapper represents the loginResponse element
type LoginResponseWrapper struct {
	XMLName       xml.Name      `xml:"loginResponse"`
	LoginResponse LoginResponse `xml:"LoginResponse"`
}

// GetUserData represents the getUserData element
type GetUserData struct {
	XMLName  xml.Name `xml:"getUserData"`
	UserData UserData `xml:"UserData"`
}
