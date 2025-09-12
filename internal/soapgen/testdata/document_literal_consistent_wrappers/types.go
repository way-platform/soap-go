package document_literal_consistent_wrappers

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

// UserInfo represents the UserInfo element
type UserInfo struct {
	XMLName xml.Name `xml:"UserInfo"`
	UserId  int64    `xml:"userId"`
	Name    string   `xml:"name"`
	Email   string   `xml:"email"`
}

// GetUserResponse represents the GetUserResponse element
type GetUserResponse struct {
	XMLName xml.Name `xml:"GetUserResponse"`
	User    string   `xml:"user"`
	Role    string   `xml:"role"`
}

// LogoutWrapper represents the logout element
type LogoutWrapper struct {
	XMLName   xml.Name `xml:"http://example.com/document-literal-test logout"`
	SessionId string   `xml:"sessionId"`
}

// LogoutResponseWrapper represents the logoutResponse element
type LogoutResponseWrapper struct {
	XMLName xml.Name `xml:"http://example.com/document-literal-test logoutResponse"`
	Success bool     `xml:"success"`
}

// ServerConfig represents the ServerConfig element
type ServerConfig struct {
	XMLName     xml.Name `xml:"ServerConfig"`
	Version     string   `xml:"version"`
	Environment string   `xml:"environment"`
}

// LoginWrapper represents the login element
type LoginWrapper struct {
	XMLName      xml.Name     `xml:"http://example.com/document-literal-test login"`
	LoginRequest LoginRequest `xml:"LoginRequest"`
}

// LoginResponseWrapper represents the loginResponse element
type LoginResponseWrapper struct {
	XMLName       xml.Name      `xml:"http://example.com/document-literal-test loginResponse"`
	LoginResponse LoginResponse `xml:"LoginResponse"`
}

// GetUserWrapper represents the getUser element
type GetUserWrapper struct {
	XMLName  xml.Name `xml:"http://example.com/document-literal-test getUser"`
	UserInfo UserInfo `xml:"UserInfo"`
}

// GetUserResponseWrapper represents the getUserResponse element
type GetUserResponseWrapper struct {
	XMLName         xml.Name        `xml:"http://example.com/document-literal-test getUserResponse"`
	GetUserResponse GetUserResponse `xml:"GetUserResponse"`
}
