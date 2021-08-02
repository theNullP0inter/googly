package resource

import (
	"github.com/sirupsen/logrus"
)

type CrudResourceManagerInterface interface {
	ResourceManagerInterface
	Create(m DataInterface) (DataInterface, error)
	List(parameters DataInterface) (DataInterface, error)
	Get(id DataInterface) (DataInterface, error)

	Update(item DataInterface) (DataInterface, error)

	Delete(id DataInterface) error
}

type CrudResourceManager struct {
	*ResourceManager
	Implementor CrudResourceManagerInterface
}

func NewCrudResourceManager(logger *logrus.Logger,
	resource Resource,
	crud_implementor CrudResourceManagerInterface,
) *CrudResourceManager {
	rm := NewResourceManager(logger, resource)

	return &CrudResourceManager{
		ResourceManager: rm,
		Implementor:     crud_implementor,
	}
}
