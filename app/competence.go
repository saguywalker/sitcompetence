package app

import "github.com/saguywalker/sitcompetence/model"

func (ctx *Context) GetCompetenceByID(id uint16) (*model.Competence, error) {
	competence, err := ctx.Database.GetCompetenceByID(id)
	if err != nil {
		return nil, err
	}

	return competence, nil
}

func (ctx *Context) GetCompetences() ([]*model.Competence, error) {
	competences, err := ctx.Database.GetCompetences()
	if err != nil {
		return nil, err
	}

	return competences, nil
}

func (ctx *Context) CreateCompetence(competence *model.Competence) error {
	return ctx.Database.CreateCompetence(competence)
}
