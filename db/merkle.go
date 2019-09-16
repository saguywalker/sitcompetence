package db

import (
	"github.com/saguywalker/sitcompetence/model"
)

// CreateMerkle insert new merkle into merkle table
func (db *Database) CreateMerkle(merkle *model.Merkle) error {
	for _, item := range merkle.ItemHash {
		stmt, err := db.Prepare("INSERT INTO merkle(merkleRoot, itemHash) VALUES($1, $2)")
		if err != nil {
			return err
		}

		_, err = stmt.Exec(merkle.MerkleRoot, item)
		if err != nil {
			return err
		}
	}

	return nil
}

// GetMerkles returns all merkles in table
func (db *Database) GetMerkles() ([]*model.Merkle, error) {
	rows, err := db.Query("SELECT * FROM merkle")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var merkles []*model.Merkle

	for rows.Next() {
		var merkle model.Merkle
		err = rows.Scan(&merkle.MerkleRoot, &merkle.ItemHash)
		if err != nil {
			return nil, err
		}
		merkles = append(merkles, &merkle)
	}

	return merkles, nil
}

// GetItemsByMerkleRoot returns all transaction in corresponding merkle root
func (db *Database) GetItemsByMerkleRoot(merkleRoot []byte) ([][]byte, error) {
	row, err := db.Query("SELECT itemhash FROM merkle WHERE merkleRoot = $1", merkleRoot)
	if err != nil {
		return nil, err
	}

	var items [][]byte

	for row.Next() {
		var item []byte
		err := row.Scan(&item)
		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return items, nil
}

// DeleteMerkle deletes all merkle root and its transactions
func (db *Database) DeleteMerkle(merkleRoot []byte) error {
	stmt, err := db.Prepare("DELETE FROM merkle WHERE merkleRoot = $1")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(merkleRoot)
	if err != nil {
		return err
	}

	return nil
}
