package service

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/theNullP0inter/googly/logger"
	"github.com/theNullP0inter/googly/resource"
)

// TODO: Test handleResourceErrors()

func TestBaseCrudDbService(t *testing.T) {
	l := new(logger.MockGooglyLogger)
	r := new(resource.MockResource)
	c := new(resource.MockCrudImplementor)

	rm := resource.NewBaseCrudResourceManager(l, r, c)

	s := NewBaseCrudDbService(l, rm)

	c.On("Create", mock.Anything).Return(mock.Anything, nil)
	s.Create(new(MockDataInterface))
	c.AssertExpectations(t)

	c.On("List", mock.Anything).Return(mock.Anything, nil)
	s.GetList(new(MockDataInterface))
	c.AssertExpectations(t)

	c.On("Get", mock.Anything).Return(mock.Anything, nil)
	s.GetItem(new(MockDataInterface))
	c.AssertExpectations(t)

	c.On("Update", mock.Anything, mock.Anything).Return(nil)
	s.Update(new(MockDataInterface), new(MockDataInterface))
	c.AssertExpectations(t)

	c.On("Delete", mock.Anything).Return(nil)
	s.Delete(new(MockDataInterface))
	c.AssertExpectations(t)

}
