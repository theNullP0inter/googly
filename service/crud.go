package service

import (
	"github.com/theNullP0inter/account-management/logger"
	"github.com/theNullP0inter/account-management/resource"
)

type CrudServiceImplementorInterface interface {
	ServiceInterface
	GetItem(id DataInterface) (DataInterface, error)
	GetList(req DataInterface) (DataInterface, error)
	Create(req DataInterface) (DataInterface, error)
	Update(item DataInterface) (DataInterface, error)
	Delete(id DataInterface) error
}

type CrudServiceInterface interface {
	ServiceInterface
	CrudServiceImplementorInterface
}
type CrudService struct {
	CrudServiceInterface
	*Service
	ResourceManager resource.CrudResourceManagerInterface
}

func NewCrudService(logger logger.LoggerInterface, rm resource.CrudResourceManagerInterface, implementor CrudServiceImplementorInterface) *CrudService {
	service := NewService(logger).(*Service)
	return &CrudService{
		implementor,
		service,
		rm,
	}
}
