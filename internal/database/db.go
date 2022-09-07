package database

import (
	"fmt"

	"github.com/danilosales/api-street-markets/config"
	"github.com/danilosales/api-street-markets/config/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitializeGorm(conf *config.DbConf, logger *logger.Logger) {

	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", conf.Host, conf.Username, conf.Password, conf.DbName, conf.Port)

	var err error
	DB, err = gorm.Open(postgres.Open(connectionString))

	if err != nil {
		logger.Fatal().Err(err).Msg("Error to connect on database")
	}
}
