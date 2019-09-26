package db

import (
	"github.com/saguywalker/sitcompetence/model"
)

// GetActivityByID returns an activity from activityID
func (db *Database) GetActivityByID(id uint32) (*model.Activity, error) {
	var ac model.Activity

	row, err := db.Query("SELECT * FROM activity WHERE activityId = $1", id)
	if err != nil {
		return nil, err
	}

	for row.Next() {
		err := row.Scan(&ac.ActivityID, &ac.ActivityName, &ac.Description, &ac.Date, &ac.Time, &ac.Creator, &ac.Organizer, &ac.Category, &ac.Location, &ac.StudentSite)
		if err != nil {
			return nil, err
		}
	}

	return &ac, nil
}

// GetActivities returns all activities in a table
func (db *Database) GetActivities() ([]*model.Activity, error) {
	var activities []*model.Activity

	rows, err := db.Query("SELECT * FROM activity")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var ac model.Activity
		err := rows.Scan(&ac.ActivityID, &ac.ActivityName, &ac.Description, &ac.Date, &ac.Time, &ac.Creator, &ac.Organizer, &ac.Category, &ac.Location, &ac.StudentSite)
		if err != nil {
			return nil, err
		}
		activities = append(activities, &ac)
	}

	return activities, nil
}

// GetActivitiesByStaff returns all activities in a table
func (db *Database) GetActivitiesByStaff(id string) ([]*model.Activity, error) {
	var activities []*model.Activity

	rows, err := db.Query("SELECT * FROM activity WHERE creator = $1", id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var ac model.Activity
		err := rows.Scan(&ac.ActivityID, &ac.ActivityName, &ac.Description, &ac.Date, &ac.Time, &ac.Creator, &ac.Organizer, &ac.Category, &ac.Location, &ac.StudentSite)
		if err != nil {
			return nil, err
		}
		activities = append(activities, &ac)
	}

	return activities, nil
}

// GetActivitiesByStudent returns all activities in a table
func (db *Database) GetActivitiesByStudent(id string) ([]*model.Activity, error) {
	var activities []*model.Activity

	rows, err := db.Query("SELECT ac.activityId, ac.activityName, ac.description, ac.date,"+
		"ac.time, ac.creator, ac.organizer, ac.category, ac.location, ac.studentSite"+
		"FROM activity as ac, attendedActivity as at"+
		"WHERE ac.activityId = at.activityId AND ac.activityId = $1", id)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var ac model.Activity
		err := rows.Scan(&ac.ActivityID, &ac.ActivityName, &ac.Description, &ac.Date, &ac.Time, &ac.Creator, &ac.Organizer, &ac.Category, &ac.Location, &ac.StudentSite)
		if err != nil {
			return nil, err
		}
		activities = append(activities, &ac)
	}

	return activities, nil
}

// CreateActivity create an activity from activity struct
func (db *Database) CreateActivity(a *model.Activity) (int64, error) {
	stmt, err := db.Prepare("INSERT INTO activity(activityName, description, date," +
		"time, creator, organizer, category, location, semester, studentSite) " +
		"VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9);")
	if err != nil {
		return -1, err
	}

	result, err := stmt.Exec(a.ActivityName, a.Description, a.Date, a.Time, a.Creator, a.Organizer, a.Category, a.Location, a.Semester, a.StudentSite)
	if err != nil {
		return -1, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

// UpdateActivity update an exisiting activity
func (db *Database) UpdateActivity(a *model.Activity) error {
	stmt, err := db.Prepare("UPDATE activity " +
		"set activityName=$1, description=$2, date=$3, time=$4, " +
		"creator=$5, organizer=$6, category=$7, location=$8, studentSite=$9 " +
		"WHERE activityId=$10")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(a.ActivityName, a.Description, a.Date, a.Time, a.Creator, a.Organizer, a.Category, a.Location, a.StudentSite, a.ActivityID)
	if err != nil {
		return err
	}

	return nil
}

// DeleteActivity deletes a activity from activityID
func (db *Database) DeleteActivity(activityID uint32) error {
	stmt, err := db.Prepare("DELETE FROM activity WHERE activityId = $1")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(activityID)
	if err != nil {
		return err
	}

	return nil
}
