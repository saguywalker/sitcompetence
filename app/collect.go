package app

import (
	"fmt"

	"github.com/saguywalker/sitcompetence/model"
)

func (ctx *Context) GetCollectedCompetences() ([]*model.Competence, error) {
	return nil, fmt.Errorf("unimplemented")
}

// GetCollectedByCompetenceID returns competence struct from competence id
func (ctx *Context) GetCollectedByCompetenceID(id uint16) ([]*model.Student, error) {
	competence, err := ctx.Database.GetCollectedByCompetenceID(id)
	if err != nil {
		return nil, err
	}

	return competence, fmt.Errorf("unimplemented")
}

// GetCollectedByStudentID returns all of activities
func (ctx *Context) GetCollectedByStudentID(id string) ([]*model.Competence, error) {
	competences, err := ctx.Database.GetCollectedByStudentID(id)
	if err != nil {
		return nil, err
	}

	return competences, fmt.Errorf("unimplemented")
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
