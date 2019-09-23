package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/saguywalker/sitcompetence/app"
	"github.com/saguywalker/sitcompetence/model"
)

// GetActivities response with all of activities
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
	return nil
}

func (a *API) SearchActivities(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	params := r.URL.Query()

	var activities []*model.Activity

	if params.Get("activity_id") != "" {
		activityID, err := strconv.ParseUint(params.Get("activity_id"), 10, 32)
		if err != nil {
			return err
		}

		activity, err := ctx.GetActivityByID(uint32(activityID))
		if err != nil {
			return err
		}

		activities = append(activities, activity)
	} else {
		var err error

		if params.Get("staff_id") != "" {
			activities, err = ctx.GetActivitiesByStaff(params.Get("staff_id"))
		} else if params.Get("student_id") != "" {
			activities, err = ctx.GetActivitiesByStudent(params.Get("student_id"))
		} else {
			activities, err = ctx.GetActivities()
		}

		if err != nil {
			return err
		}
	}

	data, err := json.Marshal(activities)
	if err != nil {
		return err
	}

	_, err = w.Write(data)
	return nil
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

	return nil
}

/*
// UpdateActivity update an existing activity
func (a *API) UpdateActivity(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	id := getIdFromRequest("id", r)
	var input model.Activity

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, &input); err != nil {
		return err
	}

	if err := ctx.UpdateActivity(&input); err != nil {
		return err
	}

	return nil
}
*/

// DeleteActivityByID delete an existing activity activity
func (a *API) DeleteActivity(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	id := getIdFromRequest("id", r)

	intID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return err
	}

	if err := ctx.DeleteActivity(uint32(intID)); err != nil {
		return err
	}

	return nil
}
