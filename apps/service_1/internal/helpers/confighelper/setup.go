package confighelper

import (
	"github.com/caarlos0/env/v10"
)

type Configs struct {
	IsDebug             bool
	HTTPServerAddress   string
	KafkaAddress        string
	Service2HttpAddress string
	Service2GrpcAddress string
}

func NewConfigs() (*Configs, error) {
	var envVars envVars
	if err := env.Parse(&envVars); err != nil {
		return nil, err
	}

	cfg := &Configs{
		IsDebug:             envVars.ProjectEnv == "dev",
		HTTPServerAddress:   ":" + envVars.Service1Port,
		KafkaAddress:        envVars.KafkaHost + ":" + envVars.KafkaBrokerPort,
		Service2HttpAddress: envVars.Service2Host + ":" + envVars.Service2HttpPort,
		Service2GrpcAddress: envVars.Service2Host + ":" + envVars.Service2GrpcPort,
	}
	return cfg, nil
}

type envVars struct {
	ProjectEnv       string `env:"PROJECT_ENV,required"`
	KafkaHost        string `env:"KAFKA_HOST,required"`
	KafkaBrokerPort  string `env:"KAFKA_BROKER_PORT,required"`
	Service1Port     string `env:"SERVICE_1_PORT,required"`
	Service2Host     string `env:"SERVICE_2_HOST,required"`
	Service2HttpPort string `env:"SERVICE_2_HTTP_PORT,required"`
	Service2GrpcPort string `env:"SERVICE_2_GRPC_PORT,required"`
}
