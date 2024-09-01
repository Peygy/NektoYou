package mocks

import "github.com/stretchr/testify/mock"

type LoggerMock struct {
	mock.Mock
}

func (m *LoggerMock) Debug(args ...interface{}) {
	m.Called(args...)
}

func (m *LoggerMock) Debugf(template string, args ...interface{}) {
	m.Called(template, args)
}

func (m *LoggerMock) Info(args ...interface{}) {
	m.Called(args...)
}

func (m *LoggerMock) Infof(template string, args ...interface{}) {
	m.Called(template, args)
}

func (m *LoggerMock) Warn(args ...interface{}) {
	m.Called(args...)
}

func (m *LoggerMock) Warnf(template string, args ...interface{}) {
	m.Called(template, args)
}

func (m *LoggerMock) Error(args ...interface{}) {
	m.Called(args...)
}

func (m *LoggerMock) Errorf(template string, args ...interface{}) {
	m.Called(template, args)
}

func (m *LoggerMock) DPanic(args ...interface{}) {
	m.Called(args...)
}

func (m *LoggerMock) DPanicf(template string, args ...interface{}) {
	m.Called(template, args)
}

func (m *LoggerMock) Panic(args ...interface{}) {
	m.Called(args...)
}

func (m *LoggerMock) Panicf(template string, args ...interface{}) {
	m.Called(template, args)
}

func (m *LoggerMock) Fatal(args ...interface{}) {
	m.Called(args...)
}

func (m *LoggerMock) Fatalf(template string, args ...interface{}) {
	m.Called(template, args)
}

func (m *LoggerMock) Sync() error {
	args := m.Called()
	return args.Error(0)
}
