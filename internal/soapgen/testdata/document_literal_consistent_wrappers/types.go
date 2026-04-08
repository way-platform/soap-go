package document_literal_consistent_wrappers

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

// UserInfoWrapper represents the UserInfo element
type UserInfoWrapper struct {
	XMLName xml.Name `xml:"UserInfo"`
	UserId  int64    `xml:"userId"`
	Name    string   `xml:"name"`
	Email   string   `xml:"email"`
}

// GetUserResponseWrapper represents the GetUserResponse element
type GetUserResponseWrapper struct {
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

// ServerConfigWrapper represents the ServerConfig element
type ServerConfigWrapper struct {
	XMLName     xml.Name `xml:"ServerConfig"`
	Version     string   `xml:"version"`
	Environment string   `xml:"environment"`
}

// LoginWrapper represents the login element
type LoginWrapper struct {
	XMLName      xml.Name            `xml:"http://example.com/document-literal-test login"`
	LoginRequest LoginRequestWrapper `xml:"LoginRequest"`
}

// GetUserWrapper represents the getUser element
type GetUserWrapper struct {
	XMLName  xml.Name        `xml:"http://example.com/document-literal-test getUser"`
	UserInfo UserInfoWrapper `xml:"UserInfo"`
}
