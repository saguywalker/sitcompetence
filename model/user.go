package model

import "time"

// User contains user information
type User struct {
	UserID    string    `json:"uid"`
	Name      string    `json:"username"`
	Group     string    `json:"group"`
	LoginTime time.Time `json:"login_time"`
}

// NewUser return user struct
func NewUser(user, name, group string) *User {
	return &User{
		UserID:    user,
		Name:      name,
		Group:     group,
		LoginTime: time.Now(),
	}
}
