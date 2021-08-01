package service

import (
	"github.com/sirupsen/logrus"
)

type DataInterface interface {
}

type ValidateDataInterface interface {
	DataInterface
	Validate() map[string][]string
}

type ServiceInterface interface {
}

type Service struct {
	ServiceInterface
	Logger *logrus.Logger
}

func NewService(logger *logrus.Logger) *Service {
	return &Service{
		Logger: logger,
	}
}
