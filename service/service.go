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
	Logger logger.LoggerInterface
}

func NewService(logger logger.LoggerInterface) *Service {
	return &Service{
		Logger: logger,
	}
}
