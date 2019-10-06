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

func NewLDAPClient(username, password, ou string) *ldap.LDAPClient {
	return &ldap.LDAPClient{
		Base:               "dc=sit,dc=kmutt,dc=ac,dc=th",
		Host:               "ld0620.sit.kmutt.ac.th",
		Port:               636,
		GroupFilter:        "(memberUid=%s)",
		InsecureSkipVerify: true,
		UseSSL:             true,
		BindDN:             fmt.Sprintf("uid=%s,ou=People,ou=%s,dc=sit,dc=kmutt,dc=ac,dc=th", username, ou),
		BindPassword:       password,
		UserFilter:         "(uid=%s)",
		Attributes:         []string{"mail", "uid", "radiusGroupName", "cn"},
	}
}

func (a *App) CheckPassword(username, password string) (*UserResponse, error) {
	staffClient := NewLDAPClient(username, password, "staff")

	var ok bool
	var user map[string]string
	var err error

	ok, user, err = staffClient.Authenticate(username, password)
	if err != nil {
		stdClient := NewLDAPClient(username, password, "st")
		ok, user, err = stdClient.Authenticate(username, password)
	}

	if !ok {
		return nil, fmt.Errorf("Authenticating failed for user %s", username)
	}

	if err != nil {
		return nil, err
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
