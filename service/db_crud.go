package service

import (
	"github.com/theNullP0inter/googly/errors"
	"github.com/theNullP0inter/googly/logger"
	"github.com/theNullP0inter/googly/resource"
)

type DbCrudServiceInterface interface {
	CrudServiceImplementorInterface
}

type DbCrudService struct {
	*DbService
}

func (s *DbCrudService) Delete(id DataInterface) *errors.GooglyError {
	return s.DbResourceManagerIntereface.Delete(id)
}

func (s *DbCrudService) GetItem(id DataInterface) (DataInterface, *errors.GooglyError) {
	return s.DbResourceManagerIntereface.Get(id)
}

func (s *DbCrudService) GetList(req DataInterface) (DataInterface, *errors.GooglyError) {
	return s.DbResourceManagerIntereface.List(req)
}

func (s *DbCrudService) Create(item DataInterface) (DataInterface, *errors.GooglyError) {
	return s.DbResourceManagerIntereface.Create(item)
}

func (s *DbCrudService) Update(item DataInterface) (DataInterface, *errors.GooglyError) {
	return s.DbResourceManagerIntereface.Update(item)
}

func NewDbCrudService(logger logger.LoggerInterface, rm resource.DbResourceManagerIntereface) DbCrudServiceInterface {
	service := NewDbService(logger, rm).(*DbService)
	return &DbCrudService{service}
}
