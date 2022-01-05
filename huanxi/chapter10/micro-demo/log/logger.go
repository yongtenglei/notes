package log

import (
	"go.uber.org/zap"
)

var Logger *zap.Logger
var err error

func init() {
	Logger, err = NewLogger()
	if err != nil {
		panic(err)
	}
}

func NewLogger() (*zap.Logger, error) {
	prolog := zap.NewProductionConfig()
	prolog.OutputPaths = append(prolog.OutputPaths, "./accountHandler.log")
	return prolog.Build()

}
