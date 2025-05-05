package helpers

import "os"

type Configs struct {
	KafkaAddress        string
	Service2HttpAddress string
	Service2GrpcAddress string
}

func LoadConfigs() *Configs {
	return &Configs{
		KafkaAddress:        os.Getenv("KAFKA_HOST") + ":" + os.Getenv("KAFKA_BROKER_PORT"),
		Service2HttpAddress: os.Getenv("SERVICE_2_HOST") + ":" + os.Getenv("SERVICE_2_HTTP_PORT"),
		Service2GrpcAddress: os.Getenv("SERVICE_2_HOST") + ":" + os.Getenv("SERVICE_2_GRPC_PORT"),
	}
}
