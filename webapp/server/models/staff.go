package models

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Staff struct {
	StaffID        string `json:"staff_id"`
	StaffFirstName string `json:"staff_firstname"`
	StaffLastName  string `json:"staff_lastname"`
	StaffPubkey    string `json:"staff_pubkey"`
}

type StaffCollection struct {
	Staffs []Staff `json:"staffs"`
}

func GetStaffs(db *sql.DB) StaffCollection {
	sql := "SELECT * FROM staff"
	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	result := StaffCollection{}
	for rows.Next() {
		staff := Staff{}
		err = rows.Scan(
			&staff.StaffID,
			&staff.StaffFirstName,
			&staff.StaffLastName,
			&staff.StaffPubkey,
		)
		if err != nil {
			panic(err)
		}
		result.Staffs = append(result.Staffs, staff)
	}

	return result
}

func PostStaff(db *sql.DB, s Staff) (int64, error) {
	sql, err := db.Prepare(`INSERT INTO student (staffID, staffFirstName, staffLastName, publickey)
		VALUES ($1, $2, $3, $4) 
	)`)
	if err != nil {
		return -1, err
	}

	result, err := sql.Exec(sql, s.StaffID, s.StaffFirstName, s.StaffLastName, s.StaffPubkey)
	if err != nil {
		return -1, err
	}

	return result.LastInsertId()
}

func DeleteStaff(db *sql.DB, id string) (int64, error) {
	sql, err := db.Prepare("DELETE FROM staff WHERE staffID = $1")
	if err != nil {
		return -1, err
	}

	result, err := sql.Exec(sql, id)
	if err != nil {
		return -1, err
	}

	return result.RowsAffected()
}
