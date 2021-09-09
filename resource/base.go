package resource

import "github.com/theNullP0inter/googly/logger"

// BaseResourceManager is  a  base implementation for ResourceManager
type BaseResourceManager struct {
	Logger   *logger.GooglyLogger
	Resource Resource
}

// GetResource just gets the resource already stored in BaseResourceManager
func (s *BaseResourceManager) GetResource() Resource {
	return s.Resource
}

// NewResourceManager creates a new ResourceManager
func NewBaseResourceManager(logger *logger.GooglyLogger, r Resource) *BaseResourceManager {
	return &BaseResourceManager{logger, r}
}
