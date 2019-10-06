package model

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewLogin(username, password string) *Login {
	return &Login{
		Username: username,
		Password: password,
	}
}
