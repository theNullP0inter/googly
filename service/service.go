package service

import (
	"github.com/theNullP0inter/googly/logger"
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
	Logger logger.GooglyLoggerInterface
}

func NewService(logger logger.GooglyLoggerInterface) *Service {
	return &Service{
		Logger: logger,
	}
}
