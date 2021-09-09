package resource

import (
	"github.com/theNullP0inter/googly/logger"
)

// CrudImplementor is a Generic crud definition interface
type CrudImplementor interface {
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
	CrudImplementor
}

// BaseCrudResourceManager  is a a basic CrudResourceManager.
type BaseCrudResourceManager struct {
	*BaseResourceManager
	CrudImplementor
}

// NewBaseCrudResourceManager creates a new BaseCrudResourceManager with a given crudImplementor
func NewBaseCrudResourceManager(logger *logger.GooglyLogger,
	resource Resource,
	crudImplementor CrudImplementor,
) *BaseCrudResourceManager {
	rm := NewBaseResourceManager(logger, resource)

	return &BaseCrudResourceManager{
		rm,
		crudImplementor,
	}
}
