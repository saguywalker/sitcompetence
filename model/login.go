package model

import "time"

// Login contain username and password
type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// NewLogin return Login struct
func NewLogin(username, password string) *Login {
	return &Login{
		Username: username,
		Password: password,
	}
}

// LoginAttempt log failed login attempts
type LoginAttempt struct {
	Counter   uint16 `json:"counter"`
	LoginTime time.Time
}
