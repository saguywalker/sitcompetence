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
func (ctx *Context) GetCompetences() (*[]model.Competence, error) {
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

// CreateCollectedCompetence update new competence and its transactionID to a corresponding studentID
func (ctx *Context) CreateCollectedCompetence(badges []*model.GiveBadge, txID []byte) error {
	for _, badge := range badges {
		tmp := model.NewCollectedCompetence(badge.StudentID, badge.CompetenceID, txID)
		if err := ctx.Database.CreateCollectedCompetence(tmp); err != nil {
			return err
		}
	}
	return nil
}
