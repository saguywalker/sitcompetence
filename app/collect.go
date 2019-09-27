package app

import (
	"github.com/saguywalker/sitcompetence/model"
)

/*

// GetCollectedCompetences return all competences collected by all students
func (ctx *Context) GetCollectedCompetences() ([]*model.Competence, error) {
	return nil, fmt.Errorf("unimplemented")
}

// GetCollectedByCompetenceID returns competence struct from competence id
func (ctx *Context) GetCollectedByCompetenceID(id uint16) ([]model.Student, error) {
	competence, err := ctx.Database.GetCollectedByCompetenceID(id)
	if err != nil {
		return nil, err
	}

	return competence, fmt.Errorf("unimplemented")
}
*/

// GetCollectedByStudentID returns all of activities
func (ctx *Context) GetCollectedByStudentID(id string, pageNo uint64) ([]model.Competence, error) {
	competencesID, err := ctx.Database.GetCompetencesIDByStudentID(id, ctx.PageLimit, pageNo)
	if err != nil {
		return nil, err
	}

	var competences []model.Competence

	for _, competenceID := range competencesID {
		competence, err := ctx.Database.GetCompetenceByID(competenceID)
		if err != nil {
			return nil, err
		}
		competences = append(competences, *competence)
	}

	return competences, nil
}

// CreateCollectedCompetence update new competence and its transactionID to a corresponding studentID
func (ctx *Context) CreateCollectedCompetence(badges []*model.CollectedCompetence, txID []byte) error {
	for _, badge := range badges {
		if err := ctx.Database.CreateCollectedCompetence(badge); err != nil {
			return err
		}
	}
	return nil
}
