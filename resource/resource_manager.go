package resource

import (
	"github.com/sirupsen/logrus"
)

type ResourceManagerInterface interface {
	GetResource() Resource
}

type ResourceManager struct {
	ResourceManagerInterface
	Resource Resource
	Logger   *logrus.Logger
}

func (s *ResourceManager) GetResource() Resource {
	return s.Resource
}
func NewResourceManager(logger *logrus.Logger, r Resource) *ResourceManager {
	return &ResourceManager{
		Logger:   logger,
		Resource: r,
	}
}
