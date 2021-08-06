package database

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB
var DbName string

func Connect(database_url string) {
	if database_url == "" {
		log.Fatal("environment variable DATABASE_URL not set")
	}

	log.Print("connecting to database")
	db, err := sqlx.Open("postgres", database_url)
	if err != nil {
		log.Fatal(err.Error())
	}
	if db == nil {
		log.Fatal("did not attach to database")
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	DB = db
	GetDatabaseName()
	log.Printf("connected to %s", DbName)
}

func GetDatabaseName() {
	err := DB.Get(&DbName, "SELECT current_database();")
	if err != nil {
		log.Fatalf("failed to get db name %+v", err)
	}

}
