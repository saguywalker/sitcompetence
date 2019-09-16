package model

import (
	"github.com/cbergoon/merkletree"
)

// Merkle contains merkle root hash and each items in merkle tree
type Merkle struct {
	MerkleRoot []byte               `json:"merkle_root"`
	ItemHash   []merkletree.Content `json:"item_hash"`
}

// NewMerkle creates new transaction struct
func NewMerkle(merkleRoot []byte, itemHash []merkletree.Content) *Merkle {
	return &Merkle{
		MerkleRoot: merkleRoot,
		ItemHash:   itemHash,
	}
}
