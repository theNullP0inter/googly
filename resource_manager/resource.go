package resource_manager

import "github.com/theNullP0inter/account-management/model"

type ResourceInterface interface {
}

type DataInterface interface {
}

type ValidateDataInterface interface {
	DataInterface
	Validate() map[string][]string
}

type CrudInterface interface {
	Create(m DataInterface) (DataInterface, error)

	List(parameters DataInterface) (DataInterface, error)

	Get(id DataInterface) (DataInterface, error)

	Update(item DataInterface) (DataInterface, error)

	Delete(id DataInterface) error
}

type ModelCrudInterface interface {
	Create(m DataInterface) (DataInterface, error)

	List(parameters DataInterface) (DataInterface, error)

	Get(id model.BinID) (DataInterface, error)

	Update(item DataInterface) (DataInterface, error)

	Delete(id model.BinID) error
}

type CrudResourceInterface interface {
	ResourceInterface
	CrudInterface
	ValidateDataInterface
}
