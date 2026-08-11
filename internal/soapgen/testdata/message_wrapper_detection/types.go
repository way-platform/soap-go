package message_wrapper_detection

import (
	"encoding/xml"
)

// LoginRequestWrapper represents the LoginRequest element
type LoginRequestWrapper struct {
	XMLName  xml.Name `xml:"LoginRequest"`
	Username string   `xml:"username"`
	Password string   `xml:"password"`
}

// LoginResponseWrapper represents the LoginResponse element
type LoginResponseWrapper struct {
	XMLName   xml.Name `xml:"LoginResponse"`
	SessionId string   `xml:"sessionId"`
	Success   bool     `xml:"success"`
}

// UserDataWrapper represents the UserData element
type UserDataWrapper struct {
	XMLName  xml.Name `xml:"UserData"`
	UserId   int64    `xml:"userId"`
	UserData string   `xml:"userData"`
}

// LoginWrapper represents the login element
type LoginWrapper struct {
	XMLName      xml.Name            `xml:"login"`
	LoginRequest LoginRequestWrapper `xml:"LoginRequest"`
}

// GetUserDataWrapper represents the getUserData element
type GetUserDataWrapper struct {
	XMLName  xml.Name        `xml:"getUserData"`
	UserData UserDataWrapper `xml:"UserData"`
}
