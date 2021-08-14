package logger

import (
	"github.com/sirupsen/logrus"
)

// GooglyLogrusLogger is a GogglyLogger binding on  logrus.Logger
type GooglyLogrusLogger struct {
	*logrus.Logger
}

// WithData adds extra info to your logs
func (l *GooglyLogrusLogger) WithData(data map[string]interface{}) LoggerInterface {
	return l.Logger.WithFields(data)
}

// NewGooglyLogrusLogger will create a new logrus.Logger with  GogglyLogger binding
func NewGooglyLogrusLogger() *GooglyLogrusLogger {
	return &GooglyLogrusLogger{
		logrus.New(),
	}
}
