package helpers

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func SwitchLoggerToHumanReadableMode() {
	log.Logger = log.Output( //nolint:reassign // Library's suggested way
		zerolog.ConsoleWriter{Out: os.Stderr}, //nolint:exhaustruct // Library's suggested way
	)
}
