package logger

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func SwitchToHumanReadableMode() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}
