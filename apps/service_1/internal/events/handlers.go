package events

import (
	"strconv"

	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
)

func myTopicHandler(msg kafka.Message) error {
	log.Info().
		Str("partition", strconv.Itoa(msg.Partition)).
		Str("offset", strconv.FormatInt(msg.Offset, 10)).
		Msgf("received event: %s", string(msg.Value))
	return nil
}
