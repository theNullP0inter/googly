package service

import (
	"github.com/theNullP0inter/account-management/logger"
	"github.com/theNullP0inter/account-management/resource"
)

type RdbCrudImplementorInterface interface {
	CrudServiceImplementorInterface
}

type RdbCrudImplementor struct {
	*Service
	RdbCrudImplementorInterface
	ResourceManager resource.RdbCrudResourceManagerIntereface
}

func (s *RdbCrudImplementor) GetModel() DataInterface {
	return s.ResourceManager.GetModel()
}

func (s *RdbCrudImplementor) Create(m DataInterface) (DataInterface, error) {
	return s.ResourceManager.Create(m)

}

func (s *RdbCrudImplementor) Delete(id DataInterface) error {
	return s.ResourceManager.Delete(id)

}

func (s *RdbCrudImplementor) GetItem(id DataInterface) (DataInterface, error) {
	return s.ResourceManager.Get(id)
}

func (s *RdbCrudImplementor) GetList(req DataInterface) (DataInterface, error) {
	return s.ResourceManager.List(req)

}

func (s *RdbCrudImplementor) Update(item DataInterface) (DataInterface, error) {
	return s.ResourceManager.Update(item)
}

func NewRdbCrudImplementor(logger logger.LoggerInterface, rm resource.RdbCrudResourceManagerIntereface) *RdbCrudImplementor {
	service := NewService(logger)
	return &RdbCrudImplementor{
		Service:         service,
		ResourceManager: rm,
	}
}
