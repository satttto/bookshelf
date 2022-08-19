package logger

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	zap *zap.Logger
}

type EnvType string

const (
	LOCAL EnvType = "local"
	DEV   EnvType = "dev"
	PRD   EnvType = "prd"
)

func New(level string, env EnvType) (*Logger, error) {
	l, err := logLevel(level)
	if err != nil {
		return nil, err
	}

	var config zap.Config
	if env == PRD {
		config = zap.NewProductionConfig()
	} else {
		config = zap.NewDevelopmentConfig()
	}

	config.Level = zap.NewAtomicLevelAt(l)
	config.DisableStacktrace = true
	config.OutputPaths = []string{"stdout"}
	config.ErrorOutputPaths = []string{"stderr"}

	zapLogger, err := config.Build()
	if err != nil {
		return nil, errors.Wrap(err, "failed to build zap config")
	}

	return &Logger{
		zap: zapLogger,
	}, nil
}

// logLevel converts string level into zap log level
func logLevel(level string) (zapcore.Level, error) {
	var l zapcore.Level

	// Log Level: https://github.com/uber-go/zap/blob/1e46f5e6d5d0714ba826e0b7ccdbbad9eab79311/zapcore/level.go#L34-L61
	switch strings.ToUpper(level) {
	case "INFO":
		l = zapcore.InfoLevel
	case "DEBUG":
		l = zapcore.DebugLevel
	case "ERROR":
		l = zapcore.ErrorLevel
	default:
		return l, errors.New(fmt.Sprintf("invalid log level given: %s", level))
	}
	return l, nil
}

// Info outputs an INFO log. Fields: https://pkg.go.dev/go.uber.org/zap#Field
func (l *Logger) Info(msg string, fields ...zapcore.Field) {
	l.zap.Info(msg, fields...)
}

// Debug outputs a debug log. Fields: https://pkg.go.dev/go.uber.org/zap#Field
func (l *Logger) Debug(msg string, fields ...zapcore.Field) {
	l.zap.Debug(msg, fields...)
}

// Error outputs a debug log. Fields: https://pkg.go.dev/go.uber.org/zap#Field
func (l *Logger) Error(msg string, fields ...zapcore.Field) {
	l.zap.Error(msg, fields...)
}
