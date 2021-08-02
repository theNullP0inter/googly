package service

import (
	"github.com/theNullP0inter/account-management/logger"
	"github.com/theNullP0inter/account-management/resource"
)

type DbCrudServiceInterface interface {
	CrudServiceImplementorInterface
}

type DbCrudService struct {
	*DbService
}

func (s *DbCrudService) Delete(id DataInterface) error {
	return s.DbResourceManagerIntereface.Delete(id)
}

func (s *DbCrudService) GetItem(id DataInterface) (DataInterface, error) {
	return s.DbResourceManagerIntereface.Get(id)
}

func (s *DbCrudService) GetList(req DataInterface) (DataInterface, error) {
	return s.DbResourceManagerIntereface.List(req)
}

func (s *DbCrudService) Create(item DataInterface) (DataInterface, error) {
	return s.DbResourceManagerIntereface.Create(item)
}

func (s *DbCrudService) Update(item DataInterface) (DataInterface, error) {
	return s.DbResourceManagerIntereface.Update(item)
}

func NewDbCrudService(logger logger.LoggerInterface, rm resource.DbResourceManagerIntereface) DbCrudServiceInterface {
	service := NewDbService(logger, rm).(*DbService)
	return &DbCrudService{service}
}
