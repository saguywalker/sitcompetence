package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/saguywalker/sitcompetence/app"
)

// GetCollectedCompetences response with all of competences
func (a *API) GetCollectedCompetences(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	collected, err := ctx.GetCollectedCompetences()
	if err != nil {
		return err
	}

	data, err := json.Marshal(collected)
	if err != nil {
		return err
	}

	_, err = w.Write(data)
	return err
}

// GetCompetenceByID response a competence from requested competenceID
func (a *API) GetCollectedByStudentID(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	id := getCompetenceIDFromRequest("id", r)
	competence, err := ctx.GetCompetenceByID(id)
	if err != nil {
		return err
	}

	data, err := json.Marshal(competence)
	if err != nil {
		return err
	}

	_, err = w.Write(data)
	return err
}

func getCompetenceIDFromRequest(param string, r *http.Request) uint16 {
	vars := mux.Vars(r)
	id := vars[param]

	intID, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		return 0
	}

	return uint16(intID)
}
