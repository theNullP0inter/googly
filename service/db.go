package service

import (
	"github.com/theNullP0inter/googly/logger"
	"github.com/theNullP0inter/googly/resource"
)

// DbService is just a service that can provide apis for a DbResource
type DbService interface {
	Service
	GetDbResourceManager() resource.DbResourceManagerIntereface
}

// BaseDbService is a basic DbService.
//
// You can embed this service and create your expose your own APIs
type BaseDbService struct {
	*BaseService
	ResourceManager resource.DbResourceManagerIntereface
}

// GetDbResourceManager() will return it's ResourceManager
func (s *BaseDbService) GetDbResourceManager() resource.DbResourceManagerIntereface {
	return s.ResourceManager
}

// NewBaseDbService will create a new BaseDbService
func NewBaseDbService(logger logger.GooglyLoggerInterface, rm resource.DbResourceManagerIntereface) *BaseDbService {
	service := NewBaseService(logger)
	return &BaseDbService{service, rm}
}
