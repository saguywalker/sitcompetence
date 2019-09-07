package app

import (
	"github.com/sirupsen/logrus"

	"github.com/saguywalker/sitcompetence/db"
)

// App contains Config and Database
type App struct {
	Config   *Config
	Database *db.Database
}

// NewContext returns reference to context struct
func (a *App) NewContext() *Context {
	return &Context{
		Logger:   logrus.StandardLogger(),
		Database: a.Database,
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
