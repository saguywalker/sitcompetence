package db

import "github.com/saguywalker/sitcompetence/model"

// CreateTransaction insert new transaction into transaction table
func (db *Database) CreateTransaction(transaction *model.TransactionLink) error {
	stmt, err := db.Prepare("INSERT INTO transaction(transactionID, merkleRoot) VALUES(?, ?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(transaction.TransactionID, transaction.MerkleRoot)
	if err != nil {
		return err
	}

	return nil
}

// GetTransactions returns all transactions from transaction table
func (db *Database) GetTransactions() (*[]model.TransactionLink, error) {
	rows, err := db.Query("SELECT * FROM transaction_set")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactionSet []model.TransactionLink

	for rows.Next() {
		var transaction model.TransactionLink
		err = rows.Scan(&transaction.TransactionID, &transaction.MerkleRoot)
		if err != nil {
			return nil, err
		}
		transactionSet = append(transactionSet, transaction)
	}

	return &transactionSet, nil
}

// GetTransactionByID returns a transaction from transactionID
func (db *Database) GetTransactionByID(transactionID []byte) (*model.TransactionLink, error) {
	row, err := db.Query("SELECT * FROM transaction_set WHERE transactionId = ?", transactionID)
	if err != nil {
		return nil, err
	}

	var transaction model.TransactionLink

	for row.Next() {
		err := row.Scan(&transaction.TransactionID, &transaction.MerkleRoot)
		if err != nil {
			return nil, err
		}
	}

	return &transaction, nil
}

// DeleteTransaction deletes a transaction from transactionID
func (db *Database) DeleteTransaction(transactionID []byte) error {
	stmt, err := db.Prepare("DELETE FROM transaction_set WHERE transactionId = ?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(transactionID)
	if err != nil {
		return err
	}

	return nil
}
