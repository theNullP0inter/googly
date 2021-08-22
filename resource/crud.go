package resource

import (
	"github.com/theNullP0inter/googly/logger"
)

// CrudImplementorInterface is a Generic crud definition interface
type CrudImplementorInterface interface {
	Create(m DataInterface) (DataInterface, error)
	List(parameters DataInterface) (DataInterface, error)
	Get(id DataInterface) (DataInterface, error)
	Update(id DataInterface, item DataInterface) error
	Delete(id DataInterface) error
}

// CrudResourceManagerInterface should be implemented by any resource manager that intends to provide CRUD functionality.
//
// generally a db manager
type CrudResourceManager interface {
	ResourceManager
	CrudImplementorInterface
}

// BaseCrudResourceManager  is a a basic CrudResourceManager.
type BaseCrudResourceManager struct {
	*BaseResourceManager
	CrudImplementorInterface
}

// NewBaseCrudResourceManager creates a new BaseCrudResourceManager with a given crudImplementor
func NewBaseCrudResourceManager(logger logger.GooglyLogger,
	resource Resource,
	crudImplementor CrudImplementorInterface,
) *BaseCrudResourceManager {
	rm := NewBaseResourceManager(logger, resource)

	return &BaseCrudResourceManager{
		rm,
		crudImplementor,
	}
}
