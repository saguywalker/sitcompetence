package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/saguywalker/sitcompetence/app"
	"github.com/saguywalker/sitcompetence/model"
)

// Login authenticate user with LDAP
func (a *API) Login(w http.ResponseWriter, r *http.Request) {
	var input model.Login

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	if err := json.Unmarshal(body, &input); err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	fmt.Printf("%+v\n", input)

	if len(input.Username) == 0 || len(input.Password) == 0 {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}
	/*
		if utf8.ValidString(input.Username) || utf8.ValidString(input.Password) {
			http.Error(w, "Invalid username or password", http.StatusMethodNotAllowed)
		}
	*/
	respStruct, err := a.App.CheckPassword(input.Username, input.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	ctx := a.App.NewContext()
	ctx.WithUser(respStruct.User)

	resp, err := json.Marshal(respStruct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	w.Write(resp)

}

// Logout remove session toke from particular token
func (a *API) Logout(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	token := r.Header.Get("X-Session-Token")

	a.App.TokenUser[token] = nil

	w.Write([]byte(fmt.Sprintf("%s has been logged out.", token)))

	return nil
}

// GetUserDetail return user detail from token
func (a *API) GetUserDetail(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	token := r.Header.Get("X-Session-Token")

	if val, ok := a.App.TokenUser[token]; ok {
		data, err := json.Marshal(val)
		if err != nil {
			return err
		}

		w.Write(data)
		return nil
	}

	w.Write([]byte("Token not found"))
	return fmt.Errorf("Token not found")
}
