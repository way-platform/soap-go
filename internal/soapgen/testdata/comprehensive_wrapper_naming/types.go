package comprehensive_wrapper_naming

import (
	"encoding/xml"
	"time"
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

// LoginData represents the LoginData element
type LoginData struct {
	XMLName       xml.Name  `xml:"LoginData"`
	LoginAttempts int32     `xml:"loginAttempts"`
	LastLogin     time.Time `xml:"lastLogin"`
}

// LoginStats represents the loginStats element
type LoginStats struct {
	XMLName          xml.Name `xml:"loginStats"`
	TotalAttempts    int32    `xml:"totalAttempts"`
	SuccessfulLogins int32    `xml:"successfulLogins"`
}

// LoginWrapper represents the login element
type LoginWrapper struct {
	XMLName      xml.Name     `xml:"login"`
	LoginRequest LoginRequest `xml:"LoginRequest"`
}

// LoginResponseWrapper represents the loginResponse element
type LoginResponseWrapper struct {
	XMLName       xml.Name      `xml:"loginResponse"`
	LoginResponse LoginResponse `xml:"LoginResponse"`
}

// GetUserInfoWrapper represents the getUserInfo element
type GetUserInfoWrapper struct {
	XMLName  xml.Name `xml:"getUserInfo"`
	UserInfo UserInfo `xml:"UserInfo"`
}
