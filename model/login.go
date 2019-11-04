package model

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
