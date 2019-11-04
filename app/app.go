package app

import (
	"crypto/sha256"
	"encoding/gob"

	"github.com/gorilla/sessions"
	"github.com/sirupsen/logrus"

	"github.com/saguywalker/sitcompetence/db"
	"github.com/saguywalker/sitcompetence/model"
)

// App contains Config and Database
type App struct {
	Config   *Config
	Database *db.Database
	// ProfileSession   *sessions.CookieStore
	UserSession      *sessions.CookieStore
	CurrentPeerIndex uint64
}

// NewContext returns reference to context struct
func (a *App) NewContext() *Context {
	return &Context{
		Logger:    logrus.StandardLogger(),
		Database:  a.Database,
		PageLimit: 8,
	}
}

// New return reference to app struct
func New() (app *App, err error) {
	app = &App{}
	app.Config, err = InitConfig()
	if err != nil {
		return nil, err
	}

	dbConfig, err := db.InitConfig()
	if err != nil {
		return nil, err
	}

	app.Database, err = db.New(dbConfig)
	if err != nil {
		return nil, err
	}

	// app.TokenUser = make(map[string]*model.User, 0)
	gob.Register(model.User{})
	keyHash := sha256.Sum256(app.Config.SecretKey)
	app.UserSession = sessions.NewCookieStore(keyHash[:])
	app.UserSession.Options = &sessions.Options{
		MaxAge:   60 * 15,
		HttpOnly: false,
		Path:     "/",
	}

	app.CurrentPeerIndex = 0

	return app, err
}

// Close closes a  database connection
func (a *App) Close() error {
	return a.Database.Close()
}

// ValidationError contains message
type ValidationError struct {
	Message string `json:"message"`
}

func (e *ValidationError) Error() string {
	return e.Message
}

// UserError contains message and statuscode
type UserError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"-"`
}

func (e *UserError) Error() string {
	return e.Message
}
