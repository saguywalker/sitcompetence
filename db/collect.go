package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/saguywalker/sitcompetence/model"
)

// CreateCollectedCompetence insert new competence to a student
func (db *Database) CreateCollectedCompetence(c *model.CollectedCompetence) error {
	stmt, err := db.Prepare("INSERT INTO collectedCompetence(studentId, competenceId, semester, giver, transactionId) " +
		"VALUES($1, $2, $3, $4, $5)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(c.StudentID, c.CompetenceID, c.Semester, c.Giver, c.TxID)
	if err != nil {
		return err
	}

	return nil
}

// GetCollectedCompetence returns all collected competence in table
func (db *Database) GetCollectedCompetence() ([]*model.CollectedCompetence, error) {
	log.Println("SELECT * FROM collectedCompetence")

	rows, err := db.Query("SELECT * FROM collectedCompetence")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var badges []*model.CollectedCompetence

	for rows.Next() {
		var badge model.CollectedCompetence
		err = rows.Scan(&badge.StudentID, &badge.CompetenceID, &badge.Semester, &badge.Giver, &badge.TxID)
		if err != nil {
			return nil, err
		}
		badges = append(badges, &badge)
	}

	return badges, nil
}

// GetCompetencesByStudentID query competence from student id
func (db *Database) GetCompetencesByStudentID(id string, pageLimit, pageNo uint32) ([]model.CollectedCompetence, error) {
	var collectedList []model.CollectedCompetence

	var rows *sql.Rows
	var err error
	if pageNo == 0 || pageLimit == 0 {
		log.Println(fmt.Sprintf("SELECT studentId, competenceId, semester, giver, transactionId FROM collectedCompetence WHERE studentId=%s;", id))

		rows, err = db.Query("SELECT studentId, competenceId, semester, giver, transactionId FROM collectedCompetence WHERE studentId=$1;", id)
	} else {
		log.Println(fmt.Sprintf("SELECT studentId, competenceId, semester, giver, transactionId FROM collectedCompetence WHERE studentId=%s "+
			"ORDER BY competenceId LIMIT %d OFFSET %d", id, pageLimit, (pageNo-1)*pageLimit))

		rows, err = db.Query("SELECT studentId, competenceId, semester, giver, transactionId FROM collectedCompetence WHERE studentId=$1 "+
			"ORDER BY competenceId LIMIT $2 OFFSET $3", id, pageLimit, (pageNo-1)*pageLimit)
	}
	/*
		rows, err := db.Query("SELECT studentId, competenceId, semester, giver, transactionId FROM collectedCompetence WHERE studentId=$1 "+
			"ORDER BY competenceId LIMIT $2 OFFSET $3", id, uint32(pageLimit), uint32((pageNo-1)*pageLimit))
	*/
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var c model.CollectedCompetence
		if err := rows.Scan(&c.StudentID, &c.CompetenceID, &c.Semester, &c.Giver, &c.TxID); err != nil {
			return nil, err
		}
		collectedList = append(collectedList, c)
	}

	return collectedList, nil
}

// GetCompetencesIDByStudentID query competence id from student id
func (db *Database) GetCompetencesIDByStudentID(id string, pageLimit, pageNo uint32) ([]uint16, error) {
	var competences []uint16

	var rows *sql.Rows
	var err error

	if pageNo == 0 || pageLimit == 0 {
		log.Println(fmt.Sprintf("SELECT competenceId FROM collectedCompetence WHERE studentId=%s;", id))

		rows, err = db.Query("SELECT competenceId FROM collectedCompetence WHERE studentId=$1;", id)
	} else {
		log.Println(fmt.Sprintf("SELECT competenceId FROM collectedCompetence WHERE studentId=%s "+
			"ORDER BY competenceId LIMIT %d OFFSET %d", id, pageLimit, (pageNo-1)*pageLimit))

		rows, err = db.Query("SELECT competenceId FROM collectedCompetence WHERE studentId=$1 "+
			"ORDER BY competenceId LIMIT $2 OFFSET $3", id, pageLimit, (pageNo-1)*pageLimit)
	}

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var competenceID uint16
		if err := rows.Scan(&competenceID); err != nil {
			return nil, err
		}
		competences = append(competences, competenceID)
	}

	return competences, nil
}

// DeleteCollectedByStudentID delete a collected pair from studentId and competenceId
func (db *Database) DeleteCollectedByStudentID(studentID string, competenceID, semester uint16) error {
	stmt, err := db.Prepare("DELETE FROM collectedCompetence WHERE studentId = $1 AND competenceId = $2 AND giver =$3")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(studentID, competenceID, semester)
	if err != nil {
		return err
	}

	return nil
}
