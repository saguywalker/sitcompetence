package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime/debug"
	"strconv"
	"strings"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"

	"github.com/saguywalker/sitcompetence/app"
	"github.com/saguywalker/sitcompetence/model"
)

type statusCodeRecorder struct {
	http.ResponseWriter
	http.Hijacker
	StatusCode int
}

/*
func (r *statusCodeRecorder) WriteHeader(statusCode int) {
	r.StatusCode = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
}
*/

// API struct
type API struct {
	App    *app.App
	Config *Config
}

// New returns API struct
func New(a *app.App) (api *API, err error) {
	api = &API{App: a}
	api.Config, err = InitConfig()
	if err != nil {
		return nil, err
	}
	return api, nil
}

// Init routes api with handler
func (a *API) Init(r *mux.Router) {
	// user methods
	r.Handle("/giveBadge", a.handler(a.GiveBadge)).Methods("POST")
	r.Handle("/approveActivity", a.handler(a.ApproveActivity)).Methods("POST")
	r.Handle("/verify", a.handler(a.VerifyTX)).Methods("POST")

	r.Handle("/competence", a.handler(a.GetCompetences)).Methods("GET")
	r.Handle("/competence", a.handler(a.CreateCompetence)).Methods("POST")
	r.Handle("/competence", a.handler(a.UpdateCompetence)).Methods("PUT")
	r.Handle("/competence/{id:[0-9]+}", a.handler(a.DeleteCompetence)).Methods("DELETE")

	r.Handle("/activity", a.handler(a.GetActivities)).Methods("GET")
	r.Handle("/activity", a.handler(a.CreateActivity)).Methods("POST")
	r.Handle("/activity", a.handler(a.UpdateActivity)).Methods("PUT")
	r.Handle("/activity/{id:[0-9]+}", a.handler(a.DeleteActivity)).Methods("DELETE")

	r.Handle("/student", a.handler(a.GetStudents)).Methods("GET")
	r.Handle("/student", a.handler(a.CreateStudent)).Methods("POST")
	r.Handle("/student", a.handler(a.UpdateStudent)).Methods("PUT")
	r.Handle("/student/{id:[0-9]+}", a.handler(a.DeleteStudent)).Methods("DELETE")

	r.Handle("/staff", a.handler(a.GetStaffs)).Methods("GET")
	r.Handle("/staff", a.handler(a.CreateStaff)).Methods("POST")
	r.Handle("/staff", a.handler(a.UpdateStaff)).Methods("PUT")
	r.Handle("/staff/{id:[0-9]+}", a.handler(a.DeleteStaff)).Methods("DELETE")

	r.HandleFunc("/login", a.Login).Methods("POST")
	r.Handle("/logout", a.handler(a.Logout)).Methods("POST")

	r.Handle("/userDetail", a.handler(a.GetUserDetail)).Methods("GET")

	searchRoute := r.PathPrefix("/search").Subrouter()
	searchRoute.Handle("/competence", a.handler(a.SearchCompetences)).Methods("GET")
	searchRoute.Handle("/activity", a.handler(a.SearchActivities)).Methods("GET")
	searchRoute.Handle("/student", a.handler(a.SearchStudents)).Methods("GET")
	searchRoute.Handle("/staff", a.handler(a.SearchStaffs)).Methods("GET")
}

func (a *API) handler(f func(*app.Context, http.ResponseWriter, *http.Request) error) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// r.Body = http.MaxBytesReader(w, r.Body, 100*1024*1024)

		beginTime := time.Now()
		token := r.Header.Get("X-Session-Token")

		token = strings.TrimSpace(token)

		if len(token) == 0 {
			http.Error(w, "Missing auth token", http.StatusForbidden)
			return
		}

		hijacker, _ := w.(http.Hijacker)
		w = &statusCodeRecorder{
			ResponseWriter: w,
			Hijacker:       hijacker,
		}

		ctx := a.App.NewContext().WithRemoteAddress(a.IPAddressForRequest(r))
		ctx.WithUser(a.App.TokenUser[token])
		//ctx = ctx.WithLogger(ctx.Logger.WithField("request_id", base64.RawURLEncoding.EncodeToString(model.NewId())))

		defer func() {
			statusCode := w.(*statusCodeRecorder).StatusCode
			if statusCode == 0 {
				statusCode = 200
			}
			duration := time.Since(beginTime)

			logger := ctx.Logger.WithFields(logrus.Fields{
				"duration":    duration,
				"status_code": statusCode,
				"remote":      ctx.RemoteAddress,
			})
			logger.Info(r.Method + " " + r.URL.RequestURI())
		}()

		defer func() {
			if r := recover(); r != nil {
				ctx.Logger.Error(fmt.Errorf("%v: %s", r, debug.Stack()))
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}
		}()
		// w.Header().Set("Allow", "http://localhost:8082")
		// w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate, private")
		w.Header().Set("Pragma", "no-cache")

		if err := f(ctx, w, r); err != nil {
			if verr, ok := err.(*app.ValidationError); ok {
				data, err := json.Marshal(verr)
				if err == nil {
					// w.WriteHeader(http.StatusBadRequest)
					_, err = w.Write(data)
				}

				if err != nil {
					ctx.Logger.Error(err)
					http.Error(w, "internal server error", http.StatusInternalServerError)
				}
			} else if uerr, ok := err.(*app.UserError); ok {
				data, err := json.Marshal(uerr)
				if err == nil {
					// w.WriteHeader(uerr.StatusCode)
					_, err = w.Write(data)
				}

				if err != nil {
					ctx.Logger.Error(err)
					// w.WriteHeader(http.StatusInternalServerError)
					http.Error(w, "internal server error", http.StatusInternalServerError)
				}
			} else {
				ctx.Logger.Error(err)
				// w.WriteHeader(http.StatusInternalServerError)
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}
		}

	})
}

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

	if len(input.Username) == 0 || len(input.Password) == 0 {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

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

// IPAddressForRequest returns ip address from request
func (a *API) IPAddressForRequest(r *http.Request) string {
	addr := r.RemoteAddr
	if a.Config.ProxyCount > 0 {
		h := r.Header.Get("X-Forwarded-For")
		if h != "" {
			clients := strings.Split(h, ",")
			if a.Config.ProxyCount > len(clients) {
				addr = clients[0]
			} else {
				addr = clients[len(clients)-a.Config.ProxyCount]
			}
		}
	}
	return strings.Split(strings.TrimSpace(addr), ":")[0]
}

// getIDFromRequest returnn id from request
func getIDFromRequest(param string, r *http.Request) string {
	vars := mux.Vars(r)
	id := vars[param]

	return id
}

func getPageParam(r *http.Request) (uint64, error) {
	params := r.URL.Query()
	pageStr := params.Get("page")
	if pageStr == "" {
		return 0, nil
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return 0, err
	}

	return uint64(page), nil
}
