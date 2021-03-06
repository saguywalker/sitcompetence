package db

import (
	"github.com/saguywalker/sitcompetence/model"
)

// GetAttendeesByActivityID return student id who joined a corresponding activity id
func (db *Database) GetAttendeesByActivityID(id uint32) ([]model.Student, error) {
	var students []model.Student

	rows, err := db.Query("SELECT s.studentId, s.firstname, s.department "+
		"FROM attendedActivity as a, student as s "+
		"WHERE a.activityId=$1 AND s.studentId=a.studentId;", id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var student model.Student
		if err := rows.Scan(&student.StudentID, &student.FirstName, &student.Department); err != nil {
			return nil, err
		}

		students = append(students, student)
	}

	return students, nil
}

// AddAttendee add student to attended activity
func (db *Database) AddAttendee(activityID uint32, studentID string) error {
	stmt, err := db.Prepare("INSERT INTO attendedActivity(activityId, studentId) " +
		"VALUES($1, $2);")
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(activityID, studentID); err != nil {
		return err
	}

	return nil
}

// ApproveAttended add approved student to attended activity
func (db *Database) ApproveAttended(a *model.AttendedActivity) error {
	stmt, err := db.Prepare("UPDATE attendedActivity SET transactionId=$1, approver=$2 " +
		"WHERE activityId=$3 AND studentId=$4")
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(a.TransactionID, a.Approver, a.ActivityID, a.StudentID); err != nil {
		return err
	}

	return nil
}

// DeleteAttendedActivity delete all student from a corresponding activity id
func (db *Database) DeleteAttendedActivity(activityID uint32) error {
	stmt, err := db.Prepare("DELETE FROM attendedActivity WHERE activityId=$1")
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(activityID); err != nil {
		return err
	}

	return nil
}

// DeleteAttendedActivityByStudentID delete all student from a corresponding student id
func (db *Database) DeleteAttendedActivityByStudentID(activityID uint32, studentID string) error {
	stmt, err := db.Prepare("DELETE FROM attendedActivity WHERE activityId=$1 AND studentId=$2")
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(activityID, studentID); err != nil {
		return err
	}

	return nil
}
