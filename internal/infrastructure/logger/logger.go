package logger

import (
	"context"
	"io"
	"log"

	"github.com/AAAAAlexeyyyyy/s3_service/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var global *zap.SugaredLogger

func Init() {
	defaultLogger, err := zap.NewProduction()
	if err != nil {
		log.Printf("Failed to initialize default global modules: %v", err)
	}
	global = defaultLogger.Sugar()
}

func LoadLogger(conf *config.AppConfig, w io.Writer, options ...zap.Option) *zap.SugaredLogger {
	if conf == nil {
		return global
	}
	var logLevel zapcore.Level
	encoderConfig := zap.NewProductionEncoderConfig()
	encoder := zapcore.NewJSONEncoder(encoderConfig)

	if conf.Logger.Environment == "development" {
		encoderConfig = zap.NewDevelopmentEncoderConfig()
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	if err := logLevel.UnmarshalText([]byte(conf.Logger.Level)); err != nil {
		log.Printf("Invalid log level '%s'. Defaulting to info.", conf.Logger.Level)
		logLevel = zapcore.InfoLevel
	}

	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	return zap.New(zapcore.NewCore(encoder, zapcore.AddSync(w), logLevel), options...).Sugar()
}

type ctxLoggerKey struct{}

func WithLogger(ctx context.Context, logger *zap.SugaredLogger) context.Context {
	return context.WithValue(ctx, ctxLoggerKey{}, logger)
}

func FromContext(ctx context.Context) *zap.SugaredLogger {
	l := ctx.Value(ctxLoggerKey{})
	if l != nil {
		return l.(*zap.SugaredLogger)
	}
	return global
}
