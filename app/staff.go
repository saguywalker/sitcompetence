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

// SetPubkey set publickey
func (ctx *Context) SetPubkey(id string, pubkey []byte) error {
	return ctx.Database.SetPubkey(id, pubkey)
}

// CheckKey check publickey
func (ctx *Context) CheckKey(id string) (bool, error) {
	ok, err := ctx.Database.CheckKey(id)
	if err != nil {
		return false, err
	}

	ctx.Logger.Infoln(ok)
	return ok, nil
}

// GetStaffs returns all of activities
func (ctx *Context) GetStaffs(pageNo uint32) ([]*model.Staff, error) {
	staffs, err := ctx.Database.GetStaffs(ctx.PageLimit, pageNo)
	if err != nil {
		//ctx.Logger.Errorln(err.Error())
		return nil, err
	}

	return staffs, nil
}

// CreateStaff creates new staff
func (ctx *Context) CreateStaff(staff *model.Staff) (string, error) {
	return ctx.Database.CreateStaff(staff)
}

// UpdateStaff update staff
func (ctx *Context) UpdateStaff(staff *model.Staff) error {
	if err := ctx.Database.UpdateStaff(staff); err != nil {
		return err
	}

	return nil
}

// DeleteStaff delete student from activity id
func (ctx *Context) DeleteStaff(staffID string) error {
	return ctx.Database.DeleteStaff(staffID)
}
