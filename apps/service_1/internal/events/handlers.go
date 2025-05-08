package events

import (
	"github.com/IBM/sarama"
	"github.com/rs/zerolog/log"
)

type MyTopicHandler struct{}

func (h *MyTopicHandler) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (h *MyTopicHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (h *MyTopicHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message, ok := <-claim.Messages():
			if !ok {
				log.Error().Msg("message channel was closed")
				return nil
			}
			log.Info().
				Bytes("value", message.Value).
				Time("timestamp", message.Timestamp).
				Str("topic", message.Topic).
				Msg("event received")
			session.MarkMessage(message, "")
		case <-session.Context().Done():
			return nil
		}
	}
}
