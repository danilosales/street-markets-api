package main

import (
	"github.com/danilosales/api-street-markets/config"
	lr "github.com/danilosales/api-street-markets/config/logger"
)

func main() {

	appConf := config.AppConfig()

	logger := lr.New(appConf.LogLevel)

	logger.Info().Msg("Teste de log")
}
