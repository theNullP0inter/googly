package resource

import (
	"github.com/theNullP0inter/googly/logger"
)

type CrudImplementorInterface interface {
	Create(m DataInterface) (DataInterface, error)
	List(parameters DataInterface) (DataInterface, error)
	Get(id DataInterface) (DataInterface, error)

	Update(id DataInterface, item DataInterface) error

	Delete(id DataInterface) error
}

type CrudResourceManagerInterface interface {
	ResourceManagerInterface
	CrudImplementorInterface
}

type CrudResourceManager struct {
	*ResourceManager
	CrudImplementorInterface
}

func NewCrudResourceManager(logger logger.GooglyLoggerInterface,
	resource Resource,
	crud_implementor CrudImplementorInterface,
) CrudResourceManagerInterface {
	rm := NewResourceManager(logger, resource)

	return &CrudResourceManager{
		rm.(*ResourceManager),
		crud_implementor,
	}
}
