package logger

import (
	"go.uber.org/zap"
)

type ILogger interface {
	Debug(args ...interface{})
	Debugf(template string, args ...interface{})

	Info(args ...interface{})
	Infof(template string, args ...interface{})

	Warn(args ...interface{})
	Warnf(template string, args ...interface{})

	Error(args ...interface{})
	Errorf(template string, args ...interface{})

	DPanic(args ...interface{})
	DPanicf(template string, args ...interface{})

	Panic(args ...interface{})
	Panicf(template string, args ...interface{})

	Fatal(args ...interface{})
	Fatalf(template string, args ...interface{})

	Sync() error
}

type appLogger struct {
	*zap.SugaredLogger
}

func NewLogger() (ILogger, error) {
	zapLogger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}
	sugaredLogger := zapLogger.Sugar()

	return &appLogger{sugaredLogger}, nil
}
