package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

var DB *sqlx.DB
var DbName string

func Connect(database_url string) {
	if database_url == "" {
		logrus.Fatal("environment variable DATABASE_URL not set")
	}

	logrus.Info("connecting to database")
	db, err := sqlx.Open("postgres", database_url)
	if err != nil {
		logrus.Fatal(err.Error())
	}
	if db == nil {
		logrus.Fatal("did not attach to database")
	}

	if err = db.Ping(); err != nil {
		logrus.Fatal(err)
	}

	DB = db
	GetDatabaseName()
	logrus.Infof("connected to %s", DbName)
}

func GetDatabaseName() {
	err := DB.Get(&DbName, "SELECT current_database();")
	if err != nil {
		logrus.Fatalf("failed to get db name %+v", err)
	}

}
