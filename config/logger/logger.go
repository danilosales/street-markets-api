package logger

import (
	"io"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func New(logLevel string) *zerolog.Logger {

	level, err := zerolog.ParseLevel(logLevel)

	if err != nil {
		log.Panic().Msg("the log level is invalid. Valid values: trace, debug, info, warn, error, fatal, panic")
		os.Exit(-1)
	}

	zerolog.SetGlobalLevel(level)

	file, err := os.OpenFile("street-markets.log", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)

	if err != nil {
		log.Panic().Err(err).Msg("Error to create/open log file")
	}

	writers := io.MultiWriter(zerolog.ConsoleWriter{Out: os.Stderr}, file)

	logger := zerolog.New(writers).With().Timestamp().Logger()

	return &logger

}
