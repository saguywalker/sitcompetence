package db

import (
	"fmt"
	"strings"

	"github.com/saguywalker/sitcompetence/model"
)

// GetStudentByID returns an student from studentID
func (db *Database) GetStudentByID(id string) (*model.Student, error) {
	row, err := db.Query("SELECT * FROM student WHERE studentID = $1", id)
	if err != nil {
		return nil, err
	}

	var student model.Student

	for row.Next() {
		err := row.Scan(&student)
		if err != nil {
			return nil, nil
		}
	}

	return &student, nil
}

// GetStudents returns all students in a table
func (db *Database) GetStudents(pageLimit uint64, pageNo uint64, dp string, year uint16) ([]*model.Student, error) {
	commands := make([]string, 1)
	commands[0] = "SELECT * FROM student"
	params := make([]string, 0)

	if dp != "" {
		params = append(params, dp)
		commands = append(commands, "WHERE dp=$1")
	}

	if year != 0 {
		if len(params) == 0 {
			commands = append(commands, "WHERE")
		}
		params = append(params, string(year))
		commands = append(commands, fmt.Sprintf("studentId LIKE '%d%%'", len(params)))
	}

	if pageLimit != 0 && pageNo != 0 {
		params = append(params, string(pageLimit))
		params = append(params, string((pageNo-1)*pageLimit))
		paramsLen := len(params)
		commands = append(commands, fmt.Sprintf("ORDER BY studentId LIMIT $%d OFFSET $%d", paramsLen, paramsLen-1))
	}

	command := strings.Join(commands, " ")
	rows, err := db.Query(command, params)
	if err != nil {
		return nil, err
	}

	var students []*model.Student

	for rows.Next() {
		var student model.Student
		err := rows.Scan(&student)
		if err != nil {
			return nil, err
		}
		students = append(students, &student)
	}

	return students, nil
}

// CreateStudent inserts a new student
func (db *Database) CreateStudent(student *model.Student) error {
	stmt, err := db.Prepare("INSERT INTO student(studentId, firstName, lastName, department) VALUES($1, $2, $3, $4)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(student.StudentID, student.FirstName, student.LastName, student.Department)
	if err != nil {
		return err
	}

	return nil
}

// UpdateStudent update student from staff id
func (db *Database) UpdateStudent(student *model.Student) error {
	stmt, err := db.Prepare("UPDATE student set firstName=$1, lastName=$2, department=$3 " +
		"WHERE studentId=$4")

	if err != nil {
		return err
	}

	_, err = stmt.Exec(student.FirstName, student.LastName, student.Department, student.StudentID)
	if err != nil {
		return err
	}

	return nil
}

// DeleteStudent deletes a student from studentID
func (db *Database) DeleteStudent(studentID string) error {
	stmt, err := db.Prepare("DELETE FROM student WHERE studentId = $1")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(studentID)
	if err != nil {
		return err
	}

	return nil
}
