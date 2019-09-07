package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // postgres driver
	"github.com/pkg/errors"
)

// Database struct
type Database struct {
	*gorm.DB
}

// New returns reference to database struct
func New(config *Config) (*Database, error) {
	db, err := gorm.Open("postgres", config.DatabaseURI)
	if err != nil {
		return nil, errors.Wrap(err, "unable to connect to database")
	}
	return &Database{db}, nil
}
