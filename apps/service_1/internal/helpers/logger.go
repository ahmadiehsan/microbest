package helpers

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func SwitchLoggerToHumanReadableMode() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}
