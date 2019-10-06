package model

type User struct {
	User  string `json:"uid"`
	Name  string `json:"username"`
	Group string `json:"group"`
}

func NewUser(user, name, group string) *User {
	return &User{
		User:  user,
		Name:  name,
		Group: group,
	}
}
