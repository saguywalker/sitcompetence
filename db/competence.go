package db

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"

	"github.com/saguywalker/sitcompetence/model"
)

func (db *Database) GetCompetenceByID(id uint16) (*model.Competence, error) {
	var competence model.Competence

	if err := db.First(&competence, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}

		return nil, errors.Wrap(err, "unable to get todo")
	}

	return &competence, nil
}

func (db *Database) GetCompetences() ([]*model.Competence, error) {
	var competences []*model.Competence
	err := errors.Wrap(db.Find(&competences).Error, "unable to get competences")
	return competences, err
}

func (db *Database) CreateCompetence(competence *model.Competence) error {
	return errors.Wrap(db.Create(competence).Error, "unable to create competence")
}
