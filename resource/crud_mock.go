package resource

import "github.com/stretchr/testify/mock"

type MockCrudImplementor struct {
	mock.Mock
}

func (c *MockCrudImplementor) Create(m DataInterface) (DataInterface, error) {
	return new(DataInterface), nil
}
func (c *MockCrudImplementor) List(parameters DataInterface) (DataInterface, error) {
	return new(DataInterface), nil
}
func (c *MockCrudImplementor) Get(id DataInterface) (DataInterface, error) {
	return new(DataInterface), nil
}
func (c *MockCrudImplementor) Update(id DataInterface, item DataInterface) error { return nil }
func (c *MockCrudImplementor) Delete(id DataInterface) error                     { return nil }
