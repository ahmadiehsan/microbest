package helpers

import (
	"os"
	"sync"
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
			IsDebug:             os.Getenv("PROJECT_ENV") == "dev",
			KafkaAddress:        os.Getenv("KAFKA_HOST") + ":" + os.Getenv("KAFKA_BROKER_PORT"),
			Service2HttpAddress: os.Getenv("SERVICE_2_HOST") + ":" + os.Getenv("SERVICE_2_HTTP_PORT"),
			Service2GrpcAddress: os.Getenv("SERVICE_2_HOST") + ":" + os.Getenv("SERVICE_2_GRPC_PORT"),
		}
	})
	return instance
}
