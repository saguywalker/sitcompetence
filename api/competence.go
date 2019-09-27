package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/saguywalker/sitcompetence/app"
	"github.com/saguywalker/sitcompetence/model"
)

// GetCompetences response with all of competences
func (a *API) GetCompetences(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	page, err := getPageParam(r)
	if err != nil {
		return err
	}

	competences, err := ctx.GetCompetences(page)
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

	if _, err := ctx.CreateCompetence(&input); err != nil {
		return err
	}
	/*
		id, err := ctx.CreateCompetence(&input)
		if err != nil {
			return err
		}

		resp := make(map[string]int64)
		resp["competence_id"] = id

		data, err := json.Marshal(resp)
		if err != nil {
			return err
		}

		if _, err := w.Write(data); err != nil {
			return err
		}
	*/
	return nil
}

// SearchCompetences search competence from key
func (a *API) SearchCompetences(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	params := r.URL.Query()

	var competences []model.Competence

	if params.Get("competence_id") != "" {
		competenceID, err := strconv.ParseUint(params.Get("competence_id"), 10, 16)
		if err != nil {
			return err
		}
		competence, err := ctx.GetCompetenceByID(uint16(competenceID))
		if err != nil {
			return err
		}

		competences = append(competences, *competence)
	} else {
		page, err := getPageParam(r)
		if err != nil {
			return err
		}

		if params.Get("activity_id") != "" {
			activityID, err := strconv.ParseUint(params.Get("activity_id"), 10, 32)
			if err != nil {
				return err
			}
			competences, err = ctx.GetCompetencesByActivityID(uint32(activityID), page)
		} else if params.Get("student_id") != "" {
			competences, err = ctx.GetCollectedByStudentID(params.Get("student_id"), page)
		} else {
			competences, err = ctx.GetCompetences(page)
		}

		if err != nil {
			return err
		}
	}

	data, err := json.Marshal(competences)
	if err != nil {
		return err
	}

	_, err = w.Write(data)

	return err
}

// GetCompetenceByID response a competence from requested competenceID
func (a *API) GetCompetenceByID(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	id := getIDFromRequest("id", r)
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

// UpdateCompetence update competence table
func (a *API) UpdateCompetence(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	var input model.Competence

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, &input); err != nil {
		return err
	}

	return ctx.UpdateCompetence(&input)
}

// DeleteCompetence delete competence from table
func (a *API) DeleteCompetence(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	id := getIDFromRequest("id", r)
	intID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	return ctx.DeleteCompetence(uint16(intID))
}
