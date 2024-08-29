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

func (l *appLogger) Debug(args ...interface{}) {
	l.SugaredLogger.Debug(append([]interface{}{"[DEBUG]"}, args...)...)
}

func (l *appLogger) Debugf(template string, args ...interface{}) {
	l.SugaredLogger.Debugf("[DEBUG] "+template, args...)
}

func (l *appLogger) Info(args ...interface{}) {
	l.SugaredLogger.Info(append([]interface{}{"[INFO]"}, args...)...)
}

func (l *appLogger) Infof(template string, args ...interface{}) {
	l.SugaredLogger.Infof("[INFO] "+template, args...)
}

func (l *appLogger) Warn(args ...interface{}) {
	l.SugaredLogger.Warn(append([]interface{}{"[WARN]"}, args...)...)
}

func (l *appLogger) Warnf(template string, args ...interface{}) {
	l.SugaredLogger.Warnf("[WARN] "+template, args...)
}

func (l *appLogger) Error(args ...interface{}) {
	l.SugaredLogger.Error(append([]interface{}{"[ERROR]"}, args...)...)
}

func (l *appLogger) Errorf(template string, args ...interface{}) {
	l.SugaredLogger.Errorf("[ERROR] "+template, args...)
}

func (l *appLogger) DPanic(args ...interface{}) {
	l.SugaredLogger.DPanic(append([]interface{}{"[DPANIC]"}, args...)...)
}

func (l *appLogger) DPanicf(template string, args ...interface{}) {
	l.SugaredLogger.DPanicf("[DPANIC] "+template, args...)
}

func (l *appLogger) Panic(args ...interface{}) {
	l.SugaredLogger.Panic(append([]interface{}{"[PANIC]"}, args...)...)
}

func (l *appLogger) Panicf(template string, args ...interface{}) {
	l.SugaredLogger.Panicf("[PANIC] "+template, args...)
}

func (l *appLogger) Fatal(args ...interface{}) {
	l.SugaredLogger.Fatal(append([]interface{}{"[FATAL]"}, args...)...)
}

func (l *appLogger) Fatalf(template string, args ...interface{}) {
	l.SugaredLogger.Fatalf("[FATAL] "+template, args...)
}

func (l *appLogger) Sync() error {
	return l.SugaredLogger.Sync()
}
