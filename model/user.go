package model

type User struct {
	User  string
	Name  string
	Group string
}

func NewUser(user, name, group string) *User {
	return &User{
		User:  user,
		Name:  name,
		Group: group,
	}
}
