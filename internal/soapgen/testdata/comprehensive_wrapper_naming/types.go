package comprehensive_wrapper_naming

import (
	"encoding/xml"
	"time"
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

// LoginDataWrapper represents the LoginData element
type LoginDataWrapper struct {
	XMLName       xml.Name  `xml:"LoginData"`
	LoginAttempts int32     `xml:"loginAttempts"`
	LastLogin     time.Time `xml:"lastLogin"`
}

// LoginStatsWrapper represents the loginStats element
type LoginStatsWrapper struct {
	XMLName          xml.Name `xml:"loginStats"`
	TotalAttempts    int32    `xml:"totalAttempts"`
	SuccessfulLogins int32    `xml:"successfulLogins"`
}

// LoginWrapper represents the login element
type LoginWrapper struct {
	XMLName      xml.Name            `xml:"http://example.com/test login"`
	LoginRequest LoginRequestWrapper `xml:"LoginRequest"`
}

// GetUserInfoWrapper represents the getUserInfo element
type GetUserInfoWrapper struct {
	XMLName  xml.Name        `xml:"http://example.com/test getUserInfo"`
	UserInfo UserInfoWrapper `xml:"UserInfo"`
}
