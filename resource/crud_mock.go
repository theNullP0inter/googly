package resource

import "github.com/stretchr/testify/mock"

type MockCrudImplementor struct {
	mock.Mock
}

func (c *MockCrudImplementor) Create(m DataInterface) (DataInterface, error) {
	c.Called(m)
	return new(DataInterface), nil
}
func (c *MockCrudImplementor) List(parameters DataInterface) (DataInterface, error) {
	c.Called(parameters)
	return new(DataInterface), nil
}
func (c *MockCrudImplementor) Get(id DataInterface) (DataInterface, error) {
	c.Called(id)
	return new(DataInterface), nil
}
func (c *MockCrudImplementor) Update(id DataInterface, item DataInterface) error {
	c.Called(id, item)
	return nil
}
func (c *MockCrudImplementor) Delete(id DataInterface) error {
	c.Called(id)
	return nil
}
