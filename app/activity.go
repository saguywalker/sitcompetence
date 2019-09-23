package app

import "github.com/saguywalker/sitcompetence/model"

// GetActivityByID returns activity struct from activity id
func (ctx *Context) GetActivityByID(id uint32) (*model.Activity, error) {
	activity, err := ctx.Database.GetActivityByID(id)
	if err != nil {
		return nil, err
	}

	competences, err := ctx.Database.GetCompetencesByActivityID(id)
	if err != nil {
		return nil, err
	}

	activity.Competences = competences

	return activity, nil
}

// GetActivities returns all of activities
func (ctx *Context) GetActivities() (*[]model.Activity, error) {
	activities, err := ctx.Database.GetActivities()
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(activites); i++ {
		competences, err := ctx.Database.GetCompetencesByActivityID(activites[i].ActivityID)
		if err != nil {
			return nil, err
		}

		activites[i].Competences = competences
	}

	return activites, erra
}

func (ctx *Context) GetActivitiesByStaff(id string) (*[]model.Activity, error) {
	activities, err := ctx.Database.GetActivitiesByStaff(id)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(activites); i++ {
		competences, err := ctx.Database.GetCompetencesByActivityID(activites[i].ActivityID)
		if err != nil {
			return nil, err
		}

		activites[i].Competences = competences
	}

	return activites, err
}

func (ctx *Context) GetActivitiesByStudent(id string) (*[]model.Activity, error) {
	activites, err := ctx.Database.GetActivitiesByStudent(id)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(activites); i++ {
		competences, err := ctx.Database.GetCompetencesByActivityID(activites[i].ActivityID)
		if err != nil {
			return nil, err
		}

		activites[i].Competences = competences
	}

	return activites, err
}

// CreateActivity creates new activity
func (ctx *Context) CreateActivity(activity *model.Activity) error {
	return ctx.Database.CreateActivity(activity)
}

/*
func (ctx *Context) UpdateActivity(id uint32) error {
	return ctx.Database.UpdateActivity(id)
}
*/
func (ctx *Context) DeleteActivity(id uint32) error {
	return ctx.Database.DeleteActivity(id)
}
