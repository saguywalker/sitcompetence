package app

import (
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/saguywalker/sitcompetence/db"
	"github.com/saguywalker/sitcompetence/model"
)

// Context contains request state
type Context struct {
	Logger           logrus.FieldLogger
	RemoteAddress    string
	Database         *db.Database
	PageLimit        uint32
	CurrentPeerIndex uint64
	User             model.User
}

// WithLogger sets a logger
func (ctx *Context) WithLogger(logger logrus.FieldLogger) *Context {
	ret := *ctx
	ret.Logger = logger
	return &ret
}

// WithRemoteAddress sets a remote address
func (ctx *Context) WithRemoteAddress(address string) *Context {
	ret := *ctx
	ret.RemoteAddress = address
	return &ret
}

// WithUser set user to a context
func (ctx *Context) WithUser(user model.User) *Context {
	ret := *ctx
	ret.User = user
	return &ret
}

// AuthorizationError return a reference to UserError struct
func (ctx *Context) AuthorizationError() *UserError {
	return &UserError{Message: "unauthorized", StatusCode: http.StatusForbidden}
}
