package db

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

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

	if dp != "" {
		commands = append(commands, fmt.Sprintf("WHERE LOWER(department)=LOWER(%s) ", dp))
	}

	if year != 0 {
		if dp == "" {
			commands = append(commands, "WHERE ")
		} else {
			commands = append(commands, "AND ")
		}
		commands = append(commands, fmt.Sprintf("studentId LIKE '%d%%' ", year))
	}

	if pageLimit != 0 && pageNo != 0 {
		commands = append(commands, fmt.Sprintf("ORDER BY studentId LIMIT %d OFFSET %d", pageLimit, (pageNo-1)*pageLimit))
	}

	command := strings.Join(commands, " ")

	log.Println("command: ", command)

	rows, err := db.Query(command)
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

// UpdateShareProfile update url and expired date
func (db *Database) UpdateShareProfile(studentID string, expire time.Time, url string) error {
	stmt, err := db.Prepare("UPDATE student SET url=$1, expire=$2 WHERE studentId=$3")
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(url, expire, studentID); err != nil {
		return err
	}

	return nil
}

// GetStudentURL return student detail from url
func (db *Database) GetStudentURL(url string) (*model.Student, error) {
	var s model.Student
	row := db.QueryRow("SELECT studentId, firstname, lastname, department FROM student WHERE url=$1 AND CURRENT_TIMESTAMP < expire;", url)
	if err := row.Scan(&s.StudentID, &s.FirstName, &s.LastName, &s.Department); err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return &s, nil
}

// CheckExpire reset url when expired
func (db *Database) CheckExpire(url string) error {
	if _, err := db.Exec("UPDATE student SET url='' WHERE url=$1 AND CURRENT_TIMESTAMP > expire", url); err != nil {
		return err
	}

	return nil
}

// UpdateStudentProfile update profile path and motto
func (db *Database) UpdateStudentProfile(id, filePath, motto string) error {
	var command string
	var err error
	var params []interface{}
	if filePath != "" && motto != "" {
		command = "UPDATE student SET profilePath=$1, motto=$2 WHERE studentId=$3"
		params = []interface{}{filePath, motto, id}
	} else if filePath != "" {
		command = "UPDATE student SET profilePath=$1 WHERE studentId=$2"
		params = []interface{}{filePath, id}
	} else {
		command = "UPDATE student SET motto=$1 WHERE studentId=$2"
		params = []interface{}{motto, id}
	}

	stmt, err := db.Prepare(command)
	if err != nil {
		return err
	}



	if _, err := stmt.Exec(params...); err != nil {
		return err
	}

	return nil
}
