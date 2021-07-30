package main

import (
	"github.com/joho/godotenv"
	d "github.com/nstoker-clixifix/find_script/internal/database"
	"github.com/sirupsen/logrus"
)

func main() {
	// Load the (optional) environment variables file
	godotenv.Load(".env")

	d.Connect()
	// Ensures the database connection is closed on exit
	defer d.DB.Close()

	d.ScanTables()

	logrus.Info("exiting")
}
