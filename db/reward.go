package db

// CreateCompetenceReward append competence that can get from activity
func (db *Database) CreateCompetenceReward(activityID uint32, competenceID uint16) error {
	stmt, err := db.Prepare("INSERT INTO competenceReward(activityId, competenceId) VALUES($1, $2)")
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(activityID, competenceID); err != nil {
		return err
	}

	return nil
}
