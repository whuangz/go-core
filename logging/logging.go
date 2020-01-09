package logging

import (
	"fmt"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func main() {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:      "time",
		LevelKey:     "level",
		NameKey:      "logger",
		CallerKey:    "caller",
		MessageKey:   "msg",
		EncodeLevel:  zapcore.LowercaseLevelEncoder,
		EncodeTime:   zapcore.ISO8601TimeEncoder,
		EncodeCaller: zapcore.FullCallerEncoder,
	}

	atomicLevel := zap.NewAtomicLevelAt(zap.DebugLevel)

	config := zap.Config{
		Level:            atomicLevel,
		Development:      true,
		Encoding:         "json",
		EncoderConfig:    encoderConfig,
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	logger, err := config.Build()
	if err != nil {
		panic(fmt.Sprintf("log Failed to initialize: %v", err))
	}
	logger.Info("log successfully initialized")

	logger.Info("Failed to get Url",
		zap.String("url", "http://www.google.com"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}

func Info(msg string, fields ...zap.Field) {
	log.Info(msg, fields...)
}

func Error(msg string, err error, fields ...zap.Field) {
	fields = append(fields, zap.Error(err))

	log.Error(msg, err...)
}
