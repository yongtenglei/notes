package main

import (
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any

	// 第一个参数为msg, 之后为附加
	logger.Info("from info",
		zap.String("status: ", "ok"),
		zap.Int("code: ", 200),
	)

}
