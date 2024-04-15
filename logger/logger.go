package logger

import (
	"go.uber.org/zap"
)

var logger *zap.SugaredLogger

func GetLogger() *zap.SugaredLogger {
	if logger != nil {
		return logger
	}
	sugar := zap.NewExample().Sugar()
	defer sugar.Sync()
	logger = sugar
	return sugar
}
