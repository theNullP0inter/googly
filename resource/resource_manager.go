package resource

import "github.com/theNullP0inter/googly/logger"

type ResourceManagerInterface interface {
	GetResource() Resource
}

type ResourceManager struct {
	ResourceManagerInterface
	Resource Resource
	Logger   logger.GooglyLoggerInterface
}

func (s *ResourceManager) GetResource() Resource {
	return s.Resource
}
func NewResourceManager(logger logger.GooglyLoggerInterface, r Resource) ResourceManagerInterface {
	return &ResourceManager{
		Logger:   logger,
		Resource: r,
	}
}
