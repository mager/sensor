package logger

import (
	"go.uber.org/zap"
)

func ProvideLogger() *zap.SugaredLogger {
	logger, _ := zap.NewProduction()

	return logger.Sugar()
}

var Options = ProvideLogger
