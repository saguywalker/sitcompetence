package app

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"

	"github.com/jtblin/go-ldap-client"
	"github.com/saguywalker/sitcompetence/model"
)

// UserResponse contain user and his token
type UserResponse struct {
	Token string     `json:"token"`
	User  model.User `json:"user"`
}

// NewLDAPClient return new LDAPClient struct
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

// CheckPassword authenticate user and password
func (a *App) CheckPassword(username, password string) (*UserResponse, error) {
	staffClient := NewLDAPClient(username, password, "staff")

	var user map[string]string
	var err error

	_, user, err = staffClient.Authenticate(username, password)
	if err != nil {
		stdClient := NewLDAPClient(username, password, "st")
		_, user, err = stdClient.Authenticate(username, password)
	}
	/*
		if !ok {
			return nil, fmt.Errorf("Authenticating failed for user %s", username)
		}
	*/
	if err != nil {
		return nil, err
	}

	userStruct := model.NewUser(username, user["cn"], user["radiusGroupName"])

	userJSON, err := json.Marshal(userStruct)
	if err != nil {
		return nil, err
	}

	token := fmt.Sprintf("%x", sha256.Sum256(userJSON))

	resp := &UserResponse{
		Token: token,
		User:  *userStruct,
	}

	a.TokenUser[token] = userStruct
	fmt.Println(a.TokenUser)

	return resp, nil
}
