package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/saguywalker/sitcompetence/app"
	"github.com/saguywalker/sitcompetence/model"
)

// SearchStaffs query
func (a *API) SearchStaffs(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	var staffs []*model.Staff
	params := r.URL.Query()

	if params.Get("staff_id") != "" {
		staff, err := ctx.GetStaffByID(params.Get("staff_id"))
		if err != nil {
			return err
		}

		staffs = append(staffs, staff)
	} else {
		page, err := getPageParam(r)
		if err != nil {
			return err
		}

		staffs, err = ctx.GetStaffs(page)

		if err != nil {
			return err
		}
	}

	data, err := json.Marshal(staffs)
	if err != nil {
		return err
	}

	if _, err := w.Write(data); err != nil {
		return err
	}

	return nil
}

// SetPubkey set pubkey from corresponding staff
func (a *API) SetPubkey(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	staffID := ctx.User.UserID

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	if err := ctx.SetPubkey(staffID, body); err != nil {
		return err
	}

	return nil
}

// GetStaffs responses with all of staffs
func (a *API) GetStaffs(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	page, err := getPageParam(r)
	if err != nil {
		return err
	}

	staffs, err := ctx.GetStaffs(page)
	if err != nil {
		return err
	}

	data, err := json.Marshal(staffs)
	if err != nil {
		return err
	}

	_, err = w.Write(data)
	return nil
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

	id, err := ctx.CreateStaff(&input)
	if err != nil {
		return err
	}

	w.Write([]byte(id))

	return nil
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
	return nil
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

	data, err := json.Marshal(input)
	if err != nil {
		return err
	}

	if _, err := w.Write(data); err != nil {
		return err
	}

	return ctx.UpdateStaff(&input)
}

// DeleteStaff delete staff from table
func (a *API) DeleteStaff(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	id := getIDFromRequest("id", r)
	return ctx.DeleteStaff(id)
}
