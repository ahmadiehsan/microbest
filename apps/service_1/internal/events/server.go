package events

import (
	"context"
	"errors"

	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
	"service_1/internal/helpers"
)

type Server struct {
	configs        *helpers.Configs
	readerHandlers []readerHandler
}

type readerHandler struct {
	reader  *kafka.Reader
	handler func(kafka.Message) error
}

func NewServer(cfg *helpers.Configs) (func() error, *Server) {
	var closeFuncs []func() error
	readerHandlers := []readerHandler{}

	myTopicReader := newReader(cfg, "my_topic", "service_1_my_topic_consumer")
	closeFuncs = append(closeFuncs, myTopicReader.Close)
	readerHandlers = append(readerHandlers, readerHandler{myTopicReader, myTopicHandler})

	srv := &Server{
		configs:        cfg,
		readerHandlers: readerHandlers,
	}

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
	errChan := make(chan error, len(s.readerHandlers))
	for _, rh := range s.readerHandlers {
		go s.listenForReader(ctx, rh, errChan)
	}

	return <-errChan
}

func (s *Server) listenForReader(ctx context.Context, rha readerHandler, errChan chan error) {
	log.Info().Msgf("start listening for %s events", rha.reader.Config().Topic)

	for {
		msg, err := rha.reader.FetchMessage(ctx)
		if err != nil {
			errChan <- err
			return
		}

		err = rha.handler(msg)
		if err != nil {
			errChan <- err
			return
		}

		err = rha.reader.CommitMessages(ctx, msg)
		if err != nil {
			errChan <- err
			return
		}
	}
}

func newReader(cfg *helpers.Configs, topic string, groupID string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{cfg.KafkaAddress},
		Topic:       topic,
		GroupID:     groupID,
		StartOffset: kafka.FirstOffset,
	})
}
