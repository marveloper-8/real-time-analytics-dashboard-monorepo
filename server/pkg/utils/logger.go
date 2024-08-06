package utils

import (
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger	*zap.Logger
	once		sync.Once
)

func InitLogger() *zap.Logger {
	once.Do(func() {
		var config zap.Config

		if os.Getenv("ENV") == "production" {
			config = zap.NewProductionConfig()
		} else {
			config = zap.NewDevelopmentConfig()
		}

		config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

		var err error
		logger, err = config.Build()
		if err != nil {
			panic("Failed to initialize logger: " + err.Error())
		}
	})

	return logger
}

func GetLogger() *zap.Logger {
	if logger == nil {
		return InitLogger()
	}
	return logger
}

func Info(msg string, fields ...zap.Field) {
	GetLogger().Info(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	GetLogger().Error(msg, fields...)
}

func Debug(msg string, fields ...zap.Field)  {
	GetLogger().Debug(msg, fields...)
}

func Warn(msg string, fields ...zap.Field)  {
	GetLogger().Warn(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field)  {
	GetLogger().Fatal(msg, fields...)
}

func With(fields ...zap.Field) *zap.Logger  {
	return GetLogger().With(fields...)
}

func Sync() error  {
	return GetLogger().Sync()
}