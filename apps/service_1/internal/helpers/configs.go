package helpers

import (
	"os"

	"github.com/rs/zerolog/log"
)

type Configs struct {
	IsDebug             bool
	KafkaAddress        string
	Service2HttpAddress string
	Service2GrpcAddress string
}

func NewConfigs() *Configs {
	return &Configs{
		IsDebug:             mustGetenv("PROJECT_ENV") == "dev",
		KafkaAddress:        mustGetenv("KAFKA_HOST") + ":" + mustGetenv("KAFKA_BROKER_PORT"),
		Service2HttpAddress: mustGetenv("SERVICE_2_HOST") + ":" + mustGetenv("SERVICE_2_HTTP_PORT"),
		Service2GrpcAddress: mustGetenv("SERVICE_2_HOST") + ":" + mustGetenv("SERVICE_2_GRPC_PORT"),
	}
}

func mustGetenv(envKey string) string {
	val := os.Getenv(envKey)
	if val == "" {
		log.Panic().Msgf("environment variable %q not set", envKey)
	}
	return val
}
