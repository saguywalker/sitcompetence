package api

import (
	"encoding/json"
	"errors"
	"fmt"
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

		students, err = ctx.GetStudents(page, params.Get("dp"), params.Get("year"))
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

	params := r.URL.Query()
	dp := params.Get("dp")
	year := params.Get("year")

	students, err := ctx.GetStudents(page, dp, year)
	if err != nil {
		return err
	}

	for i := range students {
		collected, index, err := ctx.GetCollectedWithDetail(students[i].StudentID, a.App.CurrentPeerIndex, a.Config.Peers)
		if err.Error() == "not exists" {
			continue
		}
		if err != nil {
			return err
		}
		students[i].Collected = collected
		a.App.CurrentPeerIndex = index
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

	id, err := ctx.CreateStudent(&input)
	if err != nil {
		return err
	}

	w.Write([]byte(id))

	return nil
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

// ShareProfile generate a shareable link for student porfolio
func (a *API) ShareProfile(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	if ctx.User.Group != "std_group" {
		return errors.New("only students are allowed")
	}

	session, err := a.App.ProfileSession.Get(r, "url-session")
	if err != nil {
		return err
	}

	if _, ok := session.Values["studentid"]; !ok {
		session.Values["studentid"] = ctx.User.UserID
	}

	if err := session.Save(r, w); err != nil {
		return err
	}

	fullURL := fmt.Sprintf("%s://%s/%s", r.URL.Scheme, r.URL.Host, []byte(session.ID)[:32])

	w.Write([]byte(fullURL))

	return nil
}

// ViewProfile view a specific student's profile
func (a *API) ViewProfile(w http.ResponseWriter, r *http.Request) {
	ctx := a.App.NewContext().WithRemoteAddress(a.IPAddressForRequest(r))
	ctx.Logger.Infoln(r.RemoteAddr, r.RequestURI)

	session, err := a.App.ProfileSession.Get(r, "url-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	studentID, ok := session.Values["studentid"]
	if !ok {
		http.Error(w, "incorrect url", http.StatusNotFound)
		return
	}

	collected, index, err := ctx.GetCollectedWithDetail(studentID.(string), a.App.CurrentPeerIndex, a.Config.Peers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	a.App.CurrentPeerIndex = index

	collectedBytes, err := json.Marshal(collected)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Write(collectedBytes)
}
