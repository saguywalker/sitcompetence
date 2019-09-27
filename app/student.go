package app

import (
	"strconv"

	"github.com/saguywalker/sitcompetence/model"
)

// GetStudentByID returns student detail from student id
func (ctx *Context) GetStudentByID(id string) (*model.Student, error) {
	student, err := ctx.Database.GetStudentByID(id)
	if err != nil {
		//ctx.Logger.Errorln(err.Error())
		return student, err
	}

	return student, nil
}

// GetStudents returns all of students
func (ctx *Context) GetStudents(pageNo uint64, dp, semester, year string) ([]*model.Student, error) {
	var semes uint64
	var yr uint64
	var err error

	if semester == "" {
		semes = 0
	} else {
		semes, err = strconv.ParseUint(semester, 0, 64)
		if err != nil {
			return nil, err
		}
	}

	if year == "" {
		yr = 0
	} else {
		yr, err = strconv.ParseUint(year, 0, 64)
		if err != nil {
			return nil, err
		}
	}

	students, err := ctx.Database.GetStudents(ctx.PageLimit, pageNo, dp, uint16(semes), uint16(yr))
	if err != nil {
		//ctx.Logger.Errorln(err.Error())
		return nil, err
	}

	return students, nil
}

// CreateStudent create new student
func (ctx *Context) CreateStudent(student *model.Student) error {
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
