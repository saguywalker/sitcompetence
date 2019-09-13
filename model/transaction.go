package model

// TransactionLink table in Postgres
type TransactionLink struct {
	TransactionID  []byte `json:"transaction_id"`
	MerkleRootHash []byte `json:"merkle_root"`
}

// NewTransactionLink creates new transaction struct
func NewTransactionLink(transactionID, merkleRootHah []byte) *TransactionLink {
	return &TransactionLink{
		TransactionID:  transactionID,
		MerkleRootHash: merkleRootHah,
	}
}
