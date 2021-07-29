package main

import (
	"github.com/joho/godotenv"
	d "github.com/nstoker-clixifix/find_script/internal/database"
	"github.com/nstoker-clixifix/find_script/internal/logger"
)

func main() {
	// Load the (optional) environment variables file
	godotenv.Load(".env")

	logger := logger.StartLogger()
	defer logger.Sync()

	d.Connect()
	// Ensures the database connection is closed on exit
	defer d.DB.Close()

	if err := d.ScanTables(); err != nil {
		logger.Fatal(err)
	}

	logger.Info("exiting")
}
