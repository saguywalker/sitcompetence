package db

import (
	"github.com/saguywalker/sitcompetence/model"
)

// GetCompetenceByID returns a competence from competenceID
func (db *Database) GetCompetenceByID(id uint16) (*model.Competence, error) {
	var competence model.Competence

	row, err := db.Query("SELECT * FROM competence WHERE competenceId = $1")
	if err != nil {
		return nil, err
	}

	for row.Next() {
		err := row.Scan(&competence.CompetenceID, &competence.CompetenceName, &competence.TotalActivitiesRequired, &competence.BadgeIconURL)
		if err != nil {
			return nil, err
		}
	}

	return &competence, nil
}

// GetCompetences returns all competences in a table
func (db *Database) GetCompetences() (*[]model.Competence, error) {
	var competences []model.Competence

	rows, err := db.Query("SELECT * FROM competence")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var competence model.Competence
		err := rows.Scan(&competence.CompetenceID, &competence.CompetenceName, &competence.TotalActivitiesRequired, &competence.BadgeIconURL)
		if err != nil {
			return nil, err
		}
		competences = append(competences, competence)
	}

	return &competences, nil
}

// CreateCompetence inserts a new competence
func (db *Database) CreateCompetence(competence *model.Competence) error {
	stmt, err := db.Prepare("INSERT INTO competence(competenceId, competenceName, totalActivitiesRequired, badgeIconURL) VALUES($1, $2, $3, $4)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(competence.CompetenceID, competence.CompetenceName, competence.TotalActivitiesRequired, competence.BadgeIconURL)
	if err != nil {
		return err
	}

	return nil
}

// DeleteCompetence deletes a competence from competenceID
func (db *Database) DeleteCompetence(competenceID []byte) error {
	stmt, err := db.Prepare("DELETE FROM competence WHERE competenceId = $1")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(competenceID)
	if err != nil {
		return err
	}

	return nil
}
