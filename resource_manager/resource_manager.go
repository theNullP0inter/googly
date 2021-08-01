package resource_manager

import "github.com/sirupsen/logrus"

type ResourceManagerInterface interface {
	GetResource() ResourceInterface
}

type ResourceManager struct {
	ResourceManagerInterface
	Resource ResourceInterface
	Logger   *logrus.Logger
}

func (s *ResourceManager) GetResource() ResourceInterface {
	return s.Resource
}
func NewResourceManager(logger *logrus.Logger, resource ResourceInterface) *ResourceManager {
	return &ResourceManager{
		Logger:   logger,
		Resource: resource,
	}
}
