package helpers

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func SetupLogger(cfg *Configs, processName string) {
	if cfg.IsDebug {
		switchToHumanReadableMode()
	}
	log.Logger = log.Logger.With().Str("process_name", processName).Logger()
}

func switchToHumanReadableMode() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}
