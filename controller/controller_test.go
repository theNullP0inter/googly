package controller

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theNullP0inter/googly/logger"
)

func TestNewBaseController(t *testing.T) {
	l := logger.NewGooglyLogger()
	ctl := NewBaseController(l)

	assert.Equal(t, l, ctl.Logger)
}
