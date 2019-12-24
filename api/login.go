package api

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

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

	input.Username = strings.TrimSpace(input.Username)
	input.Password = strings.TrimSpace(input.Password)

	if len(input.Username) == 0 || len(input.Password) == 0 {
		http.Error(w, "username and password must not be empty", http.StatusForbidden)
		return
	}

	// Check for blocking login attempts
	if loginLog, exists := a.App.FailedLogin[input.Username]; exists {
		if loginLog.Counter > 5 && loginLog.LoginTime.Add(time.Minute*10).Local().Before(time.Now()) {
			fmt.Println("blocking login attempts")
			http.Error(w, "login attempt is blocked", http.StatusMethodNotAllowed)
			return
		}
	}

	decPassword, err := hex.DecodeString(input.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	iv, err := hex.DecodeString("00112233445566778899aabbccddeeff")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}

	block, err := aes.NewCipher(a.App.Config.SecretKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// mode := cipher.NewCBCDecrypter(block, iv)
	// mode.CryptBlocks(decPassword, decPassword)
	mode := cipher.NewOFB(block, iv)
	mode.XORKeyStream(decPassword, decPassword)

	user, err = a.App.CheckPassword(input.Username, string(decPassword))
	if err != nil {
		fmt.Println("invalid username or password")
		if value, exists := a.App.FailedLogin[input.Username]; exists {
			a.App.FailedLogin[input.Username].Counter = value.Counter + 1
			a.App.FailedLogin[input.Username].LoginTime = time.Now()
		} else {
			a.App.FailedLogin[input.Username] = &model.LoginAttempt{Counter: 1, LoginTime: time.Now()}
		}
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	if _, exists := a.App.FailedLogin[input.Username]; exists {
		a.App.FailedLogin[input.Username] = nil
	}

	ctx := a.App.NewContext()
	ctx.WithUser(user)

	mapResp := make(map[string]interface{})

	if user.Group == "inst_group" {
		pubkey, err := ctx.Database.GetStaffPublicKey(user.UserID)
		if err != nil {
			ctx.Logger.Infof("error pubkey: %x\n", pubkey)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if len(pubkey) == 0 {
			mapResp["first"] = true
		} else {
			mapResp["first"] = false
		}
	} else {
		var err error
		user.Additional, err = ctx.FetchStudentProfile(user.UserID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}
	}

	mapResp["user"] = user

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
