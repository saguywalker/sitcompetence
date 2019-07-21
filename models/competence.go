package models

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Competence struct {
	ID                      string `json:"competence_id"`
	Name                    string `json:"competence_name"`
	Description             string `json:"description"`
	TotalRequiredActivities uint32 `json:"totalactivities"`
	Creator                 string `json:"staff_id"`
}

type CompetenceCollection struct {
	Competencies []Competence `json:"items"`
}

func GetCompetencies(db *sql.DB) CompetenceCollection {
	sql := "SELECT * FROM competence"
	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	result := CompetenceCollection{}
	for rows.Next() {
		competence := Competence{}
		err = rows.Scan(
			&competence.ID,
			&competence.Name,
			&competence.Description,
			&competence.TotalRequiredActivities,
			&competence.Creator,
		)
		if err != nil {
			panic(err)
		}
		result.Competencies = append(result.Competencies, competence)
	}

	return result
}

func PutCompetence(db *sql.DB, c Competence) (int64, error) {
	sql, err := db.Prepare(`INSERT INTO student (competenceID, competenceName, description, badgeIconUrl, totalActivityRequire)
		VALUES ($1, $2, $3, $4, $5) 
	)`)
	if err != nil {
		return -1, err
	}

	result, err := sql.Exec(sql, c.ID, c.Name, c.Description, c.TotalRequiredActivities, c.Creator)
	if err != nil {
		return -1, err
	}

	return result.LastInsertId()
}

func DeleteCompetence(db *sql.DB, id string) (int64, error) {
	sql, err := db.Prepare("DELETE FROM competence WHERE competenceID = $1")
	if err != nil {
		return -1, err
	}

	result, err := sql.Exec(sql, id)
	if er != nil {
		return -1, err
	}

	return result.RowsAffected()
}
