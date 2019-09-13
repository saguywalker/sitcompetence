package app

import "github.com/saguywalker/sitcompetence/model"

// GetStudentByID returns student detail from student id
func (ctx *Context) GetStudentByID(id string) (*model.Student, error) {
	student, err := ctx.Database.GetStudentByID(id)
	if err != nil {
		return student, err
	}

	return student, nil
}

// GetStudents returns all of students
func (ctx *Context) GetStudents() (*[]model.Student, error) {
	students, err := ctx.Database.GetStudents()
	if err != nil {
		return nil, err
	}

	return students, nil
}

// CreateStudent create new student
func (ctx *Context) CreateStudent(student *model.Student) error {
	return ctx.Database.CreateStudent(student)
}
