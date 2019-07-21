package main

import (
	"database/sql"
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
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

	e.GET("/competence", func(c echo.Context) error {
		return c.JSON(200, "GET competencies")
	})
	e.PUT("/competence", func(c echo.Context) error {
		return c.JSON(200, "PUT competencies")
	})
	e.DELETE("/competence/:id", func(c echo.Context) error {
		return c.JSON(200, "DELETE competence: "+c.Param("id"))
	})

	e.Run(standard.New(":8000"))
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

	sqlStmt := `CREATE TABLE staff (
		staffID VARCHAR PRIMARY KEY,
		staffFirstName TEXT,
		staffLastName TEXT,
		publicKey BYTEA
	);`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		return err
	}

	sqlStmt := `CREATE TABLE activity (
		activityID VARCHAR PRIMARY KEY,
		activityName TEXT,
		date DATE,
		creator VARCHAR REFERENCES staff(staffID),
	);`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		return err
	}

	sqlStmt := `CREATE TABLE acquired_competence (
		activityID VARCHAR REFERENCES activity(activityID),
		competenceID VARCHAR REFERENCES competence(competenceID)
	);`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		return err
	}

	sqlStmt := `CREATE TABLE attended_activity (
		activityID VARCHAR REFERENCES activity(activityID),
		studentID VARCHAR(11) REFERENCES student(studentID),
		approver VARCHAR REFERENCES staff(staffID)
	);`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		return err
	}

	sqlStmt := `CREATE TABLE competence (
		competenceID VARCHAR PRIMARY KEY,
		competenceName TEXT,
		description TEXT,
		badgeIconUrl VARCHAR,
		totalActivityRequire INTEGER,
		creator VARCHAR REFERENCES staff(staffID)
	);`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		return err
	}

	sqlStmt := `CREATE TABLE student (
		studentID VARCHAR(11) REFERENCES student(studentID),
		competenceID VARCHAR REFERENCES competence(competenceID)
	);`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		return err
	}

}
