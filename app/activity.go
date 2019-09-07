package app

import "github.com/saguywalker/sitcompetence/model"

func (ctx *Context) GetActivityByID(id uint32) (*model.Activity, error) {
	activity, err := ctx.Database.GetActivityByID(id)
	if err != nil {
		return nil, err
	}

	return activity, nil
}

func (ctx *Context) GetActivities() ([]*model.Activity, error) {
	activities, err := ctx.Database.GetActivities()
	if err != nil {
		return nil, err
	}

	return activities, nil
}

func (ctx *Context) CreateActivity(activity *model.Activity) error {
	return ctx.Database.CreateActivity(activity)
}
