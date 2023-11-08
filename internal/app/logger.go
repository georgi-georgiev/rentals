package app

import (
	"go.uber.org/zap"
)

func NewLogger(conf *Config) *zap.Logger {
	logger, _ := zap.NewDevelopment()

	return logger
}
