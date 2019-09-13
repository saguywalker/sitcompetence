package db

import "github.com/saguywalker/sitcompetence/model"

// CreateTransaction insert new transaction into transaction table
func (db *Database) CreateTransaction(transaction *model.TransactionLink) error {
	stmt, err := db.Prepare("INSERT INTO transaction(transactionID, merkleRootHash) VALUES(?, ?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(transaction.TransactionID, transaction.MerkleRootHash)
	if err != nil {
		return err
	}

	return nil
}

// GetTransactions returns all transactions from transaction table
func (db *Database) GetTransactions() ([]model.TransactionLink, error) {
	rows, err := db.Query("SELECT * FROM transaction")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactionSet []model.TransactionLink

	for rows.Next() {
		var transaction model.TransactionLink
		err = rows.Scan(&transaction.TransactionID, &transaction.MerkleRootHash)
		if err != nil {
			return nil, err
		}
		transactionSet = append(transactionSet, transaction)
	}

	return transactionSet, nil
}

// GetTransactionByID returns a transaction from transactionID
func (db *Database) GetTransactionByID(transactionID []byte) (model.TransactionLink, error) {
	var transaction model.TransactionLink

	row, err := db.Query("SELECT * FROM transaction WHERE transactionID = ?", transactionID)
	if err != nil {
		return transaction, err
	}

	for row.Next() {
		err := row.Scan(&transaction.TransactionID, &transaction.MerkleRootHash)
		if err != nil {
			return transaction, err
		}
	}

	return transaction, nil
}

// DeleteTransaction deletes a transaction from transactionID
func (db *Database) DeleteTransaction(transactionID []byte) error {
	stmt, err := db.Prepare("DELETE FROM transaction WHERE transactionID = ?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(transactionID)
	if err != nil {
		return err
	}

	return nil
}
