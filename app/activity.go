package app

import "github.com/saguywalker/sitcompetence/model"

// GetActivityByID returns activity struct from activity id
func (ctx *Context) GetActivityByID(id uint32) (*model.Activity, error) {
	activity, err := ctx.Database.GetActivityByID(id)
	if err != nil {
		return activity, err
	}

	return activity, nil
}

// GetActivities returns all of activities
func (ctx *Context) GetActivities() (*[]model.Activity, error) {
	activities, err := ctx.Database.GetActivities()
	if err != nil {
		return nil, err
	}

	return activities, nil
}

// CreateActivity creates new activity
func (ctx *Context) CreateActivity(activity *model.Activity) error {
	return ctx.Database.CreateActivity(activity)
}
