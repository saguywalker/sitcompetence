package model

// User contains user information
type User struct {
	UserID        string `json:"uid"`
	Name          string `json:"username"`
	Group         string `json:"group"`
	Authenticated bool   `json:"authenticated"`
}

// NewUser return user struct
func NewUser(user, name, group string, auth bool) User {
	return User{
		UserID:        user,
		Name:          name,
		Group:         group,
		Authenticated: auth,
	}
}
