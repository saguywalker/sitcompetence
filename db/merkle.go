package db

import "github.com/saguywalker/sitcompetence/model"

// CreateMerkle insert new merkle into merkle table
func (db *Database) CreateMerkle(merkle *model.Merkle) error {
	stmt, err := db.Prepare("INSERT INTO merkle(merkleRoot, itemHash) VALUES(?, ?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(merkle.MerkleRootHash, merkle.ItemHash)
	if err != nil {
		return err
	}

	return nil
}

// GetMerkles returns all merkles in table
func (db *Database) GetMerkles() (*[]model.Merkle, error) {
	rows, err := db.Query("SELECT * FROM merkle")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var merkles []model.Merkle

	for rows.Next() {
		var merkle model.Merkle
		err = rows.Scan(&merkle.MerkleRootHash, &merkle.ItemHash)
		if err != nil {
			return nil, err
		}
		merkles = append(merkles, merkle)
	}

	return &merkles, nil
}

// GetMerkleByID returns all transaction in corresponding merkle root
func (db *Database) GetMerkleByID(merkleRoot []byte) (*model.Merkle, error) {
	row, err := db.Query("SELECT * FROM merkle WHERE merkleRootHash = ?", merkleRoot)
	if err != nil {
		return nil, err
	}

	var merkle model.Merkle

	for row.Next() {
		var root []byte
		var item []byte
		err := row.Scan(&root, &item)
		if err != nil {
			return nil, err
		}

		merkle.ItemHash = append(merkle.ItemHash, item)
	}

	merkle.MerkleRootHash = merkleRoot

	return &merkle, nil
}

// DeleteMerkle deletes all merkle root and its transactions
func (db *Database) DeleteMerkle(merkleRoot []byte) error {
	stmt, err := db.Prepare("DELETE FROM merkle WHERE merkleRootHash = ?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(merkleRoot)
	if err != nil {
		return err
	}

	return nil
}
