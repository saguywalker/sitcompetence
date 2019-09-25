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

// GetCompetences returns all of activities
func (ctx *Context) GetCompetences() ([]*model.Competence, error) {
	competences, err := ctx.Database.GetCompetences()
	if err != nil {
		return nil, err
	}

	return competences, nil
}

// CreateCompetence creates new competence
func (ctx *Context) CreateCompetence(competence *model.Competence) error {
	return ctx.Database.CreateCompetence(competence)
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
