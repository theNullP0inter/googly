package controller

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theNullP0inter/googly/logger"
)

func TestNewBaseController(t *testing.T) {
	l := new(logger.MockGooglyLogger)
	ctl := NewBaseController(l)

	assert.Equal(t, l, ctl.Logger)
}
