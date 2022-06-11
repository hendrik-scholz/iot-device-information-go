package logging

import (
	"sync"

	"go.uber.org/zap"
)

var zapLogger *zap.Logger
var once sync.Once

func GetLogger() *zap.Logger {
	once.Do(func() {
		zapLogger, _ = zap.NewProduction()
	})
	return zapLogger
}
