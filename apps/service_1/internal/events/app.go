package events

import (
	"context"
	"errors"

	"github.com/IBM/sarama"
	"github.com/rs/zerolog/log"
	"service_1/internal/helpers"
)

type App struct {
	configs               *helpers.Configs
	consumerGroupHandlers []consumerGroupHandler
}

type consumerGroupHandler struct {
	consumerGroup sarama.ConsumerGroup
	topics        []string
	handler       sarama.ConsumerGroupHandler
}

func NewApp(cfg *helpers.Configs) (*App, func() error, error) {
	var closeFuncs []func() error

	app := &App{
		configs: cfg,
	}
	err := app.setupConsumerGroupHandlers(&closeFuncs)
	if err != nil {
		return nil, nil, err
	}

	shutdown := func() error {
		var errShut error
		for _, fn := range closeFuncs {
			errShut = errors.Join(errShut, fn())
		}
		closeFuncs = nil
		return errShut
	}

	return app, shutdown, nil
}

func (a *App) Listen(ctx context.Context) error {
	errChan := make(chan error, 1)
	for _, cgh := range a.consumerGroupHandlers {
		go a.listenForReader(ctx, cgh, errChan)
	}
	return <-errChan
}

func (a *App) listenForReader(ctx context.Context, cgh consumerGroupHandler, errChan chan error) {
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

func (a *App) setupConsumerGroupHandlers(closeFuncs *[]func() error) error {
	saramaCfg := sarama.NewConfig()
	saramaCfg.Version = sarama.V4_0_0_0
	saramaCfg.Consumer.Offsets.Initial = sarama.OffsetOldest
	brokers := []string{a.configs.KafkaAddress}

	client, err := sarama.NewConsumerGroup(brokers, "service_1_my_topic_consumer", saramaCfg)
	if err != nil {
		return err
	}
	*closeFuncs = append(*closeFuncs, client.Close)
	a.consumerGroupHandlers = append(
		a.consumerGroupHandlers,
		consumerGroupHandler{
			consumerGroup: client,
			topics:        []string{"my_topic"},
			handler:       &MyTopicHandler{},
		},
	)
	return nil
}
