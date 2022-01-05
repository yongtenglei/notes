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

	// 第一个参数为msg, 之后为附加
	logger.Info("from info",
		zap.String("status", "ok"),
		zap.Int("code", 200),
	)

}
