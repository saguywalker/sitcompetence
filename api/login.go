package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/saguywalker/sitcompetence/app"
	"github.com/saguywalker/sitcompetence/model"
)

// Login authenticate user with LDAP
func (a *API) Login(w http.ResponseWriter, r *http.Request) {
	session, err := a.App.UserSession.Get(r, "x-session-token")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	user := model.User{}
	var ok bool
	userInterface := session.Values["user"]
	user, ok = userInterface.(model.User)
	if ok && user.Authenticated {
		if err := session.Save(r, w); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

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

	if len(input.Username) == 0 || len(input.Password) == 0 {
		http.Error(w, "username and password must not be empty", http.StatusForbidden)
		return
	}

	user, err = a.App.CheckPassword(input.Username, input.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	ctx := a.App.NewContext()
	ctx.WithUser(user)

	mapResp := make(map[string]interface{})

	if user.Group == "inst_group" {
		pubkey, err := ctx.Database.GetStaffPublicKey(user.UserID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if len(pubkey) == 0 {
			mapResp["first"] = true
		} else {
			mapResp["first"] = false
		}
	}

	resp, err := json.Marshal(mapResp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session.Values["user"] = user

	mapResp["user"] = user

	ctx.Logger.Infof("%+v\n", session.Values["user"])
	ctx.Logger.Infoln(string(resp))

	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(resp)
}

// Logout remove session toke from particular token
func (a *API) Logout(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	session, err := a.App.UserSession.Get(r, "x-session-token")
	if err != nil {
		return err
	}

	val := session.Values["user"]
	user, ok := val.(model.User)
	if !ok {
		return errors.New("user is not found")
	}

	session.Values["user"] = model.User{}
	session.Options.MaxAge = -1

	w.Write([]byte(fmt.Sprintf("%s has been logged out.", user.UserID)))

	return nil
}

// GetUserDetail return user detail from token
func (a *API) GetUserDetail(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	user := ctx.User

	userBytes, err := json.Marshal(user)
	if err != nil {
		return err
	}

	w.Write(userBytes)

	return nil
}
