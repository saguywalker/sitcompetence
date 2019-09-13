package db

import (
	"github.com/saguywalker/sitcompetence/model"
)

// GetActivityByID returns an activity from activityID
func (db *Database) GetActivityByID(id uint32) (*model.Activity, error) {
	var activity model.Activity

	row, err := db.Query("SELECT * FROM activity WHERE activityID = ?", id)
	if err != nil {
		return nil, err
	}

	for row.Next() {
		err := row.Scan(&activity.ActivityID, &activity.ActivityName, &activity.Creator, &activity.Date)
		if err != nil {
			return nil, nil
		}
	}

	return &activity, nil
}

// GetActivities returns all activities in a table
func (db *Database) GetActivities() (*[]model.Activity, error) {
	var activities []model.Activity

	rows, err := db.Query("SELECT * FROM activity")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var activity model.Activity
		err := rows.Scan(&activity.ActivityID, &activity.ActivityName, &activity.Creator, &activity.Date)
		if err != nil {
			return nil, err
		}
		activities = append(activities, activity)
	}

	return &activities, nil
}

// CreateActivity inserts a new activity
func (db *Database) CreateActivity(activity *model.Activity) error {
	stmt, err := db.Prepare("INSERT INTO activity(activityID, activityName, date, creator) VALUES(?, ?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(activity.ActivityID, activity.ActivityName, activity.Date, activity.Creator)
	if err != nil {
		return err
	}

	return nil
}

// DeleteActivity deletes a activity from activityID
func (db *Database) DeleteActivity(activityID []byte) error {
	stmt, err := db.Prepare("DELETE FROM activity WHERE activityID = ?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(activityID)
	if err != nil {
		return err
	}

	return nil
}
