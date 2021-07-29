package main

import (
	d "github.com/nstoker-clixifix/find_script/internal/database"
	"github.com/nstoker-clixifix/find_script/internal/logger"
)

func main() {
	logger := logger.StartLogger()
	defer logger.Sync()

	d.Connect()
}
