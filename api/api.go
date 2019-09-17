package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"github.com/saguywalker/sitcompetence/app"
)

type statusCodeRecorder struct {
	http.ResponseWriter
	http.Hijacker
	StatusCode int
}

func (r *statusCodeRecorder) WriteHeader(statusCode int) {
	r.StatusCode = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
}

// API struct
type API struct {
	App              *app.App
	Config           *Config
	CurrentPeerIndex int
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
	//r.Handle("/users/", a.handler(a.CreateUser)).Methods("POST")
	//r.Handle("/", a.handler(a.HelloHandler)).Methods("GET")
	r.Handle("/home", a.handler(a.Home)).Methods("GET")
	r.Handle("/giveBadge", a.handler(a.GiveBadge)).Methods("POST")
	r.Handle("/approveActivity", a.handler(a.ApproveActivity)).Methods("POST")
	r.Handle("/verify", a.handler(a.VerifyTX)).Methods("POST")
	r.Handle("/competences", a.handler(a.GetCompetences)).Methods("GET")
	r.Handle("/competence/{id:[0-9]+}", a.handler(a.GetCompetenceByID)).Methods("GET")
	r.Handle("/competence", a.handler(a.CreateCompetence)).Methods("POST")
	r.Handle("/activities", a.handler(a.GetActivities)).Methods("GET")
	r.Handle("/activity/{id:[0-9]+}", a.handler(a.GetActivityByID)).Methods("GET")
	r.Handle("/activity", a.handler(a.CreateActivity)).Methods("POST")
	r.Handle("/students", a.handler(a.GetStudents)).Methods("GET")
	r.Handle("/student/{id:[0-9]+}", a.handler(a.GetStudentByID)).Methods("GET")
	r.Handle("/student", a.handler(a.CreateStudent)).Methods("POST")
	r.Handle("/staffs", a.handler(a.GetStaffs)).Methods("GET")
	r.Handle("/staff/{id:[0-9]+}", a.handler(a.GetStaffByID)).Methods("GET")
	r.Handle("/staff", a.handler(a.CreateStaff)).Methods("POST")

}

func (a *API) handler(f func(*app.Context, http.ResponseWriter, *http.Request) error) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Body = http.MaxBytesReader(w, r.Body, 100*1024*1024)

		//beginTime := time.Now()

		hijacker, _ := w.(http.Hijacker)
		w = &statusCodeRecorder{
			ResponseWriter: w,
			Hijacker:       hijacker,
		}

		ctx := a.App.NewContext().WithRemoteAddress(a.IPAddressForRequest(r))
		//ctx = ctx.WithLogger(ctx.Logger.WithField("request_id", base64.RawURLEncoding.EncodeToString(model.NewId())))
		/*
				if username, password, ok := r.BasicAuth(); ok {
					user, err := a.App.GetUserByEmail(username)

					if user == nil || err != nil {
						if err != nil {
							ctx.Logger.WithError(err).Error("unable to get user")
						}
						http.Error(w, "invalid credentials", http.StatusForbidden)
						return
					}

					if ok := user.CheckPassword(password); !ok {
						http.Error(w, "invalid credentials", http.StatusForbidden)
						return
					}

					ctx = ctx.WithUser(user)
				}

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
		*/
		w.Header().Set("Content-Type", "application/json")

		if err := f(ctx, w, r); err != nil {
			if verr, ok := err.(*app.ValidationError); ok {
				data, err := json.Marshal(verr)
				if err == nil {
					w.WriteHeader(http.StatusBadRequest)
					_, err = w.Write(data)
				}

				if err != nil {
					ctx.Logger.Error(err)
					http.Error(w, "internal server error", http.StatusInternalServerError)
				}
			} else if uerr, ok := err.(*app.UserError); ok {
				data, err := json.Marshal(uerr)
				if err == nil {
					w.WriteHeader(uerr.StatusCode)
					_, err = w.Write(data)
				}

				if err != nil {
					ctx.Logger.Error(err)
					http.Error(w, "internal server error", http.StatusInternalServerError)
				}
			} else {
				ctx.Logger.Error(err)
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}
		}

	})
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

// Home is for testing purpose only
func (a *API) Home(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	vals := r.URL.Query()
	params, ok := vals["name"]
	if !ok {
		return fmt.Errorf("missing data parameter")
	}

	_, err := w.Write([]byte(fmt.Sprintf("Hello %s", params[0])))
	if err != nil {
		return err
	}

	return nil
}
