package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theNullP0inter/googly/logger"
)

func TestNewBaseService(t *testing.T) {
	l := new(logger.MockGooglyLogger)
	s := NewBaseService(l)
	assert.Equal(t, l, s.Logger, "Not the same Logger")
}
