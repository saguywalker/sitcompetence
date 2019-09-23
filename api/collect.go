package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/saguywalker/sitcompetence/app"
	"github.com/saguywalker/sitcompetence/model"
)

// GetCollectedCompetences response with all of competences
func (a *API) GetCollectedCompetences(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	collected, err := ctx.GetCollectedCompetences
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

// CreateCollectedCompetence create new competence from request
func (a *API) CreateCollectedCompetence(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	var input model.CollectedCompetence

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, &input); err != nil {
		return err
	}

	if err := ctx.CreateCollectedCompetence(&input); err != nil {
		return err
	}

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
