package helpers

import (
	"os"
	"sync"

	"github.com/rs/zerolog/log"
)

type Configs struct {
	IsDebug             bool
	KafkaAddress        string
	Service2HttpAddress string
	Service2GrpcAddress string
}

var (
	instance *Configs
	once     sync.Once
)

func GetConfigs() *Configs {
	once.Do(func() {
		instance = &Configs{
			IsDebug:             mustGetenv("PROJECT_ENV") == "dev",
			KafkaAddress:        mustGetenv("KAFKA_HOST") + ":" + mustGetenv("KAFKA_BROKER_PORT"),
			Service2HttpAddress: mustGetenv("SERVICE_2_HOST") + ":" + mustGetenv("SERVICE_2_HTTP_PORT"),
			Service2GrpcAddress: mustGetenv("SERVICE_2_HOST") + ":" + mustGetenv("SERVICE_2_GRPC_PORT"),
		}
	})
	return instance
}

func mustGetenv(envKey string) string {
	val := os.Getenv(envKey)
	if val == "" {
		log.Fatal().Msgf("environment variable %q not set", envKey)
	}
	return val
}
