package resource

import "github.com/theNullP0inter/googly/logger"

// BaseResourceManager is  a  base implementation for ResourceManagerInterface
type BaseResourceManager struct {
	Logger   logger.GooglyLoggerInterface
	Resource Resource
}

// GetResource just gets the resource already stored in BaseResourceManager
func (s *BaseResourceManager) GetResource() Resource {
	return s.Resource
}

// NewResourceManager creates a new ResourceManager and checks the implementation with ResourceManagerInterface
func NewBaseResourceManager(logger logger.GooglyLoggerInterface, r Resource) *BaseResourceManager {
	return &BaseResourceManager{logger, r}
}
