package app

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"

	"github.com/cbergoon/merkletree"

	"github.com/saguywalker/sitcompetence/model"
)

/*
// UpdateMerkleTransaction insert a new transactionID and merkleRootHash into transaction table
// It also add a merkleRoot and all of its transaction into merkle table
func (ctx *Context) UpdateMerkleTransaction(transactionID, merkleRoot []byte, transactionSet []merkletree.Content) (*model.TransactionLink, error) {
	merkle := model.NewMerkle(merkleRoot, transactionSet)
	if err := ctx.Database.CreateMerkle(merkle); err != nil {
		return nil, err
	}

	transaction := model.NewTransactionLink(transactionID, merkleRoot)
	if err := ctx.Database.CreateTransaction(transaction); err != nil {
		return nil, err
	}

	return transaction, nil
}
*/

// VerifyTX verify data with a given merkle root
func (ctx *Context) VerifyTX(merkleRoot, data []byte) (bool, error) {
	var trimData bytes.Buffer
	if err := json.Compact(&trimData, data); err != nil {
		return false, err
	}
	hashData := sha256.Sum256(trimData.Bytes())
	ctx.Logger.Infof("H(data): %x\n", hashData)

	items, err := ctx.Database.GetItemsByMerkleRoot(merkleRoot)
	if err != nil {
		return false, err
	}
	for i, x := range items {
		ctx.Logger.Printf("%d: %x\n", i, x)
	}

	itemContents := make([]merkletree.Content, len(items))
	for i, item := range items {
		itemContents[i] = model.MyHash(item)
	}

	items = nil

	tree, err := merkletree.NewTree(itemContents)
	if err != nil {
		return false, err
	}

	if !bytes.Equal(tree.MerkleRoot(), merkleRoot) {
		return false, fmt.Errorf("merkleroots are not equal (%x) and (%x)", tree.MerkleRoot(), merkleRoot)
	}

	vc, err := tree.VerifyContent(model.MyHash(hashData[:]))
	if err != nil {
		return false, err
	}

	return vc, nil
}
