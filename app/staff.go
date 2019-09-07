package app

import "github.com/saguywalker/sitcompetence/model"

func (ctx *Context) GetStaffByID(id string) (*model.Staff, error) {
	staff, err := ctx.Database.GetStaffByID(id)
	if err != nil {
		return nil, err
	}

	return staff, nil
}

func (ctx *Context) GetStaffs() ([]*model.Staff, error) {
	staffs, err := ctx.Database.GetStaffs()
	if err != nil {
		return nil, err
	}

	return staffs, nil
}

func (ctx *Context) CreateStaff(staff *model.Staff) error {
	return ctx.Database.CreateStaff(staff)
}
