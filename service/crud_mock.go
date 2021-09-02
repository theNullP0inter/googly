package service

import (
	"github.com/stretchr/testify/mock"
	"github.com/theNullP0inter/googly/resource"
)

type MockCrudService struct {
	mock.Mock
}

func (s *MockCrudService) GetItem(id DataInterface) (DataInterface, *ServiceError) {
	s.Called()
	return new(resource.MockResource), nil
}

func (s *MockCrudService) GetList(req DataInterface) (DataInterface, *ServiceError) {
	s.Called()
	return []*resource.MockResource{new(resource.MockResource)}, nil
}

func (s *MockCrudService) Create(req DataInterface) (DataInterface, *ServiceError) {
	s.Called()
	return new(resource.MockResource), nil
}

func (s *MockCrudService) Update(id DataInterface, update DataInterface) *ServiceError {
	s.Called()
	return nil
}

func (s *MockCrudService) Delete(id DataInterface) *ServiceError {
	s.Called()
	return nil
}
