package main

import (
	"database/sql"
	"fmt"

	"github.com/labstack/echo"
	"github.com/sit-competence/handlers"
)

const (
	host     = "localhost"
	port     = 5432
	user     = ""
	password = ""
	dbname   = ""
)

func main() {
	fmt.Println("main")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = initDB(db)
	if err != nil {
		panic(err)
	}

	e := echo.New()

	e.GET("/competence", handlers.GetCompetencies(db))
	e.POST("/competence", handlers.PostCompetence(db))
	e.DELETE("/competence/:id", handlers.DeleteCompetence(db))

	e.GET("/staff", handlers.GetCompetencies(db))
	e.POST("/staff", handlers.PostCompetence(db))
	e.DELETE("/staff/:id", handlers.DeleteCompetence(db))

	e.GET("/student", handlers.GetCompetencies(db))
	e.POST("/student", handlers.PostCompetence(db))
	e.DELETE("/student/:id", handlers.DeleteCompetence(db))

	e.GET("/activity", handlers.GetCompetencies(db))
	e.POST("/activity", handlers.PostCompetence(db))
	e.DELETE("/activity/:id", handlers.DeleteCompetence(db))

	e.GET("/competence", handlers.GetCompetencies(db))
	e.POST("/competence", handlers.PostCompetence(db))
	e.DELETE("/competence/:id", handlers.DeleteCompetence(db))

	e.GET("/competence", handlers.GetCompetencies(db))
	e.POST("/competence", handlers.PostCompetence(db))
	e.DELETE("/competence/:id", handlers.DeleteCompetence(db))

	e.GET("/competence", handlers.GetCompetencies(db))
	e.POST("/competence", handlers.PostCompetence(db))
	e.DELETE("/competence/:id", handlers.DeleteCompetence(db))

}

func initDB(db *sql.DB) error {
	sqlStmt := `CREATE TABLE student (
		studentID VARCHAR(11) PRIMARY KEY,
		studentFirstName TEXT,
		studentLastName TEXT 
	);`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		return err
	}

	sqlStmt = `CREATE TABLE staff (
		staffID VARCHAR PRIMARY KEY,
		staffFirstName TEXT,
		staffLastName TEXT,
		publicKey BYTEA
	);`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		return err
	}

	sqlStmt = `CREATE TABLE activity (
		activityID VARCHAR PRIMARY KEY,
		activityName TEXT,
		date DATE,
		creator VARCHAR REFERENCES staff(staffID),
	);`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		return err
	}

	sqlStmt = `CREATE TABLE acquired_competence (
		activityID VARCHAR REFERENCES activity(activityID),
		competenceID VARCHAR REFERENCES competence(competenceID)
	);`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		return err
	}

	sqlStmt = `CREATE TABLE attended_activity (
		activityID VARCHAR REFERENCES activity(activityID),
		studentID VARCHAR(11) REFERENCES student(studentID),
		approver VARCHAR REFERENCES staff(staffID)
	);`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		return err
	}

	sqlStmt = `CREATE TABLE competence (
		competenceID VARCHAR PRIMARY KEY,
		competenceName TEXT,
		description TEXT,
		badgeIconUrl VARCHAR,
		totalActivityRequire INTEGER,
		creator VARCHAR REFERENCES staff(staffID)
	);`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		return err
	}

	sqlStmt = `CREATE TABLE student (
		studentID VARCHAR(11) REFERENCES student(studentID),
		competenceID VARCHAR REFERENCES competence(competenceID)
	);`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		return err
	}

}
