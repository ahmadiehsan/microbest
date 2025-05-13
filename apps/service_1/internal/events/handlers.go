package events

import (
	"github.com/IBM/sarama"
	"github.com/agoda-com/opentelemetry-go/otelzap"
	"github.com/dnwe/otelsarama"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
)

type myTopicHandler struct {
	*App
}

func (h *myTopicHandler) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (h *myTopicHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (h *myTopicHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message, ok := <-claim.Messages():
			ctx := otel.GetTextMapPropagator().Extract(session.Context(), otelsarama.NewConsumerMessageCarrier(message))
			if !ok {
				otelzap.Ctx(ctx).Error("message channel was closed")
				return nil
			}
			otelzap.Ctx(ctx).With(
				zap.ByteString("value", message.Value),
				zap.Time("timestamp", message.Timestamp),
				zap.String("topic", message.Topic),
			).Info("event received")
			session.MarkMessage(message, "")
		case <-session.Context().Done():
			return nil
		}
	}
}
