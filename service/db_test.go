package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theNullP0inter/googly/logger"
	"github.com/theNullP0inter/googly/resource"
)

func TestBaseDbService(t *testing.T) {
	l := logger.NewGooglyLogger()
	rm := new(resource.MockDbResourceManager)
	s := NewBaseDbService(l, rm)
	assert.Equal(t, s.ResourceManager, rm)

}
