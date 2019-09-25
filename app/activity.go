package app

import "github.com/saguywalker/sitcompetence/model"

// GetActivityByID returns activity struct from activity id
func (ctx *Context) GetActivityByID(id uint32) (*model.FullActivity, error) {
	activity, err := ctx.Database.GetActivityByID(id)
	if err != nil {
		return nil, err
	}

	competences, err := ctx.Database.GetCompetencesByActivityID(id)
	if err != nil {
		return nil, err
	}

	students, err := ctx.Database.GetAttendedActivityByID(id)
	if err != nil {
		return nil, err
	}

	return &model.FullActivity{ActivityDetail: activity, Competences: competences, Students: students}, nil
}

// GetActivities returns all of activities
func (ctx *Context) GetActivities() ([]*model.Activity, error) {
	activities, err := ctx.Database.GetActivities()
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(activities); i++ {
		competences, err := ctx.Database.GetCompetencesByActivityID(activities[i].ActivityID)
		if err != nil {
			return nil, err
		}

		activities[i].Competences = competences
	}

	return activities, err
}

// GetActivitiesByStaff return activities from staff id
func (ctx *Context) GetActivitiesByStaff(id string) ([]*model.Activity, error) {
	activities, err := ctx.Database.GetActivitiesByStaff(id)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(activities); i++ {
		competences, err := ctx.Database.GetCompetencesByActivityID(activities[i].ActivityID)
		if err != nil {
			return nil, err
		}

		activities[i].Competences = competences
	}

	return activities, err
}

// GetActivitiesByStudent return activities from student id
func (ctx *Context) GetActivitiesByStudent(id string) ([]*model.Activity, error) {
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

// UpdateActivity update activity from activity id
func (ctx *Context) UpdateActivity(activity *model.Activity) error {
	return ctx.Database.UpdateActivity(activity)
}

// DeleteActivity delete activity from activity id
func (ctx *Context) DeleteActivity(id uint32) error {
	return ctx.Database.DeleteActivity(id)
}
