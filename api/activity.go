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

func (a *API) GetActivities(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	activities, err := ctx.GetActivities()
	if err != nil {
		return err
	}

	data, err := json.Marshal(activities)
	if err != nil {
		return err
	}

	_, err = w.Write(data)
	return err
}

type CreateActivityResponse struct {
	ActivityID uint32 `json:"activity_id"`
}

func (a *API) CreateActivity(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	var input model.Activity

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, &input); err != nil {
		return err
	}

	if err := ctx.CreateActivity(&input); err != nil {
		return err
	}

	data, err := json.Marshal(&CreateActivityResponse{ActivityID: input.ActivityID})
	if err != nil {
		return err
	}

	_, err = w.Write(data)
	return err
}

func (a *API) GetActivityByID(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	id := getActivityIDFromRequest(r)
	activity, err := ctx.GetActivityByID(id)
	if err != nil {
		return err
	}

	data, err := json.Marshal(activity)
	if err != nil {
		return err
	}

	_, err = w.Write(data)
	return err
}

func getActivityIDFromRequest(r *http.Request) uint32 {
	vars := mux.Vars(r)
	id := vars["id"]

	intID, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		return 0
	}

	return uint32(intID)
}
