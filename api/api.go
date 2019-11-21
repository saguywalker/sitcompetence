package api

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/ed25519"

	"github.com/saguywalker/sitcompetence/app"
	"github.com/saguywalker/sitcompetence/model"
	protoTm "github.com/saguywalker/sitcompetence/proto/tendermint"
)

type statusCodeRecorder struct {
	http.ResponseWriter
	http.Hijacker
	StatusCode int
}

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

	if err := api.initKeyPair(); err != nil {
		return nil, err
	}

	return api, nil
}

// Init routes api with handler
func (a *API) Init(r *mux.Router) {
	r.HandleFunc("/", a.Index).Methods("GET")
	r.Handle("/giveBadge", a.handler(a.GiveBadge)).Methods("POST", "OPTIONS")
	r.Handle("/approveActivity", a.handler(a.ApproveActivity)).Methods("POST", "OPTIONS")
	r.HandleFunc("/verify", a.VerifyTX).Methods("POST", "OPTIONS")

	r.Handle("/competence", a.handler(a.GetCompetences)).Methods("GET")
	r.Handle("/competence", a.handler(a.CreateCompetence)).Methods("POST", "OPTIONS")
	r.Handle("/competence", a.handler(a.UpdateCompetence)).Methods("PUT", "OPTIONS")
	r.Handle("/competence/{id:[0-9]+}", a.handler(a.DeleteCompetence)).Methods("DELETE")

	r.Handle("/activity", a.handler(a.GetActivities)).Methods("GET")
	r.Handle("/activity", a.handler(a.CreateActivity)).Methods("POST", "OPTIONS")
	r.Handle("/activity", a.handler(a.UpdateActivity)).Methods("PUT", "OPTIONS")
	r.Handle("/activity/{id:[0-9]+}", a.handler(a.DeleteActivity)).Methods("DELETE")

	r.Handle("/student", a.handler(a.GetStudents)).Methods("GET")
	r.Handle("/student", a.handler(a.CreateStudent)).Methods("POST", "OPTIONS")
	r.Handle("/student", a.handler(a.UpdateStudent)).Methods("PUT", "OPTIONS")
	r.Handle("/student/{id:[0-9]+}", a.handler(a.DeleteStudent)).Methods("DELETE", "OPTIONS")
	r.Handle("/student/shareProfile", a.handler(a.ShareProfile)).Methods("GET")
	r.Handle("/student/edit", a.handler(a.EditProfile)).Methods("PUT", "OPTIONS")

	r.Handle("/staff", a.handler(a.GetStaffs)).Methods("GET")
	r.Handle("/staff", a.handler(a.CreateStaff)).Methods("POST", "OPTIONS")
	r.Handle("/staff", a.handler(a.UpdateStaff)).Methods("PUT", "OPTIONS")
	r.Handle("/staff/{id:[0-9]+}", a.handler(a.DeleteStaff)).Methods("DELETE", "OPTIONS")

	adminRoute := r.PathPrefix("/admin").Subrouter()
	adminRoute.Handle("/setkey", a.handler(a.SetPubkey)).Methods("POST", "OPTIONS")
	adminRoute.Handle("/checkKey", a.handler(a.CheckKey)).Methods("GET")
	adminRoute.Handle("/resetKey", a.handler(a.ResetKey)).Methods("GET")
	adminRoute.Handle("/comfirmResetKey", a.handler(a.ConfirmResetKey)).Methods("GET")

	r.Handle("/joinActivity", a.handler(a.JoinActivity)).Methods("POST", "OPTIONS")
	r.HandleFunc("/login", a.Login).Methods("POST", "OPTIONS")
	r.Handle("/logout", a.handler(a.Logout)).Methods("GET")

	r.Handle("/userDetail", a.handler(a.GetUserDetail)).Methods("GET")
	r.HandleFunc("/profile/{url}", a.ViewProfile).Methods("GET")

	searchRoute := r.PathPrefix("/search").Subrouter()
	searchRoute.Handle("/competence", a.handler(a.SearchCompetences)).Methods("GET")
	searchRoute.Handle("/activity", a.handler(a.SearchActivities)).Methods("GET")
	searchRoute.Handle("/student", a.handler(a.SearchStudents)).Methods("GET")
	searchRoute.Handle("/staff", a.handler(a.SearchStaffs)).Methods("GET")
}

func (a *API) handler(f func(*app.Context, http.ResponseWriter, *http.Request) error) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := a.App.UserSession.Get(r, "x-session-token")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		hijacker, _ := w.(http.Hijacker)
		w = &statusCodeRecorder{
			ResponseWriter: w,
			Hijacker:       hijacker,
		}

		// ctx := a.App.NewContext().WithRemoteAddress(a.IPAddressForRequest(r))
		ctx := a.App.NewContext().WithRemoteAddress(r.RemoteAddr)
		ctx.Logger.Infoln(r.RemoteAddr, r.RequestURI)
		user := model.User{}
		var ok bool
		userInterface := session.Values["user"]
		user, ok = userInterface.(model.User)
		if !ok && !user.Authenticated {
			http.Error(w, "unauthenticated", http.StatusForbidden)
			return
		}

		if err := session.Save(r, w); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if r.RequestURI != "/api/admin/setkey" && user.Group == "inst_group" {
			ok, err := ctx.CheckKey(user.UserID)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			if !ok {
				http.Error(w, "Publickey is unset", http.StatusUnavailableForLegalReasons)
				return
			}
		}

		ctx.Logger.Infof("%+v\n", user)

		ctx = ctx.WithUser(user)

		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate, private")
		w.Header().Set("Pragma", "no-cache")

		if err := f(ctx, w, r); err != nil {
			if err.Error() == "does not exists" {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("false"))
			} else if err.Error() == "unauthorization" {
				http.Error(w, err.Error(), http.StatusUnauthorized)
			} else {
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}
		}

	})
}

/*
IPAddressForRequest returns ip address from request
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
*/

// getIDFromRequest returnn id from request
func getIDFromRequest(param string, r *http.Request) string {
	vars := mux.Vars(r)
	id := vars[param]

	return id
}

func getPageParam(r *http.Request) (uint32, error) {
	params := r.URL.Query()
	pageStr := params.Get("page")
	if pageStr == "" {
		return 0, nil
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return 0, err
	}

	return uint32(page), nil
}

func (a *API) initKeyPair() error {
	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return err
	}
	fmt.Printf("pub: %x\npriv: %x\n", pub, priv)

	pub64 := base64.StdEncoding.EncodeToString([]byte(pub))
	priv64 := base64.StdEncoding.EncodeToString([]byte(priv))

	fmt.Printf("pub: %s\npriv: %s\n", pub64, priv64)
	fPub, err := os.Create("key.pub")
	if err != nil {
		return err
	}

	fPriv, err := os.Create("key")
	if err != nil {
		return err
	}

	if _, err := fPub.Write([]byte(pub64)); err != nil {
		return err
	}

	if _, err := fPriv.Write([]byte(priv64)); err != nil {
		return err
	}

	a.App.SK = priv

	params := []byte(fmt.Sprintf("sitcompetence:%s", pub64))

	payload := protoTm.Payload{
		Method: "AddNewService",
		Params: params,
	}

	tx := protoTm.Tx{
		Payload: &payload,
	}

	txBytes, err := proto.Marshal(&tx)
	if err != nil {
		return err
	}

	httpClient := &http.Client{
		Timeout: 3 * time.Second,
	}

	var index uint64 = 0
	for i := 0; i < len(a.Config.Peers); i++ {
		url := fmt.Sprintf("http://%s/broadcast_tx_commit?tx=0x%x", a.Config.Peers[index], txBytes)
		log.Println(url)
		resp, err := httpClient.Get(url)
		index = uint64((index + 1)) % uint64(len(a.Config.Peers))
		if err != nil {
			continue
		} else {
			_, err := ioutil.ReadAll(resp.Body)
			defer resp.Body.Close()
			if err != nil {
				continue
			}
			return nil
		}
	}

	return nil
}

// Index return status 200
func (a *API) Index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
