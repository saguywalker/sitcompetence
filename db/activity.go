package db

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/saguywalker/sitcompetence/model"
)

// GetActivityByID returns an activity from activityID
func (db *Database) GetActivityByID(id uint32) (*model.Activity, error) {
	var ac model.Activity

	row, err := db.Query("SELECT * FROM activity WHERE activityId = $1;", id)

	if err != nil {
		return nil, err
	}

	for row.Next() {
		err := row.Scan(&ac.ActivityID, &ac.ActivityName, &ac.Description, &ac.Date, &ac.Time, &ac.Creator, &ac.Organizer, &ac.Category, &ac.Location, &ac.Semester, &ac.StudentSite)
		if err != nil {
			return nil, err
		}
	}

	return &ac, nil
}

// GetActivities returns all activities in a table
func (db *Database) GetActivities(pageNo uint32, pageLimit uint32, isStudent string) ([]*model.Activity, error) {
	var activities []*model.Activity

	commands := make([]string, 1)
	commands[0] = "SELECT * FROM activity "

	if strings.ToLower(isStudent) == "true" {
		commands = append(commands, fmt.Sprintf("WHERE studentSite = %s ", strings.ToLower(isStudent)))
	}

	if pageLimit != 0 && pageNo != 0 {
		commands = append(commands, fmt.Sprintf("ORDER BY activityId LIMIT %d OFFSET %d", pageLimit, (pageNo-1)*pageLimit))
	}

	command := strings.Join(commands, "")

	rows, err := db.Query(command)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var ac model.Activity
		err := rows.Scan(&ac.ActivityID, &ac.ActivityName, &ac.Description, &ac.Date, &ac.Time, &ac.Creator, &ac.Organizer, &ac.Category, &ac.Location, &ac.Semester, &ac.StudentSite)
		if err != nil {
			return nil, err
		}
		activities = append(activities, &ac)
	}

	return activities, nil
}

// GetActivitiesByStaff returns all activities in a table
func (db *Database) GetActivitiesByStaff(id string, pageNo uint32, pageLimit uint32) ([]*model.Activity, error) {
	var activities []*model.Activity

	var rows *sql.Rows
	var err error
	if pageLimit == 0 || pageNo == 0 {
		rows, err = db.Query("SELECT * FROM activity WHERE creator = $1;", id)
	} else {
		rows, err = db.Query("SELECT * FROM activity WHERE creator = $1 "+
			"ORDER BY activityId LIMIT $1 OFFSET $2;", id, pageLimit, (pageNo-1)*pageLimit)
	}

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var ac model.Activity
		err := rows.Scan(&ac.ActivityID, &ac.ActivityName, &ac.Description, &ac.Date, &ac.Time, &ac.Creator, &ac.Organizer, &ac.Category, &ac.Location, &ac.Semester, &ac.StudentSite)
		if err != nil {
			return nil, err
		}
		activities = append(activities, &ac)
	}

	return activities, nil
}

// GetActivitiesByStudent returns all activities in a table
func (db *Database) GetActivitiesByStudent(id string, pageLimit uint32, pageNo uint32) ([]*model.Activity, error) {
	var activities []*model.Activity

	var rows *sql.Rows
	var err error
	if pageLimit == 0 || pageNo == 0 {
		rows, err = db.Query("SELECT ac.activityId, ac.activityName, ac.description, ac.date, "+
			"ac.time, ac.creator, ac.organizer, ac.category, ac.location, ac.semester, ac.studentSite "+
			"FROM activity as ac, attendedActivity as at "+
			"WHERE ac.activityId = at.activityId AND ac.activityId = $1;", id)
	} else {
		rows, err = db.Query("SELECT ac.activityId, ac.activityName, ac.description, ac.date, "+
			"ac.time, ac.creator, ac.organizer, ac.category, ac.location, ac.semester, ac.studentSite "+
			"FROM activity as ac, attendedActivity as at "+
			"WHERE ac.activityId = at.activityId AND ac.activityId = $1 "+
			"ORDER BY ac.activityId LIMIT $1 OFFSET $2;", id, pageLimit, (pageNo-1)*pageLimit)
	}

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
func (db *Database) CreateActivity(a *model.Activity) (uint32, error) {
	var id uint32
	command := "INSERT INTO activity(activityName, description, date," +
		"time, creator, organizer, category, location, semester, studentSite) " +
		"VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING activityId;"

	if err := db.QueryRow(command, a.ActivityName, a.Description, a.Date, a.Time, a.Creator, a.Organizer, a.Category, a.Location, a.Semester, a.StudentSite).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

// UpdateActivity update an exisiting activity
func (db *Database) UpdateActivity(a *model.Activity) error {
	stmt, err := db.Prepare("UPDATE activity " +
		"set activityName=$1, description=$2, date=$3, time=$4, " +
		"creator=$5, organizer=$6, category=$7, location=$8, semester=$9, studentSite=$10 " +
		"WHERE activityId=$11")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(a.ActivityName, a.Description, a.Date, a.Time, a.Creator, a.Organizer, a.Category, a.Location, a.Semester, a.StudentSite, a.ActivityID)
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
