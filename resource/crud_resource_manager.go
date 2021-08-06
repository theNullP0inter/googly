package resource

import (
	"github.com/theNullP0inter/googly/errors"
	"github.com/theNullP0inter/googly/logger"
)

type CrudImplementorInterface interface {
	Create(m DataInterface) (DataInterface, *errors.GooglyError)
	List(parameters DataInterface) (DataInterface, *errors.GooglyError)
	Get(id DataInterface) (DataInterface, *errors.GooglyError)

	Update(id DataInterface, item DataInterface) *errors.GooglyError

	Delete(id DataInterface) *errors.GooglyError
}

type CrudResourceManagerInterface interface {
	ResourceManagerInterface
	CrudImplementorInterface
}

type CrudResourceManager struct {
	*ResourceManager
	CrudImplementorInterface
}

func NewCrudResourceManager(logger logger.LoggerInterface,
	resource Resource,
	crud_implementor CrudImplementorInterface,
) CrudResourceManagerInterface {
	rm := NewResourceManager(logger, resource)

	return &CrudResourceManager{
		rm.(*ResourceManager),
		crud_implementor,
	}
}
