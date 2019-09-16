package model

// TransactionLink table in Postgres
type TransactionLink struct {
	TransactionID []byte `json:"transaction_id"`
	MerkleRoot    []byte `json:"merkle_root"`
}

// NewTransactionLink creates new transaction struct
func NewTransactionLink(transactionID, merkleRoot []byte) *TransactionLink {
	return &TransactionLink{
		TransactionID: transactionID,
		MerkleRoot:    merkleRoot,
	}
}
