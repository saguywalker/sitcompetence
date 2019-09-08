package db

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"

	"github.com/saguywalker/sitcompetence/model"
)

func (db *Database) GetActivityByID(id uint32) (*model.Activity, error) {
	var activity model.Activity

	if err := db.First(&activity, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}

		return nil, errors.Wrap(err, "unable to get todo")
	}

	return &activity, nil
}

func (db *Database) GetActivities() ([]*model.Activity, error) {
	var activities []*model.Activity
	err := errors.Wrap(db.Find(&activities).Error, "unable to get activities")
	return activities, err
}

func (db *Database) CreateActivity(activity *model.Activity) error {
	return errors.Wrap(db.Create(activity).Error, "unable to create an activity")
}
