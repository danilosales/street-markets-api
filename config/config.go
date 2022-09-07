package config

import (
	"os"
	"strconv"

	"github.com/rs/zerolog/log"
)

type Conf struct {
	LogLevel string
	Port     int
	Db       dbConf
}

type dbConf struct {
	Host     string
	Port     int
	Username string
	Password string
	DbName   string
}

var defaultValues = map[string]string{
	"LOG_LEVEL": "info",
	"PORT":      "8080",
	"DB_PORT":   "5432",
	"DB_HOST":   "localhost",
	"DB_USER":   "postgres",
	"DB_PASS":   "postgres",
	"DB_NAME":   "markets",
}

func AppConfig() *Conf {
	c := Conf{}

	port, err := strconv.Atoi(getValue("PORT"))

	if err != nil {
		log.Panic().Msg("the environment variable PORT needs to be a valid number")
		os.Exit(-1)
	}

	c.Port = port
	c.Db = *dbConfig()
	c.LogLevel = getValue("LOG_LEVEL")

	return &c
}

func dbConfig() *dbConf {
	c := dbConf{
		Host:     getValue("DB_HOST"),
		Username: getValue("DB_USER"),
		Password: getValue("DB_PASS"),
		DbName:   getValue("DB_NAME"),
	}

	port, err := strconv.Atoi(getValue("DB_PORT"))

	if err != nil {
		log.Panic().Msg("the environment variable DB_PORT needs to be a valid number")
		os.Exit(-1)
	}

	c.Port = port

	return &c
}

func getValue(key string) (value string) {

	valueFromEnv := os.Getenv(key)

	if valueFromEnv == "" {
		if value, found := defaultValues[key]; found {
			return value
		}
		log.Panic().Msgf("the environment variable %s not configured", key)
		os.Exit(-1)
	}

	return valueFromEnv
}
