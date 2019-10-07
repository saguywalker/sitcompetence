package app

import "github.com/saguywalker/sitcompetence/model"

// GetCompetenceByID returns competence struct from competence id
func (ctx *Context) GetCompetenceByID(id uint16) (*model.Competence, error) {
	competence, err := ctx.Database.GetCompetenceByID(id)
	if err != nil {
		return nil, err
	}

	return competence, nil
}

// GetCompetencesByActivityID query competence from activity id
func (ctx *Context) GetCompetencesByActivityID(activityID uint32, pageNo uint32) ([]model.Competence, error) {
	competences, err := ctx.Database.GetCompetencesByActivityID(activityID, ctx.PageLimit, pageNo)
	if err != nil {
		return nil, err
	}

	return competences, nil
}

// GetCompetences returns all of activities
func (ctx *Context) GetCompetences(pageNo uint32) ([]model.Competence, error) {
	competences, err := ctx.Database.GetCompetences(ctx.PageLimit, pageNo)
	if err != nil {
		return nil, err
	}

	return competences, nil
}

// CreateCompetence creates new competence
func (ctx *Context) CreateCompetence(competence *model.Competence) (int64, error) {
	id, err := ctx.CreateCompetence(competence)
	if err != nil {
		return -1, err
	}

	return id, nil
}

// UpdateCompetence update competence
func (ctx *Context) UpdateCompetence(competence *model.Competence) error {
	if err := ctx.Database.UpdateCompetence(competence); err != nil {
		return err
	}

	return nil
}

// DeleteCompetence delete competence from competence id
func (ctx *Context) DeleteCompetence(competenceID uint16) error {
	return ctx.Database.DeleteCompetence(competenceID)
}
