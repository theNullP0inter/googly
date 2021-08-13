package service

type CrudServiceInterface interface {
	ServiceInterface
	GetItem(id DataInterface) (DataInterface, *ServiceError)
	GetList(req DataInterface) (DataInterface, *ServiceError)
	Create(req DataInterface) (DataInterface, *ServiceError)
	Update(id DataInterface, update DataInterface) *ServiceError
	Delete(id DataInterface) *ServiceError
}
