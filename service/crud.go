package service

import (
	"github.com/theNullP0inter/googly/errors"
	"github.com/theNullP0inter/googly/logger"
	"github.com/theNullP0inter/googly/resource"
)

type CrudServiceImplementorInterface interface {
	ServiceInterface
	GetItem(id DataInterface) (DataInterface, *errors.GooglyError)
	GetList(req DataInterface) (DataInterface, *errors.GooglyError)
	Create(req DataInterface) (DataInterface, *errors.GooglyError)
	Update(item DataInterface) (DataInterface, *errors.GooglyError)
	Delete(id DataInterface) *errors.GooglyError
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
