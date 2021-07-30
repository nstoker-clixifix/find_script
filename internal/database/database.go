package database

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

var DB *sql.DB

func Connect() {
	database_url := os.Getenv("DATABASE_URL")
	if database_url == "" {
		logrus.Fatal("environment variable DATABASE_URL not set")
	}

	logrus.Info("connecting to database")
	db, err := sql.Open("postgres", database_url)
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
}
