package service

import (
	"github.com/sirupsen/logrus"
	"github.com/theNullP0inter/account-management/model"
	"github.com/theNullP0inter/account-management/resource"
)

type ModelCrudServiceInterface interface {
	ServiceInterface

	GetModel() DataInterface
	GetItem(id model.BinID) (DataInterface, error)
	GetList(req DataInterface) (DataInterface, error)
	Create(req DataInterface) (DataInterface, error)
	Update(item DataInterface) (DataInterface, error)
	Delete(id model.BinID) error
}

type ModelCrudService struct {
	*Service
	ResourceManager resource.ModelResourceManagerInterface
}

func (s *ModelCrudService) GetModel() DataInterface {
	return s.ResourceManager.GetModel()
}

func (s *ModelCrudService) Create(m DataInterface) (DataInterface, error) {
	return s.ResourceManager.Create(m)

}

func (s *ModelCrudService) Delete(id model.BinID) error {
	return s.ResourceManager.Delete(id)

}

func (s *ModelCrudService) GetItem(id model.BinID) (DataInterface, error) {
	return s.ResourceManager.Get(id)
}

func (s *ModelCrudService) GetList(req DataInterface) (DataInterface, error) {
	return s.ResourceManager.List(req)

}

func (s *ModelCrudService) Update(item DataInterface) (DataInterface, error) {
	return s.ResourceManager.Update(item)
}

func NewModelCrudService(logger *logrus.Logger, rm resource.ModelResourceManagerInterface) *ModelCrudService {
	service := NewService(logger)
	return &ModelCrudService{
		Service:         service,
		ResourceManager: rm,
	}
}
