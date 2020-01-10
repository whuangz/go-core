package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
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

	var err error

	if log, err = config.Build(); err != nil {
		panic(err)
	}

	log.Info("log successfully initialized")

	// log.Info("Failed to get Url",
	// 	zap.String("url", "http://www.google.com"),
	// 	zap.Int("attempt", 3),
	// 	zap.Duration("backoff", time.Second),
	// )
}

func Info(msg string, fields ...zap.Field) {
	log.Info(msg, fields...)
}

func Error(msg string, err error, fields ...zap.Field) {
	fields = append(fields, zap.Error(err))

	log.Error(msg, fields...)
}
