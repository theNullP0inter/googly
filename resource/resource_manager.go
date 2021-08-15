package resource

import "github.com/theNullP0inter/googly/logger"

// ResourceManagerInterface is a base interface for a resource manager
//
// ResourceManager acts like a proxy between the service and ORM( or mongo-driver).
// i.e, if you want to migrate your service from mongo to rdb, it'll expose similar APIs
type ResourceManagerInterface interface {
	GetResource() Resource
}

// ResourceManager is  a  base implementation for ResourceManagerInterface
type ResourceManager struct {
	Logger   logger.GooglyLoggerInterface
	Resource Resource
}

// GetResource just gets the resource already stored in ResourceManager
func (s *ResourceManager) GetResource() Resource {
	return s.Resource
}

// NewResourceManager creates a new ResourceManager and checks the implementation with ResourceManagerInterface
func NewResourceManager(logger logger.GooglyLoggerInterface, r Resource) *ResourceManager {
	return &ResourceManager{logger, r}
}
