package db

// CreateCompetenceReward append competence that can get from activity
func (db *Database) CreateCompetenceReward(activityID uint32, competenceID uint32) error {
	stmt, err := db.Prepare("INSERT INTO competenceReward(activityId, competenceId) VALUES($1, $2);")
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(activityID, competenceID); err != nil {
		return err
	}

	return nil
}

// GetCompetenceReward return all reward badges from specific activityID
func (db *Database) GetCompetenceReward(activityID uint32) ([]uint32, error) {
	rows, err := db.Query("SELECT competenceId FROM competenceReward WHERE activityId=$1", activityID)
	if err != nil {
		return nil, err
	}

	var competences []uint32
	for rows.Next() {
		var competence uint32
		if err := rows.Scan(&competence); err != nil {
			return nil, err
		}

		competences = append(competences, competence)
	}

	return competences, nil
}
