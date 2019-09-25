package db

// GetAttendeesByActivityID return student id who joined a corresponding activity id
func (db *Database) GetAttendeesByActivityID(id uint32) ([]string, error) {
	var studentsID []string

	rows, err := db.Query("SELECT studentId FROM attendedActivity WHERE activityId=$1", id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var studentID string
		if err := rows.Scan(&studentID); err != nil {
			return nil, err
		}

		studentsID = append(studentsID, studentID)
	}

	return studentsID, nil
}

// AddAttendee add student to attended activity
func (db *Database) AddAttendee(activityID uint32, studentID string) error {
	stmt, err := db.Prepare("INSER INTO attendedActivity(activityId, studentId)" +
		"VALUES($1, $2)")
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(activityID, studentID) {
		return err
	}

	return nil
}

// DeleteAttendedActivity delete all student from a corresponding activity id
func (db *Database) DeleteAttendedActivity(activityId uint32) error {
	stmt, err := db.Prepare("DELETE FROM attendedActivity WHERE activityId=$1")
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(activityId) {
		return err
	}

	return nil
}
