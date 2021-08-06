package service

import (
	"github.com/theNullP0inter/googly/logger"
	"github.com/theNullP0inter/googly/resource"
)

type DbCrudServiceInterface interface {
	CrudInterface
}

type DbCrudService struct {
	*DbService
}

func handleResourceErrors(err error) *ServiceError {
	if err == nil {
		return nil
	}

	if err == resource.ErrResourceNotFound {
		return NewNotFoundServiceError(err)
	}

	if err == resource.ErrInvalidQuery ||
		err == resource.ErrUniqueConstraint ||
		err == resource.ErrInvalidFormat {
		return NewInternalServiceError(err)

	}

	return NewBadRequestError(err)

}

func (s *DbCrudService) Delete(id DataInterface) *ServiceError {
	err := s.DbResourceManagerIntereface.Delete(id)
	return handleResourceErrors(err)
}

func (s *DbCrudService) GetItem(id DataInterface) (DataInterface, *ServiceError) {
	data, err := s.DbResourceManagerIntereface.Get(id)
	return data, handleResourceErrors(err)
}

func (s *DbCrudService) GetList(req DataInterface) (DataInterface, *ServiceError) {
	data, err := s.DbResourceManagerIntereface.List(req)
	return data, handleResourceErrors(err)
}

func (s *DbCrudService) Create(item DataInterface) (DataInterface, *ServiceError) {
	data, err := s.DbResourceManagerIntereface.Create(item)
	return data, handleResourceErrors(err)
}

func (s *DbCrudService) Update(id DataInterface, update DataInterface) *ServiceError {
	err := s.DbResourceManagerIntereface.Update(id, update)
	return handleResourceErrors(err)
}

func NewDbCrudService(logger logger.LoggerInterface, rm resource.DbResourceManagerIntereface) *DbCrudService {
	service := NewDbService(logger, rm)
	return &DbCrudService{service}
}
