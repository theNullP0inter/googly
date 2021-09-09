package resource

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theNullP0inter/googly/logger"
)

func TestNewBaseCrudResourceManager(t *testing.T) {
	r := new(MockResource)
	c := new(MockCrudImplementor)
	l := logger.NewGooglyLogger()
	rm := NewBaseCrudResourceManager(l, r, c)

	assert.Equal(t, rm.CrudImplementor, c)
	assert.Equal(t, rm.Resource, r)

	assert.NotNil(t, rm.BaseResourceManager)
	assert.Equal(t, rm.GetResource(), r)
}
