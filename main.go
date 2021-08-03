package main

import (
	"flag"
	"os"

	"github.com/joho/godotenv"
	d "github.com/nstoker-clixifix/find_script/internal/database"
	"github.com/sirupsen/logrus"
)

var databaseURL string
var useEnvironment bool
var showHelp bool

func init() {
	flag.StringVar(&databaseURL, "d", "", "the heroku style database url for your database")
	flag.BoolVar(&useEnvironment, "e", false, "looks in the environment for 'database_url' (this can be in a .env file ")
	flag.BoolVar(&showHelp, "help", false, "displays the help text")
}

func main() {
	// Load the (optional) environment variables file
	godotenv.Load(".env")
	flag.Parse()

	if useEnvironment {
		databaseURL = os.Getenv("DATABASE_URL")
		if os.Getenv("DATABASE_URL") == "" {
			usage()
			os.Exit(1)
		}
	} else {
		if databaseURL == "" {
			usage()
			os.Exit(1)
		}
	}
	d.Connect(databaseURL)
	// Ensures the database connection is closed on exit
	defer d.DB.Close()

	d.ScanTables()

	logrus.Info("closing connection to %s", d.DbName)
}

func usage() {
	logrus.Println("\nUsage: ")
	flag.PrintDefaults()
}
