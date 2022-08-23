package log

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"time"
)

var logger zerolog.Logger

func InitLogger(env, version string) {
	if env == "dev" {
		log.Logger = log.Output(zerolog.ConsoleWriter{
			Out:        os.Stderr,
			TimeFormat: time.RFC3339,
		})
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
	logger = log.With().
		Str("version", version).
		Str("environment", env).
		Logger()
}

func GetLogger() zerolog.Logger {
	return logger.With().Logger()
}
