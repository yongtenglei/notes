package main

import (
	"go.uber.org/zap"
)

var logger *zap.Logger

func NewLogger() (*zap.Logger, error) {
	config := zap.NewProductionConfig()
	config.OutputPaths = append(config.OutputPaths, "./mylog.log")
	return config.Build()

}

func main() {
	logger, _ = NewLogger()
	defer logger.Sync() // flushes buffer, if any

	zap.ReplaceGlobals(logger)
	zap.L().Info("global logger")
	zap.S().Info("global suger logger")

}
