package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/saguywalker/sitcompetence/app"
	"github.com/saguywalker/sitcompetence/model"
)

// SearchStudents query
func (a *API) SearchStudents(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	var students []*model.Student
	params := r.URL.Query()

	if params.Get("student_id") != "" {
		student, err := ctx.GetStudentByID(params.Get("student_id"))
		if err != nil {
			return err
		}

		students = append(students, student)
	} else {
		page, err := getPageParam(r)
		if err != nil {
			return err
		}

		students, err = ctx.GetStudents(page, params.Get("dp"), params.Get("semester"), params.Get("year"))
		if err != nil {
			return err
		}
	}

	data, err := json.Marshal(students)
	if err != nil {
		return err
	}

	if _, err := w.Write(data); err != nil {
		return err
	}

	return nil
}

// GetStudents responses with all of students
func (a *API) GetStudents(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	page, err := getPageParam(r)
	if err != nil {
		return err
	}

	students, err := ctx.GetStudents(page)
	if err != nil {
		return err
	}

	data, err := json.Marshal(students)
	if err != nil {
		return err
	}

	_, err = w.Write(data)
	return err
}

// CreateStudent creates a student from a request
func (a *API) CreateStudent(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	var input model.Student

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, &input); err != nil {
		return err
	}

	if err := ctx.CreateStudent(&input); err != nil {
		return err
	}

	return err
}

// GetStudentByID responses a student from a requested studentID
func (a *API) GetStudentByID(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]
	student, err := ctx.GetStudentByID(id)
	if err != nil {
		return err
	}

	data, err := json.Marshal(student)
	if err != nil {
		return err
	}

	_, err = w.Write(data)
	return err
}

// UpdateStudent update student table
func (a *API) UpdateStudent(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	var input model.Student

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

	return ctx.UpdateStudent(&input)
}

// DeleteStudent delete student from table
func (a *API) DeleteStudent(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	id := getIDFromRequest("id", r)
	return ctx.DeleteStudent(id)
}
