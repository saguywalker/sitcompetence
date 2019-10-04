package db

import (
	"database/sql"

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
		err := row.Scan(&competence.CompetenceID, &competence.CompetenceName, &competence.BadgeIconURL, &competence.TotalActivitiesRequired)
		if err != nil {
			return nil, err
		}
	}

	return &competence, nil
}

// GetCompetencesByActivityID query competence from activity id
func (db *Database) GetCompetencesByActivityID(id uint32, pageLimit, pageNo uint64) ([]model.Competence, error) {
	var competences []model.Competence

	var rows *sql.Rows
	var err error
	if pageLimit == 0 || pageNo == 0 {
		rows, err = db.Query("SELECT c.competenceId, c.competenceName, c.badgeIconUrl, c.totalActivitiesRequired "+
			"FROM competence as c, competenceReward as r "+
			"WHERE r.activityId = $1 AND c.competenceId = r.competenceId", id)
	} else {
		rows, err = db.Query("SELECT c.competenceId, c.competenceName, c.badgeIconUrl, c.totalActivitiesRequired "+
			"FROM competence as c, competenceReward as r "+
			"WHERE r.activityId = $1 AND c.competenceId = r.competenceId "+
			"ORDER BY c.competenceId LIMIT $2 OFFSET $3", id, pageLimit, (pageNo-1)*pageLimit)
	}

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var competence model.Competence
		err := rows.Scan(&competence.CompetenceID, &competence.CompetenceName, &competence.BadgeIconURL, &competence.TotalActivitiesRequired)
		if err != nil {
			return nil, err
		}
		competences = append(competences, competence)
	}

	return competences, nil

}

// GetCompetences returns all competences in a table
func (db *Database) GetCompetences(pageLimit uint64, pageNo uint64) ([]model.Competence, error) {
	var competences []model.Competence

	var rows *sql.Rows
	var err error

	if pageLimit == 0 || pageNo == 0 {
		rows, err = db.Query("SELECT * FROM competence")
	} else {
		rows, err = db.Query("SELECT * FROM competence ORDER BY competenceId LIMIT $1 OFFSET $2", pageLimit, (pageNo-1)*pageLimit)
	}
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var competence model.Competence
		err := rows.Scan(&competence.CompetenceID, &competence.CompetenceName, &competence.BadgeIconURL, &competence.TotalActivitiesRequired)
		if err != nil {
			return nil, err
		}
		competences = append(competences, competence)
	}

	return competences, nil
}

// CreateCompetence inserts a new competence
func (db *Database) CreateCompetence(competence *model.Competence) (int64, error) {
	stmt, err := db.Prepare("INSERT INTO competence(competenceId, competenceName, badgeIconURL, totalActivitiesRequired) VALUES($1, $2, $3, $4)")
	if err != nil {
		return -1, err
	}

	result, err := stmt.Exec(competence)
	if err != nil {
		return -1, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}

	return id, nil
}

// UpdateCompetence update an existing competence from id
func (db *Database) UpdateCompetence(competence *model.Competence) error {
	stmt, err := db.Prepare("UPDATE competence " +
		"set competenceName=$2, badgeIconURL=$3, totalActivitiesRequired=$4 " +
		"WHERE competenceId=$1;")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(competence)
	if err != nil {
		return err
	}

	return nil
}

// DeleteCompetence deletes a competence from competenceID
func (db *Database) DeleteCompetence(competenceID uint16) error {
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
