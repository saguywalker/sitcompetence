package db

import (
	"github.com/Sirupsen/logrus"
	"github.com/gocraft/dbr"
	_ "github.com/lib/pq" // Postgres driver
	"github.com/saguywalker/sit-competence/webapp/conf"
)

// Init database session
func Init() *dbr.Session {
	session := getSession()
	return session
}

func getSession() *dbr.Session {
	db, err := dbr.Open("postgres", conf.USER+":"+conf.PASSWORD+"@tcp("+conf.HOST+":"+conf.PORT+")/"+conf.DB,
		nil)

	if err != nil {
		logrus.Error(err)
	} else {
		session := db.NewSession(nil)
	}

	return nil
}
