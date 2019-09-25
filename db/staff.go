package db

import (
	"github.com/saguywalker/sitcompetence/model"
)

// GetStaffByID returns an staff from staffID
func (db *Database) GetStaffByID(id string) (*model.Staff, error) {
	row, err := db.Query("SELECT * FROM staff WHERE staffId = $1", id)
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
func (db *Database) GetStaffs() ([]*model.Staff, error) {
	rows, err := db.Query("SELECT * FROM staff")
	if err != nil {
		return nil, err
	}

	var staffs []*model.Staff

	for rows.Next() {
		var staff model.Staff
		err := rows.Scan(&staff.StaffID, &staff.FirstName, &staff.LastName)
		if err != nil {
			return nil, err
		}
		staffs = append(staffs, &staff)
	}

	return staffs, nil
}

// CreateStaff inserts a new staff
func (db *Database) CreateStaff(staff *model.Staff) error {
	stmt, err := db.Prepare("INSERT INTO staff(staffId, firstName, lastName) VALUES($1, $2, $3)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(staff.StaffID, staff.FirstName, staff.LastName)
	if err != nil {
		return err
	}

	return nil
}

// UpdateStaff update staff from staff id
func (db *Database) UpdateStaff(staff *model.Staff) error {
	stmt, err := db.Prepare("UPDATE staff set firstName=$1, lastName=$2 " +
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
