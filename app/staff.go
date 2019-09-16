package app

import "github.com/saguywalker/sitcompetence/model"

// GetStaffByID returns staff struct from staff id
func (ctx *Context) GetStaffByID(id string) (*model.Staff, error) {
	staff, err := ctx.Database.GetStaffByID(id)
	if err != nil {
		//ctx.Logger.Errorln(err.Error())
		return nil, err
	}

	return staff, nil
}

// GetStaffs returns all of activities
func (ctx *Context) GetStaffs() (*[]model.Staff, error) {
	staffs, err := ctx.Database.GetStaffs()
	if err != nil {
		//ctx.Logger.Errorln(err.Error())
		return nil, err
	}

	return staffs, nil
}

// CreateStaff creates new staff
func (ctx *Context) CreateStaff(staff *model.Staff) error {
	if err := ctx.Database.CreateStaff(staff); err != nil {
		//ctx.Logger.Errorln(err.Error())
		return err
	}

	return nil
}
