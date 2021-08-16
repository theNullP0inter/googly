package service

import (
	"github.com/theNullP0inter/googly/logger"
)

// DataInterface is a basic interafce used for data passed around in the service
type DataInterface interface {
}

// Service is an empty interface.
//
// Service can implement anything ( also nothing )
type Service interface {
}

// BaseServcie is a Service with Just a Logger
type BaseService struct {
	Logger logger.GooglyLoggerInterface
}

// NewBaseService create a new BaseService
func NewBaseService(logger logger.GooglyLoggerInterface) *BaseService {
	return &BaseService{
		Logger: logger,
	}
}
