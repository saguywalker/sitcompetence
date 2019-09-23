package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/saguywalker/sitcompetence/app"
	"github.com/saguywalker/sitcompetence/model"
)

// GetCompetences response with all of competences
func (a *API) GetCompetences(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	competences, err := ctx.GetCompetences()
	if err != nil {
		return err
	}

	data, err := json.Marshal(competences)
	if err != nil {
		return err
	}

	_, err = w.Write(data)
	return err
}

// CreateCompetenceResponse defines competenceID for response back
type CreateCompetenceResponse struct {
	CompetenceID uint16 `json:"competence_id"`
}

// CreateCompetence create new competence from request
func (a *API) CreateCompetence(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	var input model.Competence

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, &input); err != nil {
		return err
	}

	if err := ctx.CreateCompetence(&input); err != nil {
		return err
	}

	data, err := json.Marshal(&CreateCompetenceResponse{CompetenceID: input.CompetenceID})
	if err != nil {
		return err
	}

	_, err = w.Write(data)
	return err
}

// GetCompetenceByID response a competence from requested competenceID
func (a *API) GetCompetenceByID(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
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
