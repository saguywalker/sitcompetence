package database

import (
	"database/sql"
	"time"
	_ "github.com/lib/pq"
)

type MyDB struct {
	sql.DB
}

func (db *MyDB) createStudentTable() error {
	sqlStmt := `CREATE TABLE student (
		studentID VARCHAR(11) PRIMARY KEY,
		studentFirstName TEXT,
		studentLastName TEXT 
	);`
	_, err := db.Exec(sqlStmt)

	return err
}

func (db *MyDB) createStaffTable() error {
	sqlStmt := `CREATE TABLE staff (
		staffID VARCHAR PRIMARY KEY,
		staffFirstName TEXT,
		staffLastName TEXT,
		publicKey BYTEA
	);`
	_, err := db.Exec(sqlStmt)

	return err
}

func (db *MyDB) createActivityTable() error {
	sqlStmt := `CREATE TABLE student (
		activityID VARCHAR PRIMARY KEY,
		activityName TEXT,
		date DATE,
		creator VARCHAR REFERENCES staff(staffID),
	);`
	_, err := db.Exec(sqlStmt)

	return err
}

func (db *MyDB) createAcquiredCompetenceTable() error {
	sqlStmt := `CREATE TABLE student (
		activityID VARCHAR REFERENCES activity(activityID),
		competenceID VARCHAR REFERENCES competence(competenceID)
	);`
	_, err := db.Exec(sqlStmt)

	return err
}

func (db *MyDB) createAttendedActivityTable() error {
	sqlStmt := `CREATE TABLE student (
		activityID VARCHAR REFERENCES activity(activityID),
		studentID VARCHAR(11) REFERENCES student(studentID),
		approver VARCHAR REFERENCES staff(staffID)
	);`
	_, err := db.Exec(sqlStmt)

	return err
}

func (db *MyDB) createCompetenceTable() error {
	sqlStmt := `CREATE TABLE student (
		competenceID VARCHAR PRIMARY KEY,
		competenceName TEXT,
		description TEXT,
		badgeIconUrl VARCHAR,
		totalActivityRequire INTEGER,
		creator VARCHAR REFERENCES staff(staffID)
	);`
	_, err := db.Exec(sqlStmt)

	return err
}

func (db *MyDB) createCollectedCompetenceTable() error {
	sqlStmt := `CREATE TABLE student (
		studentID VARCHAR(11) REFERENCES student(studentID),
		competenceID VARCHAR REFERENCES competence(competenceID)
	);`
	_, err := db.Exec(sqlStmt)

	return err
}

/* --------------- INSERT DATA --------------- */

func (db *MyDB) InsertStudent(studentID string, studentFirstName string, studentLastName string) error {
	sqlStmt := `INSERT INTO student (studentID, studentFirstName, studentLastName)
		VALUES ($1, $2, $3) 
	);`
	_, err := db.Exec(sqlStmt, studentID, studentFirstName, studentLastName)

	return err
}

func (db *MyDB) InsertStaff(staffID, staffFirstName, staffLastName, publicKey string) error {
	sqlStmt := `INSERT INTO student (studentID, studentFirstName, studentLastName, publicKey)
		VALUES ($1, $2, $3, $4) 
	);`
	_, err := db.Exec(sqlStmt, staffID, staffFirstName, staffLastName, publicKey)

	return err
}

func (db *MyDB) InsertActivity(activityID, activitiyName string, date time.Time, creator string) error {
	sqlStmt := `INSERT INTO student (activityID, activitiyName, date, creator)
		VALUES ($1, $2, $3, $4) 
	);`
	_, err := db.Exec(sqlStmt, activityID, activitiyName, date, creator)

	return err
}

func (db *MyDB) InsertAcquiredCompetence(activityID, competenceID string) error {
	sqlStmt := `INSERT INTO student (activityID, competenceID)
		VALUES ($1, $2) 
	);`
	_, err := db.Exec(sqlStmt, activityID, competenceID)

	return err
}

func (db *MyDB) InsertAttendedActivity(activityID, studentID, approver string) error {
	sqlStmt := `INSERT INTO student (activityID, studentID, approver)
		VALUES ($1, $2, $3) 
	);`
	_, err := db.Exec(sqlStmt, activityID, studentID, approver)

	return err
}

func (db *MyDB) InsertCompetence(competenceID, competenceName, description, badgeIconUrl string, totalActivityRequire int) error {
	sqlStmt := `INSERT INTO student (competenceID, competenceName, description, badgeIconUrl, totalActivityRequire)
		VALUES ($1, $2, $3, $4, $5) 
	);`
	_, err := db.Exec(sqlStmt, competenceID, competenceName, description, badgeIconUrl, totalActivityRequire)

	return err
}

func (db *MyDB) InsertCollectedCompetence(studentID, competenceID string) error {
	sqlStmt := `INSERT INTO student (studentID, competenceID)
		VALUES ($1, $2) 
	);`
	_, err := db.Exec(sqlStmt, studentID, competenceID)

	return err
}

