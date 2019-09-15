package app

import (
	"github.com/saguywalker/sitcompetence/model"
)

// UpdateMerkleTransaction insert a new transactionID and merkleRootHash into transaction table
// It also add a merkleRoot and all of its transaction into merkle table
func (ctx *Context) UpdateMerkleTransaction(transactionID, merkleRoot string, transactionSet []string) error {
	merkle := model.NewMerkle(merkleRoot, transactionSet)
	if err := ctx.Database.CreateMerkle(merkle); err != nil {
		return err
	}

	transaction := model.NewTransactionLink(transactionID, merkleRoot)
	if err := ctx.Database.CreateTransaction(transaction); err != nil {
		return err
	}

	return nil
}

// UpdateCollectedCompetence update new competence and its transactionID to a corresponding studentID
func (ctx *Context) UpdateCollectedCompetence(badges []*model.GiveBadge, txID string) error {
	for _, badge := range badges {
		tmp := model.NewCollectedCompetence(badge.StudentID, badge.CompetenceID, txID)
		if err := ctx.Database.CreateCollectedCompetence(tmp); err != nil {
			return err
		}
	}
	return nil
}
