package app

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/saguywalker/sitcompetence/model"
)

// GetStudentByID returns student detail from student id
func (ctx *Context) GetStudentByID(id string) (*model.Student, error) {
	student, err := ctx.Database.GetStudentByID(id)
	if err != nil {
		return student, err
	}

	return student, nil
}

// GetStudents returns all of students
func (ctx *Context) GetStudents(pageNo uint32, dp, year string) ([]*model.Student, error) {
	var yr uint64
	var err error

	if year == "" {
		yr = 0
	} else {
		yr, err = strconv.ParseUint(year, 0, 16)
		if err != nil {
			return nil, err
		}
	}

	students, err := ctx.Database.GetStudents(ctx.PageLimit, pageNo, dp, uint16(yr))
	if err != nil {
		//ctx.Logger.Errorln(err.Error())
		return nil, err
	}

	return students, nil
}

// CreateStudent create new student
func (ctx *Context) CreateStudent(student *model.Student) (string, error) {
	return ctx.Database.CreateStudent(student)
}

// UpdateStudent update student
func (ctx *Context) UpdateStudent(student *model.Student) error {
	if err := ctx.Database.UpdateStudent(student); err != nil {
		return err
	}

	return nil
}

// DeleteStudent delete student from activity id
func (ctx *Context) DeleteStudent(studentID string) error {
	return ctx.Database.DeleteStudent(studentID)
}

// ShareProfile generate url for a student's portfolio
func (ctx *Context) ShareProfile(studentID string) (string, error) {
	expire := time.Now().Add(time.Hour * 24 * 7)
	expireBytes, err := expire.MarshalBinary()
	if err != nil {
		return "", err
	}

	hashed := sha256.New()
	hashed.Write([]byte(studentID))
	hashed.Write(expireBytes)
	url := hex.EncodeToString(hashed.Sum(nil))

	if err := ctx.Database.UpdateShareProfile(studentID, expire, url); err != nil {
		return "", err
	}

	return url, nil
}

// ViewProfile view a profile from url
func (ctx *Context) ViewProfile(w http.ResponseWriter, url string, index uint64, peers []string) ([]byte, error) {
	if err := ctx.Database.CheckExpire(url); err != nil {
		return nil, err
	}

	student, err := ctx.Database.GetStudentURL(url)
	if err != nil {
		return nil, err
	}

	ctx.Logger.Infof("after queried: %+v\n", student)

	if len(student.StudentID) == 0 {
		return nil, nil
	}

	collected, evidence, index, err := ctx.GetCollectedWithDetail(fmt.Sprintf("student_id=%s", student.StudentID), index, peers)
	if err != nil {
		return nil, err
	}

	student.Collected = collected
	student.Evidence = evidence

	collectedBytes, err := json.Marshal(student)
	if err != nil {
		return nil, err
	}

	return collectedBytes, nil
}

// UpdateProfilePicture save an image into static-images directory
func (ctx *Context) UpdateProfilePicture(filePath string) error {
	return ctx.Database.UpdateProfilePicture(ctx.User.UserID, filePath)
}
