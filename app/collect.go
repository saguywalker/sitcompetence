package app

import (
	"bytes"
	"encoding/json"

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
/*
// GetCollectedByStudentID returns all of activities
func (ctx *Context) GetCollectedByStudentID(id string, pageNo uint32) ([]model.Competence, error) {
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
*/

// GetCollectedWithDetail return list of collected competence from student id
func (ctx *Context) GetCollectedWithDetail(id string, index uint64, peers []string) ([]model.CollectedCompetence, uint64, error) {
	// collected, err := ctx.Database.GetCompetencesByStudentID(id, ctx.PageLimit, pageNo)
	collected, returnIndex, err := ctx.BlockchainQueryWithParams(id, index, peers)
	if err != nil {
		return nil, returnIndex, err
	}

	var returnCollected []model.CollectedCompetence
	if err := json.Unmarshal(collected, &returnCollected); err != nil {
		parts := bytes.Split(collected, []byte("|"))
		for _, x := range parts {
			var tmp model.CollectedCompetence
			if err := json.Unmarshal(x, &tmp); err != nil {
				return nil, returnIndex, err
			}

			returnCollected = append(returnCollected, tmp)
		}
	}

	ctx.Logger.Infof("from blockchain: %+v\n", returnCollected)

	return returnCollected, returnIndex, nil
}

/*
// CreateCollectedCompetence update new competence and its transactionID to a corresponding studentID
func (ctx *Context) CreateCollectedCompetence(badges []*model.CollectedCompetence, txID []byte) error {
	for _, badge := range badges {
		if err := ctx.Database.CreateCollectedCompetence(badge); err != nil {
			return err
		}
	}
	return nil
}
*/
