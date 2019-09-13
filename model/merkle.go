package model

// Merkle contains merkle root hash and each items in merkle tree
type Merkle struct {
	MerkleRoot []byte   `json:"merkle_root"`
	ItemHash   [][]byte `json:"item_hash"`
}

// NewMerkle creates new transaction struct
func NewMerkle(merkleRoot []byte, itemHash [][]byte) *Merkle {
	return &Merkle{
		MerkleRoot: merkleRoot,
		ItemHash:   itemHash,
	}
}
