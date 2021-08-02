package resource

import (
	"github.com/theNullP0inter/account-management/errors"
	"github.com/theNullP0inter/account-management/logger"
)

type CrudImplementorInterface interface {
	Create(m DataInterface) (DataInterface, *errors.GogetaError)
	List(parameters DataInterface) (DataInterface, *errors.GogetaError)
	Get(id DataInterface) (DataInterface, *errors.GogetaError)

	Update(item DataInterface) (DataInterface, *errors.GogetaError)

	Delete(id DataInterface) *errors.GogetaError
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
