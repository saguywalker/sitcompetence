package db

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/saguywalker/sitcompetence/model"
)

// GetStudentByID returns an student from studentID
func (db *Database) GetStudentByID(id string) (*model.Student, error) {
	log.Println(fmt.Sprintf("SELECT studentId, firstname, department FROM student WHERE studentID = %s", id))

	row, err := db.Query("SELECT studentId, firstname, department FROM student WHERE studentID = $1", id)
	if err != nil {
		return nil, err
	}

	var student model.Student

	for row.Next() {
		err := row.Scan(&student.StudentID, &student.FirstName, &student.Department)
		if err != nil {
			return nil, nil
		}
	}

	return &student, nil
}

// GetStudents returns all students in a table
func (db *Database) GetStudents(pageLimit uint32, pageNo uint32, dp string, year uint16) ([]*model.Student, error) {
	commands := make([]string, 1)
	commands[0] = "SELECT studentId, firstname, department FROM student "
	params := make([]interface{}, 0)

	if dp != "" {
		params = append(params, dp)
		commands = append(commands, "WHERE LOWER(department)=LOWER($1) ")
	}

	if year != 0 {
		if len(params) == 0 {
			commands = append(commands, "WHERE ")
		} else {
			commands = append(commands, "AND ")
		}
		commands = append(commands, fmt.Sprintf("studentId LIKE '%d%%' ", year))
	}

	if pageLimit != 0 && pageNo != 0 {
		params = append(params, string(pageLimit))
		params = append(params, string((pageNo-1)*pageLimit))
		var paramsLen int = len(params)
		commands = append(commands, fmt.Sprintf("ORDER BY studentId LIMIT $%d OFFSET $%d", paramsLen-1, paramsLen))
	}

	command := strings.Join(commands, " ")

	log.Println(command, ":", params)

	var rows *sql.Rows
	var err error
	if len(params) == 0 {
		rows, err = db.Query(command)
	} else {
		rows, err = db.Query(command, params...)
	}
	if err != nil {
		return nil, err
	}

	var students []*model.Student

	for rows.Next() {
		var student model.Student
		err := rows.Scan(&student.StudentID, &student.FirstName, &student.Department)
		if err != nil {
			return nil, err
		}
		students = append(students, &student)
	}

	return students, nil
}

// CreateStudent inserts a new student
func (db *Database) CreateStudent(s *model.Student) (string, error) {
	var id string
	command := "INSERT INTO student(studentId, firstname, lastname, department) VALUES($1, $2, $3, $4) RETURNING studentId;"
	if err := db.QueryRow(command, s.StudentID, s.FirstName, s.LastName, s.Department).Scan(&id); err != nil {
		return "", err
	}

	return id, nil
}

// UpdateStudent update student from staff id
func (db *Database) UpdateStudent(student *model.Student) error {
	stmt, err := db.Prepare("UPDATE student set firstname=$1, lastname=$2, department=$3 " +
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
