package logger

import (
	"log"

	"github.com/nstoker-clixifix/find_script/internal/version"
	"go.uber.org/zap"
)

var Log *zap.SugaredLogger

func StartLogger() *zap.SugaredLogger {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("failed to initialise logger %v", err)
	}

	Log = logger.Sugar()

	Log.Infof("Logger initialised %v", version.Version)

	return Log
}
