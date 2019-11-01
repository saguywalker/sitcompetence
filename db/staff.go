package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/saguywalker/sitcompetence/model"
)

// GetStaffByID returns an staff from staffID
func (db *Database) GetStaffByID(id string) (*model.Staff, error) {
	log.Println(fmt.Sprintf("SELECT staffId, firstname FROM staff WHERE staffId = %s", id))

	row, err := db.Query("SELECT staffId, firstname FROM staff WHERE staffId = $1", id)
	if err != nil {
		return nil, err
	}

	var staff model.Staff

	for row.Next() {
		err := row.Scan(&staff.StaffID, &staff.FirstName)
		if err != nil {
			return nil, err
		}
	}

	return &staff, nil
}

// GetStaffs returns all staffs in a table
func (db *Database) GetStaffs(pageLimit uint32, pageNo uint32) ([]*model.Staff, error) {
	var rows *sql.Rows
	var err error

	if pageLimit == 0 || pageNo == 0 {
		log.Println("SELECT staffId, firstname FROM staff")
		rows, err = db.Query("SELECT staffId, firstname FROM staff")
	} else {
		log.Println(fmt.Sprintf("SELECT staffId, firstname FROM staff ORDER BY staffId LIMIT %d OFFSET %d", pageLimit, (pageNo-1)*pageLimit))
		rows, err = db.Query("SELECT staffId, firstname FROM staff ORDER BY staffId LIMIT $1 OFFSET $2", pageLimit, (pageNo-1)*pageLimit)
	}
	if err != nil {
		return nil, err
	}

	var staffs []*model.Staff

	for rows.Next() {
		var staff model.Staff
		err := rows.Scan(&staff.StaffID, &staff.FirstName)
		if err != nil {
			return nil, err
		}
		staffs = append(staffs, &staff)
	}

	return staffs, nil
}

// GetStaffPublicKey return staff public key from staff id
func (db *Database) GetStaffPublicKey(id string) ([]byte, error) {
	row, err := db.Query("SELECT publicKey FROM staff WHERE staffId = $1", id)
	if err != nil {
		return nil, err
	}

	publicKey := make([]byte, 32)

	for row.Next() {
		if err := row.Scan(&publicKey); err != nil {
			return nil, err
		}
	}

	return publicKey, nil
}

// SetPubkey update database
func (db *Database) SetPubkey(id string, pubkey []byte) error {
	stmt, err := db.Prepare("UPDATE staff set publicKey=$1 WHERE staffId=$2")
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(pubkey, id); err != nil {
		return err
	}

	return nil
}

// CheckKey check publickey
func (db *Database) CheckKey(id string) (bool, error) {
	row, err := db.Query("SELECT LENGTH(publicKey) FROM staff WHERE staffId = $1", id)
	if err != nil {
		return false, err
	}

	var keyLength uint

	for row.Next() {
		if err := row.Scan(&keyLength); err != nil {
			return false, err
		}
	}

	return keyLength == 0, nil
}

// CreateStaff inserts a new staff
func (db *Database) CreateStaff(s *model.Staff) (string, error) {
	var id string
	command := "INSERT INTO staff(staffId, firstname, lastname) VALUES($1, $2, $3) RETURNING staffId;"
	if err := db.QueryRow(command, s.StaffID, s.FirstName, s.LastName).Scan(&id); err != nil {
		return "", err
	}

	return id, nil
}

// UpdateStaff update staff from staff id
func (db *Database) UpdateStaff(staff *model.Staff) error {
	stmt, err := db.Prepare("UPDATE staff set firstname=$1, lastname=$2 " +
		"WHERE staffId=$3")

	if err != nil {
		return err
	}

	_, err = stmt.Exec(staff.FirstName, staff.LastName, staff.StaffID)
	if err != nil {
		return err
	}

	return nil
}

// DeleteStaff deletes a staff from staffID
func (db *Database) DeleteStaff(staffID string) error {
	stmt, err := db.Prepare("DELETE FROM staff WHERE staffId = $1")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(staffID)
	if err != nil {
		return err
	}

	return nil
}
