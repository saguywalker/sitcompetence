package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/saguywalker/sitcompetence/app"
	"github.com/saguywalker/sitcompetence/model"
)

// GetStaffs responses with all of staffs
func (a *API) GetStaffs(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	staffs, err := ctx.GetStaffs()
	if err != nil {
		return err
	}

	data, err := json.Marshal(staffs)
	if err != nil {
		return err
	}

	_, err = w.Write(data)
	return err
}

// CreateStaffResponse defines a response with staffID
type CreateStaffResponse struct {
	StaffID string `json:"staff_id"`
}

// CreateStaff creates a staff from a request
func (a *API) CreateStaff(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	var input model.Staff

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, &input); err != nil {
		return err
	}

	if err := ctx.CreateStaff(&input); err != nil {
		return err
	}

	data, err := json.Marshal(&CreateStaffResponse{StaffID: input.StaffID})
	if err != nil {
		return err
	}

	_, err = w.Write(data)
	return err
}

// GetStaffByID responses a staff from a requested staffID
func (a *API) GetStaffByID(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]
	competence, err := ctx.GetStaffByID(id)
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

// UpdateStaff update staff table
func (a *API) UpdateStaff(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	var input model.Staff

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, &input); err != nil {
		return err
	}

	return ctx.UpdateStaff(&input)
}

// DeleteStaff delete staff from table
func (a *API) DeleteStaff(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	id := getIdFromRequest("id", r)
	return ctx.DeleteStaff(id)
}
