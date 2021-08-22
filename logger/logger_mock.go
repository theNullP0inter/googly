package logger

import "github.com/stretchr/testify/mock"

type MockLogger struct {
	mock.Mock
}

func (l *MockLogger) Debugf(format string, args ...interface{}) {}
func (l *MockLogger) Debug(args ...interface{})                 {}

func (l *MockLogger) Infof(format string, args ...interface{}) {}
func (l *MockLogger) Info(args ...interface{})                 {}

func (l *MockLogger) Warnf(format string, args ...interface{}) {}
func (l *MockLogger) Warn(args ...interface{})                 {}

func (l *MockLogger) Errorf(format string, args ...interface{}) {}
func (l *MockLogger) Error(args ...interface{})                 {}

func (l *MockLogger) Fatalf(format string, args ...interface{}) {}
func (l *MockLogger) Fatal(args ...interface{})                 {}

func (l *MockLogger) Panicf(format string, args ...interface{}) {}
func (l *MockLogger) Panic(args ...interface{})                 {}

type MockGooglyLogger struct {
	mock.Mock
	*MockLogger
}

func (l *MockGooglyLogger) WithData(map[string]interface{}) Logger {
	return new(MockLogger)
}
