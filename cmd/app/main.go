package main

import (
	"fmt"

	"github.com/danilosales/api-street-markets/config"
	lr "github.com/danilosales/api-street-markets/config/logger"
	"github.com/danilosales/api-street-markets/internal/api/v1/router"
	"github.com/danilosales/api-street-markets/internal/database"
)

func main() {

	appConf := config.AppConfig()

	logger := lr.New(appConf.LogLevel)

	database.ConnectDatabase(&appConf.Db, logger)

	router := router.New(logger)

	router.Run(fmt.Sprintf(":%d", appConf.Port))

}
