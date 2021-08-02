package service

import (
	"github.com/theNullP0inter/account-management/logger"
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
	Logger logger.LoggerInterface
}

func NewService(logger logger.LoggerInterface) *Service {
	return &Service{
		Logger: logger,
	}
}
