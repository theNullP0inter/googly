package service

import (
	"github.com/theNullP0inter/account-management/errors"
	"github.com/theNullP0inter/account-management/logger"
	"github.com/theNullP0inter/account-management/resource"
)

type CrudServiceImplementorInterface interface {
	ServiceInterface
	GetItem(id DataInterface) (DataInterface, *errors.GogetaError)
	GetList(req DataInterface) (DataInterface, *errors.GogetaError)
	Create(req DataInterface) (DataInterface, *errors.GogetaError)
	Update(item DataInterface) (DataInterface, *errors.GogetaError)
	Delete(id DataInterface) *errors.GogetaError
}

type CrudServiceInterface interface {
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
