package controller

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theNullP0inter/googly/service"
)

func TestNewHttpErrorFromServiceError(t *testing.T) {
	c := 500
	m := "mock"
	err := errors.New(m)

	serr := service.NewServiceError(c, m, err, nil)
	herr := NewHttpErrorFromServiceError(serr)

	assert.Equal(t, c, herr.Code)
	assert.Equal(t, m, herr.Message)
	assert.Equal(t, err, herr.Err)
}
