package logger

import (
	"github.com/sirupsen/logrus"
)

// GooglyLogger is the logger interface implemented in all services for logging
//
// It is just a wrapper around logrus.Logger
type GooglyLogger struct {
	*logrus.Logger
}

func NewGooglyLogger() *GooglyLogger {
	log := logrus.New()
	return &GooglyLogger{log}

}
