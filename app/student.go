package app

import "github.com/saguywalker/sitcompetence/model"

func (ctx *Context) GetStudentByID(id string) (*model.Student, error) {
	student, err := ctx.Database.GetStudentByID(id)
	if err != nil {
		return nil, err
	}

	return student, nil
}

func (ctx *Context) GetStudents() ([]*model.Student, error) {
	students, err := ctx.Database.GetStudents()
	if err != nil {
		return nil, err
	}

	return students, nil
}

func (ctx *Context) CreateStudent(student *model.Student) error {
	return ctx.Database.CreateStudent(student)
}
