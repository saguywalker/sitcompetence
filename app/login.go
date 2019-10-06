package app

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"

	"github.com/jtblin/go-ldap-client"
	"github.com/saguywalker/sitcompetence/model"
)

type UserResponse struct {
	Token string     `json:"token"`
	User  model.User `json:"user"`
}

func NewLDAPClient(username, password string) *ldap.LDAPClient {
	return &ldap.LDAPClient{
		Base:               "dc=sit,dc=kmutt,dc=ac,dc=th",
		Host:               "ld0620.sit.kmutt.ac.th",
		Port:               636,
		GroupFilter:        "(memberUid=%s)",
		InsecureSkipVerify: true,
		UseSSL:             true,
		BindDN:             fmt.Sprintf("uid=%s,ou=People,ou=staff,dc=sit,dc=kmutt,dc=ac,dc=th", username),
		BindPassword:       password,
		UserFilter:         "(uid=%s)",
		Attributes:         []string{"mail", "uid", "radiusGroupName", "cn"},
	}
}

func (a *App) CheckPassword(username, password string) (*UserResponse, error) {
	client := NewLDAPClient(username, password)
	ok, user, err := client.Authenticate(username, password)
	if err != nil {
		return nil, err
	}

	if !ok {
		return nil, fmt.Errorf("Authenticating failed for user %s", username)
	}

	userStruct := model.NewUser(username, user["cn"], user["radiusGroupName"])

	userJson, err := json.Marshal(userStruct)
	if err != nil {
		return nil, err
	}

	token := fmt.Sprintf("%x", sha256.Sum256(userJson))

	resp := &UserResponse{
		Token: token,
		User:  *userStruct,
	}

	a.TokenUser[token] = userStruct

	return resp, nil
}
