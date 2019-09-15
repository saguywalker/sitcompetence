package model

// TransactionLink table in Postgres
type TransactionLink struct {
	TransactionID string `json:"transaction_id"`
	MerkleRoot    string `json:"merkle_root"`
}

// NewTransactionLink creates new transaction struct
func NewTransactionLink(transactionID, merkleRoot string) *TransactionLink {
	return &TransactionLink{
		TransactionID: transactionID,
		MerkleRoot:    merkleRoot,
	}
}
