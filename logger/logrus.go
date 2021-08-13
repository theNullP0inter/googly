package logger

import (
	"github.com/sirupsen/logrus"
)

type GooglyLogrusLogger struct {
	*logrus.Logger
}

func (l *GooglyLogrusLogger) WithData(data map[string]interface{}) LoggerInterface {
	return l.Logger.WithFields(data)
}

func NewGooglyLogrusLogger() *GooglyLogrusLogger {
	return &GooglyLogrusLogger{
		logrus.New(),
	}
}
