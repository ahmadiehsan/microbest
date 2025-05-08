package events

import (
	"context"
	"errors"

	"github.com/IBM/sarama"
	"github.com/rs/zerolog/log"
	"service_1/internal/helpers"
)

type Server struct {
	configs               *helpers.Configs
	consumerGroupHandlers []consumerGroupHandler
}

type consumerGroupHandler struct {
	consumerGroup sarama.ConsumerGroup
	topics        []string
	handler       sarama.ConsumerGroupHandler
}

func NewServer(cfg *helpers.Configs) (func() error, *Server) {
	var closeFuncs []func() error

	srv := &Server{
		configs: cfg,
	}
	srv.setupConsumerGroupHandlers(&closeFuncs)

	shutdown := func() error {
		var errShut error
		for _, fn := range closeFuncs {
			errShut = errors.Join(errShut, fn())
		}
		closeFuncs = nil
		return errShut
	}

	return shutdown, srv
}

func (s *Server) Listen(ctx context.Context) error {
	errChan := make(chan error, 1)
	for _, cgh := range s.consumerGroupHandlers {
		go s.listenForReader(ctx, cgh, errChan)
	}

	return <-errChan
}

func (s *Server) listenForReader(ctx context.Context, cgh consumerGroupHandler, errChan chan error) {
	log.Info().Msgf("start listening for %s events", cgh.topics)

	for {
		err := cgh.consumerGroup.Consume(ctx, cgh.topics, cgh.handler)
		if err != nil {
			if errors.Is(err, sarama.ErrClosedConsumerGroup) {
				return
			}
			errChan <- err
			return
		}
	}
}

func (s *Server) setupConsumerGroupHandlers(closeFuncs *[]func() error) {
	saramaCfg := sarama.NewConfig()
	saramaCfg.Version = sarama.V4_0_0_0
	saramaCfg.Consumer.Offsets.Initial = sarama.OffsetOldest
	brokers := []string{s.configs.KafkaAddress}

	client, err := sarama.NewConsumerGroup(brokers, "service_1_my_topic_consumer", saramaCfg)
	if err != nil {
		log.Panic().Err(err).Msg("error creating consumer group")
	}
	*closeFuncs = append(*closeFuncs, client.Close)
	s.consumerGroupHandlers = append(
		s.consumerGroupHandlers,
		consumerGroupHandler{
			consumerGroup: client,
			topics:        []string{"my_topic"},
			handler:       &MyTopicHandler{},
		},
	)
}
