package db

import (
	"database/sql"

	_ "github.com/lib/pq" // postgres driver
	"github.com/pkg/errors"
)

// Database struct
type Database struct {
	*sql.DB
}

// New returns reference to database struct
func New(config *Config) (*Database, error) {
	db, err := sql.Open("postgres", config.DatabaseURI)
	if err != nil {
		return nil, errors.Wrap(err, "unable to connect to database")
	}
	return &Database{db}, nil
}
