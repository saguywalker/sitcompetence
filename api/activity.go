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

// GetActivity response with all of activities
func (a *API) GetActivity(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	params := r.URL.Query()

	var activities []*model.Activity

	if params.Get("activity_id") != "" {
		activity, err := ctx.GetActivityByID(params.Get("activity_id"))
		if err != nil {
			return err
		}
		activities = append(activities, activity)
	} else if params.Get("staff_id") != "" {
		activities, err := ctx.GetActivitiesByStaff(params.Get("staff_id"))
		if err != nil {
			return err
		}
	} else {
		activities, err := ctx.GetActivities()
		if err != nil {
			return err
		}
	}

	data, err := json.Marshal(activities)
	if err != nil {
		return err
	}

	_, err = w.Write(data)
	return err
}

// CreateActivityResponse defines activityID for response back
type CreateActivityResponse struct {
	ActivityID uint32 `json:"activity_id"`
}

// CreateActivity create new activity from request
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

// GetActivityByID response an activity from requested activityID
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
