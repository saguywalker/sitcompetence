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
func (db *Database) GetActivities(pageNo uint64, pageLimit uint64, isStudent string) ([]*model.Activity, error) {
	var activities []*model.Activity

	commands := make([]string, 1)
	commands[0] = "SELECT * FROM activity "
	params := make([]interface{}, 0)

	if strings.ToLower(isStudent) == "true" {
		params = append(params, isStudent)
		commands = append(commands, "WHERE studentSite = $1 ")
	}

	if pageLimit != 0 && pageNo != 0 {
		params = append(params, string(pageLimit))
		params = append(params, string((pageNo-1)*pageLimit))
		paramsLen := len(params)
		commands = append(commands, fmt.Sprintf("ORDER BY activityId LIMIT $%d OFFSET $%d", paramsLen, paramsLen-1))
	}

	command := strings.Join(commands, " ")

	var rows *sql.Rows
	var err error

	if len(params) == 0 {
		rows, err = db.Query(command)
	} else {
		rows, err = db.Query(command, params...)
	}

	if err != nil {
		return nil, err
	}

	/*
		var rows *sql.Rows
		var err error
		if pageLimit == 0 || pageNo == 0 {
			rows, err = db.Query("SELECT * FROM activity;")
		} else {
			rows, err = db.Query("SELECT * FROM activity ORDER BY activityId LIMIT $1 OFFSET pageNo $2;", pageLimit, (pageNo-1)*pageLimit)
		}
		if err != nil {
			return nil, err
		}
	*/
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
func (db *Database) GetActivitiesByStaff(id string, pageNo uint64, pageLimit uint64) ([]*model.Activity, error) {
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
func (db *Database) GetActivitiesByStudent(id string, pageLimit uint64, pageNo uint64) ([]*model.Activity, error) {
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
func (db *Database) CreateActivity(a *model.Activity) (int64, error) {
	stmt, err := db.Prepare("INSERT INTO activity(activityName, description, date," +
		"time, creator, organizer, category, location, semester, studentSite) " +
		"VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);")
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
