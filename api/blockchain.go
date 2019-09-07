package api

import (
	"errors"
	"net/http"

	"github.com/saguywalker/sitcompetence/app"
)

func (a *API) GiveBadge(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	return errors.New("unimplemented")
}

func (a *API) ApproveActivity(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	return errors.New("unimplemented")
}
