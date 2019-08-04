package models

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Student struct {
	StudentID        string `json:"student_id"`
	StudentFirstName string `json:"student_firstname"`
	StudentLastName  string `json:"student_lastname"`
}

type StudentCollection struct {
	Students []Student `json:"students"`
}

func GetStudents(db *sql.DB) StudentCollection {
	sql := "SELECT * FROM student"
	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	result := StudentCollection{}
	for rows.Next() {
		student := Student{}
		err = rows.Scan(
			&student.StudentID,
			&student.StudentFirstName,
			&student.StudentLastName,
		)
		if err != nil {
			panic(err)
		}
		result.Students = append(result.Students, student)
	}

	return result
}

func PostStudent(db *sql.DB, s Student) (int64, error) {
	sql, err := db.Prepare(`INSERT INTO student (studentID, studentFirstName, studentLastName)
		VALUES ($1, $2, $3) 
	)`)
	if err != nil {
		return -1, err
	}

	result, err := sql.Exec(sql, s.StudentID, s.StudentFirstName, s.StudentLastName)
	if err != nil {
		return -1, err
	}

	return result.LastInsertId()
}

func DeleteStudent(db *sql.DB, id string) (int64, error) {
	sql, err := db.Prepare("DELETE FROM student WHERE studentID = $1")
	if err != nil {
		return -1, err
	}

	result, err := sql.Exec(sql, id)
	if err != nil {
		return -1, err
	}

	return result.RowsAffected()
}
