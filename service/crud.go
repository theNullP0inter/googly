package service

// CrudService is any Service that provides CRUD functionality
type CrudService interface {
	Service
	GetItem(id DataInterface) (DataInterface, *ServiceError)
	GetList(req DataInterface) (DataInterface, *ServiceError)
	Create(req DataInterface) (DataInterface, *ServiceError)
	Update(id DataInterface, update DataInterface) *ServiceError
	Delete(id DataInterface) *ServiceError
}
