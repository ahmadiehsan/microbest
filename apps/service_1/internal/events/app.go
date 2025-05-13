package events

import (
	"context"
	"errors"

	"github.com/IBM/sarama"
	"github.com/agoda-com/opentelemetry-go/otelzap"
	"github.com/dnwe/otelsarama"
	"go.uber.org/zap"
	"service_1/internal/helpers/confighelper"
)

type AppSetupper struct {
	App        *App
	closeFuncs []func() error
}

func NewAppSetupper() *AppSetupper {
	return &AppSetupper{}
}

func (s *AppSetupper) Setup(cfg *confighelper.Configs) error {
	app := &App{
		configs: cfg,
	}

	err := app.setupConsumerGroupHandlers(&s.closeFuncs)
	if err != nil {
		return err
	}

	s.App = app
	return nil
}

func (s *AppSetupper) Shutdown() error {
	var err error
	for _, fn := range s.closeFuncs {
		err = errors.Join(err, fn())
	}
	s.closeFuncs = nil
	return err
}

type App struct {
	configs               *confighelper.Configs
	consumerGroupHandlers []consumerGroupHandler
}

type consumerGroupHandler struct {
	consumerGroup sarama.ConsumerGroup
	topics        []string
	handler       sarama.ConsumerGroupHandler
}

func (a *App) Listen(ctx context.Context) error {
	errChan := make(chan error, 1)
	for _, cgh := range a.consumerGroupHandlers {
		go a.listenForReader(ctx, cgh, errChan)
	}
	return <-errChan
}

func (a *App) listenForReader(ctx context.Context, cgh consumerGroupHandler, errChan chan error) {
	otelzap.Ctx(ctx).With(zap.Strings("topics", cgh.topics)).Info("start listening for events")
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
			handler:       otelsarama.WrapConsumerGroupHandler(&myTopicHandler{a}),
		},
	)
	return nil
}
