package resource

import "github.com/stretchr/testify/mock"

type MockDbResourceManager struct {
	*mock.Mock
	*MockCrudImplementor
	*MockResourceManager
}
