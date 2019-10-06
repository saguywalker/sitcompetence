package main

import (
	"log"

	"github.com/jtblin/go-ldap-client"
)

func main() {
	ExAuthenticate()
}

// ExAuthenticate shows how a typical application can verify a login attempt
func ExAuthenticate() {
	client := &ldap.LDAPClient{
		Base:               "dc=sit,dc=kmutt,dc=ac,dc=th",
		Host:               "ld0620.sit.kmutt.ac.th",
		Port:               636,
		GroupFilter:        "(memberUid=%s)",
		InsecureSkipVerify: true,
		UseSSL:             true,
		BindDN:             "uid=lec01,ou=People,ou=staff,dc=sit,dc=kmutt,dc=ac,dc=th",
		BindPassword:       "readonlypassword",
		UserFilter:         "(uid=%s)",
		Attributes:         []string{"givenName", "sn", "mail", "uid"},
	}
	defer client.Close()

	ok, user, err := client.Authenticate("lec01@sit.kmutt.ac.th", "password")
	if err != nil {
		log.Fatalf("Error authenticating user %s: %+v", "username", err)
	}
	if !ok {
		log.Fatalf("Authenticating failed for user %s", "username")
	}
	log.Printf("User: %+v", user)
}

// ExGetGroupsOfUser shows how to retrieve user groups
func ExGetGroupsOfUser() {
	client := &ldap.LDAPClient{
		Base:               "dc=sit,dc=kmutt,dc=ac,dc=th",
		Host:               "ld0620.sit.kmutt.ac.th",
		Port:               636,
		GroupFilter:        "(memberUid=%s)",
		InsecureSkipVerify: true,
		UseSSL:             true,
	}
	defer client.Close()
	groups, err := client.GetGroupsOfUser("lec01@sit.kmutt.ac.th")
	if err != nil {
		log.Fatalf("Error getting groups for user %s: %+v", "lec01@sit.kmutt.ac.th", err)
	}
	log.Printf("Groups: %+v", groups)
}
