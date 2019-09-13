package model

// Merkle contains merkle root hash and each items in merkle tree
type Merkle struct {
	MerkleRootHash []byte   `json:"merkle_root"`
	ItemHash       [][]byte `json:"item_hash"`
}

// NewMerkle creates new transaction struct
func NewMerkle(merkleRoot []byte, itemHash [][]byte) *Merkle {
	return &Merkle{
		MerkleRootHash: merkleRoot,
		ItemHash:       itemHash,
	}
}
