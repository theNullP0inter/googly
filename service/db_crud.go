package service

import (
	"github.com/theNullP0inter/googly/logger"
	"github.com/theNullP0inter/googly/resource"
)

// CrudDbService is any service that implements CRUD on a DbResource
//
// So, it has to be a CrudService and a DbService
type CrudDbService interface {
	CrudService
	DbService
}

// BaseCrudDbService is a basic CrudDbService implementation
//
// It uses DbResourceManagerIntereface implements CRUD
type BaseCrudDbService struct {
	*BaseDbService
}

// handleResourceErrors converts all errors from resource manager to ServiceError
func handleResourceErrors(err error) *ServiceError {
	if err == nil {
		return nil
	}

	if err == resource.ErrResourceNotFound {
		return NewNotFoundServiceError(err)
	}

	if err == resource.ErrInvalidQuery ||
		err == resource.ErrUniqueConstraint ||
		err == resource.ErrInvalidFormat ||
		err == resource.ErrInternal {
		return NewInternalServiceError(err)

	}

	return NewBadRequestError(err)

}

// Delete will delete the item with given id using the resource_manager
func (s *BaseCrudDbService) Delete(id DataInterface) *ServiceError {
	err := s.ResourceManager.Delete(id)
	return handleResourceErrors(err)
}

// GetItem will get item with given id using the resource_manager
func (s *BaseCrudDbService) GetItem(id DataInterface) (DataInterface, *ServiceError) {
	data, err := s.ResourceManager.Get(id)
	return data, handleResourceErrors(err)
}

// GetList will get item list using the resource manager
//
// Takes in a request which will be converted into QueryParameters at ResourceManager
func (s *BaseCrudDbService) GetList(req DataInterface) (DataInterface, *ServiceError) {
	data, err := s.ResourceManager.List(req)
	return data, handleResourceErrors(err)
}

// Create will create an list using the resource manager
func (s *BaseCrudDbService) Create(item DataInterface) (DataInterface, *ServiceError) {
	data, err := s.ResourceManager.Create(item)
	return data, handleResourceErrors(err)
}

// Update will update the item with given id using the resource manager
func (s *BaseCrudDbService) Update(id DataInterface, update DataInterface) *ServiceError {
	err := s.ResourceManager.Update(id, update)
	return handleResourceErrors(err)
}

// NewBaseCrudDbService creates a new BaseCrudDbService
func NewBaseCrudDbService(logger *logger.GooglyLogger, rm resource.DbResourceManager) *BaseCrudDbService {
	ser := NewBaseDbService(logger, rm)
	return &BaseCrudDbService{ser}
}
