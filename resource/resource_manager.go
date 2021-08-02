package resource

import "github.com/theNullP0inter/account-management/logger"

type ResourceManagerInterface interface {
	GetResource() Resource
}

type ResourceManager struct {
	ResourceManagerInterface
	Resource Resource
	Logger   logger.LoggerInterface
}

func (s *ResourceManager) GetResource() Resource {
	return s.Resource
}
func NewResourceManager(logger logger.LoggerInterface, r Resource) *ResourceManager {
	return &ResourceManager{
		Logger:   logger,
		Resource: r,
	}
}
