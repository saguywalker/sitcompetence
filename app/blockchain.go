package app

import (
	"github.com/saguywalker/sitcompetence/model"
)

// UpdateMerkleTransaction insert a new transactionID and merkleRootHash into transaction table
// It also add a merkleRoot and all of its transaction into merkle table
func (ctx *Context) UpdateMerkleTransaction(transactionID, merkleRoot []byte, transactionSet [][]byte) error {
	transaction := model.NewTransactionLink(transactionID, merkleRoot)
	if err := ctx.Database.CreateTransaction(transaction); err != nil {
		return err
	}

	merkle := model.NewMerkle(merkleRoot, transactionSet)
	if err := ctx.Database.CreateMerkle(merkle); err != nil {
		return err
	}

	return nil
}
