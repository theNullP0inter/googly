package resource

import "github.com/stretchr/testify/mock"

type MockResource struct {
	mock.Mock
}

type MockResourceManager struct {
	mock.Mock
}

func (m *MockResourceManager) GetResource() Resource {
	m.Called()
	return new(MockResource)
}
