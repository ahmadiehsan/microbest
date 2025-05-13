package loghelper

import (
	"github.com/agoda-com/opentelemetry-go/otelzap"
	"github.com/agoda-com/opentelemetry-logs-go/logs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"service_1/internal/helpers/confighelper"
)

func Setup(cfg *confighelper.Configs, loggerProvider logs.LoggerProvider, processName string) error {
	var logConfig zap.Config

	if cfg.IsDebug {
		logConfig = zap.NewDevelopmentConfig()
		logConfig.EncoderConfig.TimeKey = ""
		logConfig.EncoderConfig.NameKey = ""
		logConfig.EncoderConfig.StacktraceKey = ""
		logConfig.EncoderConfig.ConsoleSeparator = " "
		logConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		logConfig.EncoderConfig.EncodeCaller = callerEncoder
	} else {
		logConfig = zap.NewProductionConfig()
	}

	logger, err := logConfig.Build()
	if err != nil {
		return err
	}

	loggerInstrumented := logger.WithOptions(
		zap.WrapCore(func(core zapcore.Core) zapcore.Core {
			return zapcore.NewTee(
				core,
				otelzap.NewOtelCore(loggerProvider),
			)
		}),
	).With(
		zap.String("process_name", processName),
	)

	zap.ReplaceGlobals(loggerInstrumented)
	return nil
}
