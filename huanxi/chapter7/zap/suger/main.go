package main

import (
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any

	sugar := logger.Sugar()

	// 第一个参数为msg, 之后为附加
	sugar.Infow("from infow",
		// Structured context as loosely typed key-value pairs.
		"stutus", "ok",
	)

	// 直接写入msg
	sugar.Infof("status %s", "ok")
}
