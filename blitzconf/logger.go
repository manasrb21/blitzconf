package blitzconf

import (
	"go.uber.org/zap"
)

// Logger instance
var log *zap.Logger

// InitLogger initializes the logger
func InitLogger() {
	var err error
	log, err = zap.NewProduction()
	if err != nil {
		panic("❌ Failed to initialize logger: " + err.Error())
	}
}

// Info logs an info message
func Info(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}

// Error logs an error message
func Error(message string, fields ...zap.Field) {
	log.Error(message, fields...)
}

// Sync flushes any buffered log entries
func Sync() {
	_ = log.Sync()
}
