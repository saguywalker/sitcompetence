package db

import (
	"github.com/saguywalker/sitcompetence/model"
)

// GetStaffByID returns an staff from staffID
func (db *Database) GetStaffByID(id string) (*model.Staff, error) {
	row, err := db.Query("SELECT * FROM staff WHERE staffId = ?", id)
	if err != nil {
		return nil, err
	}

	var staff model.Staff

	for row.Next() {
		err := row.Scan(&staff.StaffID, &staff.FirstName, &staff.LastName)
		if err != nil {
			return nil, err
		}
	}

	return &staff, nil
}

// GetStaffs returns all staffs in a table
func (db *Database) GetStaffs() (*[]model.Staff, error) {
	rows, err := db.Query("SELECT * FROM staff")
	if err != nil {
		return nil, err
	}

	var staffs []model.Staff

	for rows.Next() {
		var staff model.Staff
		err := rows.Scan(&staff.StaffID, &staff.FirstName, &staff.LastName)
		if err != nil {
			return nil, err
		}
		staffs = append(staffs, staff)
	}

	return &staffs, nil
}

// CreateStaff inserts a new staff
func (db *Database) CreateStaff(staff *model.Staff) error {
	stmt, err := db.Prepare("INSERT INTO staff(staffId, firstName, lastName) VALUES(?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(staff.StaffID, staff.FirstName, staff.LastName)
	if err != nil {
		return err
	}

	return nil
}

// DeleteStaff deletes a staff from staffID
func (db *Database) DeleteStaff(staffID []byte) error {
	stmt, err := db.Prepare("DELETE FROM staff WHERE staffId = ?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(staffID)
	if err != nil {
		return err
	}

	return nil
}
