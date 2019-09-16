package model

// Merkle contains merkle root hash and each items in merkle tree
type Merkle struct {
	MerkleRoot string   `json:"merkle_root"`
	ItemHash   []string `json:"item_hash"`
}

// NewMerkle creates new transaction struct
func NewMerkle(merkleRoot string, itemHash []string) *Merkle {
	return &Merkle{
		MerkleRoot: merkleRoot,
		ItemHash:   itemHash,
	}
}
