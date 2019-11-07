package app

import (
	"bytes"
	"encoding/json"

	"github.com/saguywalker/sitcompetence/model"
)

// GetCollectedWithDetail return list of collected competence from student id
func (ctx *Context) GetCollectedWithDetail(id string, index uint64, peers []string) ([]model.CollectedCompetence, []byte, uint64, error) {
	// collected, err := ctx.Database.GetCompetencesByStudentID(id, ctx.PageLimit, pageNo)
	collected, evidence, returnIndex, err := ctx.BlockchainQueryWithParams(id, index, peers)
	if err != nil {
		return nil, evidence, returnIndex, err
	}

	var returnCollected []model.CollectedCompetence
	if err := json.Unmarshal(collected, &returnCollected); err != nil {
		parts := bytes.Split(collected, []byte("|"))
		for _, x := range parts {
			var tmp model.CollectedCompetence
			if err := json.Unmarshal(x, &tmp); err != nil {
				return nil, evidence, returnIndex, err
			}

			returnCollected = append(returnCollected, tmp)
		}
	}

	ctx.Logger.Infof("from blockchain: %+v\n", returnCollected)

	return returnCollected, evidence, returnIndex, nil
}
