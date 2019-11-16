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

		for i := range students {
			collected, evidence, index, err := ctx.GetCollectedWithDetail(fmt.Sprintf("student_id=%s", students[i].StudentID), a.App.CurrentPeerIndex, a.Config.Peers)
			if err != nil {
				if err.Error() == "does not exists" {
					continue
				}
				return errors.New(string(evidence))
			}

			students[i].Collected = collected
			students[i].Evidence = evidence
			a.App.CurrentPeerIndex = index
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
		collected, evidence, index, err := ctx.GetCollectedWithDetail(fmt.Sprintf("student_id=%s", students[i].StudentID), a.App.CurrentPeerIndex, a.Config.Peers)
		if err != nil {
			if err.Error() == "does not exists" {
				continue
			}
			return errors.New(string(evidence))
		}

		ctx.Logger.Printf("student[%d]: %+v\n", i, collected)

		students[i].Collected = collected
		students[i].Evidence = evidence
		a.App.CurrentPeerIndex = index
	}

	data, err := json.Marshal(students)
	if err != nil {
		return err
	}

	if _, err = w.Write(data); err != nil {
		return err
	}

	return nil
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
	if ctx.User.Group != "st_group" {
		return errors.New("only students are allowed")
	}

	url, err := ctx.ShareProfile(ctx.User.UserID)
	if err != nil {
		return err
	}

	w.Write([]byte(url))

	return nil
}

// ViewProfile view a specific student's profile
func (a *API) ViewProfile(w http.ResponseWriter, r *http.Request) {
	ctx := a.App.NewContext().WithRemoteAddress(r.RemoteAddr)
	ctx.Logger.Infoln(r.RemoteAddr, r.RequestURI)

	vars := mux.Vars(r)
	url, ok := vars["url"]
	if !ok {
		http.Error(w, "incorrect url", http.StatusNotFound)
		return
	}

	ctx.Logger.Printf("url: %s", url)

	student, err := ctx.ViewProfile(w, url, a.App.CurrentPeerIndex, a.Config.Peers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if student == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	studentBytes, err := json.Marshal(student)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(studentBytes)
}

// EditProfile handle a editing student's information request
func (a *API) EditProfile(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	/*
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return err
		}

		var edit model.EditProfile
		if err := json.Unmarshal(body, &edit); err != nil {
			return err
		}
	*/

	var edit model.EditProfile

	if err := r.ParseMultipartForm(20 << 20); err != nil {
		return err
	}

	if image, header, err := r.FormFile("profile_picture"); err == nil {
		ctx.Logger.Infof("Upload File: %+v\nFile size: %+v\nMIME Header: %+v\n", header.Filename, header.Size, header.Header)
		edit.ProfilePicture, err = a.App.S3.UploadProfilePicture(image, header, ctx.User.UserID)
		if err != nil {
			return err
		}
	}

	edit.Motto = r.FormValue("motto")

	if err := ctx.UpdateStudentProfile(edit.ProfilePicture, edit.Motto); err != nil {
		return err
	}

	editBytes, err := json.Marshal(edit)
	if err != nil {
		return err
	}

	w.Write(editBytes)

	return nil
}
