package message_wrapper_detection

// LoginRequest represents the LoginRequest element
type LoginRequest struct {
	Username string `xml:"username"`
	Password string `xml:"password"`
}

// LoginResponse represents the LoginResponse element
type LoginResponse struct {
	SessionId string `xml:"sessionId"`
	Success   bool   `xml:"success"`
}

// UserData represents the UserData element
type UserData struct {
	UserId   int64  `xml:"userId"`
	UserData string `xml:"userData"`
}
