package logger

import (
	"io"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Logger struct {
	logger *zerolog.Logger
}

func NewWithCustomWriter(logLevel string, w io.Writer) *Logger {
	return new(logLevel, w)
}

func New(logLevel string) *Logger {
	return new(logLevel, nil)
}

func new(logLevel string, w io.Writer) *Logger {

	level, err := zerolog.ParseLevel(logLevel)

	if err != nil {
		log.Panic().Msg("the log level is invalid. Valid values: trace, debug, info, warn, error, fatal, panic")
		os.Exit(-1)
	}

	zerolog.SetGlobalLevel(level)

	if w != nil {
		logger := zerolog.New(w).With().Timestamp().Logger()

		return &Logger{logger: &logger}
	}

	file, err := os.OpenFile("street-markets.log", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)

	if err != nil {
		log.Panic().Err(err).Msg("Error to create/open log file")
	}

	writers := io.MultiWriter(zerolog.ConsoleWriter{Out: os.Stderr}, file)

	logger := zerolog.New(writers).With().Timestamp().Logger()

	return &Logger{logger: &logger}

}

// Err starts a new message with error level with err as a field if not nil or
// with info level if err is nil.
//
// You must call Msg on the returned event in order to send the event.
func (l *Logger) Err(err error) *zerolog.Event {
	return l.logger.Err(err)
}

// Trace starts a new message with trace level.
//
// You must call Msg on the returned event in order to send the event.
func (l *Logger) Trace() *zerolog.Event {
	return l.logger.Trace()
}

// Debug starts a new message with debug level.
//
// You must call Msg on the returned event in order to send the event.
func (l *Logger) Debug() *zerolog.Event {
	return l.logger.Debug()
}

// Info starts a new message with info level.
//
// You must call Msg on the returned event in order to send the event.
func (l *Logger) Info() *zerolog.Event {
	return l.logger.Info()
}

// Warn starts a new message with warn level.
//
// You must call Msg on the returned event in order to send the event.
func (l *Logger) Warn() *zerolog.Event {
	return l.logger.Warn()
}

// Error starts a new message with error level.
//
// You must call Msg on the returned event in order to send the event.
func (l *Logger) Error() *zerolog.Event {
	return l.logger.Error()
}

// Fatal starts a new message with fatal level. The os.Exit(1) function
// is called by the Msg method.
//
// You must call Msg on the returned event in order to send the event.
func (l *Logger) Fatal() *zerolog.Event {
	return l.logger.Fatal()
}

// Panic starts a new message with panic level. The message is also sent
// to the panic function.
//
// You must call Msg on the returned event in order to send the event.
func (l *Logger) Panic() *zerolog.Event {
	return l.logger.Panic()
}
