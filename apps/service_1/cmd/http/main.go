package main

import (
	"service_1/api/http"
	"service_1/internal/logger"

	"github.com/rs/zerolog/log"
)

func main() {
	startupSetup()
	app := http.NewGinApp()

	err := app.Run()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to run server")
	}
}

func startupSetup() {
	logger.SwitchToHumanReadableMode()
}
