package service

import (
	"github.com/sirupsen/logrus"
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
	ResourceManager resource.ResourceManagerInterface
}

func NewCrudService(logger *logrus.Logger, rm resource.ResourceManagerInterface, implementor CrudServiceImplementorInterface) *CrudService {
	service := NewService(logger)
	return &CrudService{
		implementor,
		service,
		rm,
	}
}
