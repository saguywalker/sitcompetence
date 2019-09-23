package db

import (
	"github.com/saguywalker/sitcompetence/model"
)

// GetActivityByID returns an activity from activityID
func (db *Database) GetActivityByID(id uint32) (*model.Activity, error) {
	var activity model.Activity

	row, err := db.Query("SELECT * FROM activity WHERE activityId = $1", id)
	if err != nil {
		return nil, err
	}

	for row.Next() {
		err := row.Scan(&activity)
		//err := row.Scan(&activity.ActivityID, &activity.ActivityName, &activity.Date, &activity.Creator, &activity.StudentSite)
		if err != nil {
			return nil, err
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
		err := rows.Scan(&activity)
		//err := rows.Scan(&activity.ActivityID, &activity.ActivityName, &activity.Date, &activity.Creator, &activity.StudentSite)
		if err != nil {
			return nil, err
		}
		activities = append(activities, activity)
	}

	return &activities, nil
}

// GetActivitiesByStaff returns all activities in a table
func (db *Database) GetActivitiesByStaff(id uint32) (*[]model.Activity, error) {
	var activities []model.Activity

	rows, err := db.Query("SELECT * FROM activity WHERE creator = $1", id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var activity model.Activity
		err := rows.Scan(&activity)
		// err := rows.Scan(&activity.ActivityID, &activity.ActivityName, &activity.Date, &activity.Creator)
		if err != nil {
			return nil, err
		}
		activities = append(activities, activity)
	}

	return &activities, nil
}

// GetActivitiesByStudent returns all activities in a table
func (db *Database) GetActivitiesByStudent(id string) ([]model.Activity, error) {
	var activities []model.Activity

	rows, err := db.Query("SELECT activity.activityId, activity.activityName, activity.description, activity.date, activity.creator, activity.studentSite FROM activity, attendedActivity WHERE activity.activityId = attendedActivity.activityId AND activity.activityId = $1", id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var activity model.Activity
		err := rows.Scan(&activity)
		//err := rows.Scan(&activity.ActivityID, &activity.ActivityName, &activity.Date, &activity.Creator, &activity.StudentSite)
		if err != nil {
			return nil, err
		}
		activities = append(activities, activity)
	}

	return activities, nil
}

func (db *Database) GetCompetencesByActivityID(id uint32) ([]model.Competence, error) {
	var competences []model.Competence

	rows, err := db.Query("SELECT c.competenceId, c.competenceName, c.badgeIconUrl, c.totalActivitiesRequired FROM competence as c, competenceReward as r WHERE r.activityId = $1 AND c.competenceId = r.competenceId", id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var competence model.Competence
		err := rows.Scan(&competence)
		if err != nil {
			return nil, err
		}
		competences = append(competences, competence)
	}

	return competences, nil

}

// CreateActivity inserts a new activity
func (db *Database) CreateActivity(activity *model.Activity) error {
	stmt, err := db.Prepare("INSERT INTO activity(activityId, activityName, description, date, creator, studentSite) VALUES($1, $2, $3, $4, $5, $6)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(activity.ActivityID, activity.ActivityName, activity.Description, activity.Date, activity.Creator, activity.StudentSite)
	if err != nil {
		return err
	}

	return nil
}

/*
// UpdateActivity deletes a activity from activityID
func (db *Database) UpdateActivity(activityID uint32) error {
	stmt, err := db.Prepare("UPDATE FROM activity WHERE activityId = $1")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(activityID)
	if err != nil {
		return err
	}

	return nil
}
*/

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
