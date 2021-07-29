package database

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
	"github.com/nstoker-clixifix/find_script/internal/logger"
)

var DB *sql.DB

func Connect() {
	database_url := os.Getenv("DATABASE_URL")
	if database_url == "" {
		logger.Log.Fatal("environment variable DATABASE_URL not set")
	}

	logger.Log.Info("connecting to database")
	db, err := sql.Open("postgres", database_url)
	if err != nil {
		logger.Log.Fatal(err.Error())
	}
	if db == nil {
		logger.Log.Fatal("did not attach to database")
	}

	if err = db.Ping(); err != nil {
		logger.Log.Fatal(err)
	}

	DB = db
}
