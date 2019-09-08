package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/saguywalker/sitcompetence/app"
	"github.com/saguywalker/sitcompetence/model"
)

func (a *API) GetStudents(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	students, err := ctx.GetStudents()
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

type CreateStudentResponse struct {
	StudentID string `json:"student_id"`
}

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

	data, err := json.Marshal(&CreateStudentResponse{StudentID: input.StudentID})
	if err != nil {
		return err
	}

	_, err = w.Write(data)
	return err
}

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
