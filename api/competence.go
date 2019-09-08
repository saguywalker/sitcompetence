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

type CreateCompetenceResponse struct {
	CompetenceID uint16 `json:"competence_id"`
}

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

func (a *API) GetCompetenceByID(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	id := getCompetenceIDFromRequest(r)
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

func getCompetenceIDFromRequest(r *http.Request) uint16 {
	vars := mux.Vars(r)
	id := vars["id"]

	intID, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		return 0
	}

	return uint16(intID)
}

/*
type UpdateTodoInput struct {
	Name *string `json:"name"`
	Done *bool   `json:"done"`
}

func (a *API) UpdateTodoById(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	id := getIdFromRequest(r)

	var input UpdateTodoInput

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, &input); err != nil {
		return err
	}

	existingTodo, err := ctx.GetTodoById(id)
	if err != nil || existingTodo == nil {
		return err
	}

	if input.Name != nil {
		existingTodo.Name = *input.Name
	}
	if input.Done != nil {
		existingTodo.Done = *input.Done
	}

	err = ctx.UpdateTodo(existingTodo)
	if err != nil {
		return err
	}

	data, err := json.Marshal(existingTodo)
	if err != nil {
		return err
	}

	_, err = w.Write(data)
	return err
}


func (a *API) DeleteTodoById(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	id := getIdFromRequest(r)
	err := ctx.DeleteTodoById(id)

	if err != nil {
		return err
	}

	return &app.UserError{StatusCode: http.StatusOK, Message: "removed"}
}
*/
