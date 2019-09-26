package db

import (
	"github.com/saguywalker/sitcompetence/model"
)

// CreateCollectedCompetence insert new competence to a student
func (db *Database) CreateCollectedCompetence(badges *model.CollectedCompetence) error {
	stmt, err := db.Prepare("INSERT INTO collectedCompetence(studentId, competenceId, transactionId) VALUES($1, $2, $3)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(badges.StudentID, badges.CompetenceID, badges.TxID)
	if err != nil {
		return err
	}

	return nil
}

// GetCollectedCompetence returns all collected competence in table
func (db *Database) GetCollectedCompetence() ([]*model.CollectedCompetence, error) {
	rows, err := db.Query("SELECT * FROM collectedCompetence")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var badges []*model.CollectedCompetence

	for rows.Next() {
		var badge model.CollectedCompetence
		err = rows.Scan(&badge.StudentID, &badge.CompetenceID, &badge.TxID)
		if err != nil {
			return nil, err
		}
		badges = append(badges, &badge)
	}

	return badges, nil
}

// GetMerkleByStudentID returns all query collected competence from studentId
func (db *Database) GetMerkleByStudentID(studentId string) ([]*model.CollectedCompetence, error) {
	row, err := db.Query("SELECT * FROM collectedCompetence WHERE studentId = $1", studentId)
	if err != nil {
		return nil, err
	}

	var badges []*model.CollectedCompetence

	for row.Next() {
		var badge model.CollectedCompetence
		err := row.Scan(&badge.StudentID, &badge.CompetenceID, &badge.TxID)
		if err != nil {
			return nil, err
		}

		badges = append(badges, &badge)
	}

	return badges, nil
}

// GetMerkleByCompetenceID returns all query collected competence from competenceId
func (db *Database) GetMerkleByCompetenceID(competenceId uint16) ([]*model.CollectedCompetence, error) {
	row, err := db.Query("SELECT * FROM collectedCompetence WHERE competenceId = $1", competenceId)
	if err != nil {
		return nil, err
	}

	var badges []*model.CollectedCompetence

	for row.Next() {
		var badge model.CollectedCompetence
		err := row.Scan(&badge.StudentID, &badge.CompetenceID, &badge.TxID)
		if err != nil {
			return nil, err
		}

		badges = append(badges, &badge)
	}

	return badges, nil
}

// GetCompetencesByStudentID query competence from student id
func (db *Database) GetCompetencesByStudentID(id string) ([]model.Competence, error) {
	var competences []model.Competence

	rows, err := db.Query("SELECT c.competenceId, c.competenceName, c.badgeIconUrl, c.totalActivitiesRequired " +
		"FROM competence as c, collectedCompetence as s " +
		"WHERE c.competenceId = s.competenceId AND s.studentId=$1;")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var competence model.Competence
		if err := rows.Scan(&competence); err != nil {
			return nil, err
		}
		competences = append(competences, competence)
	}

	return competences, nil
}

// DeleteCollectedByStudentID delete a collected pair from studentId and competenceId
func (db *Database) DeleteCollectedByStudentID(studentID string, competenceID uint16) error {
	stmt, err := db.Prepare("DELETE FROM collectedCompetence WHERE studentId = $1 AND competenceId = $2")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(studentID, competenceID)
	if err != nil {
		return err
	}

	return nil
}
