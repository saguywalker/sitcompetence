package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/saguywalker/sitcompetence/app"
)

/*

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
*/
// GetCollectedByStudentID response a competence from requested competenceID
func (a *API) GetCollectedByStudentID(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	id := getIdFromRequest("id", r)
	intID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	competence, err := ctx.GetCompetenceByID(uint16(intID))
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
