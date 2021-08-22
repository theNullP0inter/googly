package service

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServiceError(t *testing.T) {
	m := "error_message"
	c := 200
	err := errors.New("test service error")

	serr := NewServiceError(c, m, err, nil)

	expectedError, _ := json.MarshalIndent(m, "", "  ")
	actualError := serr.Error()

	assert.Equal(t, string(expectedError), actualError)
}

func TestInternalServiceError(t *testing.T) {

	err := errors.New("test service error")
	serr := NewInternalServiceError(err)
	assert.Equal(t, 500, serr.Code)
}

func TestBadRequestError(t *testing.T) {

	err := errors.New("test service error")
	serr := NewBadRequestError(err)
	assert.Equal(t, 400, serr.Code)
}

func TestNotFoundServiceError(t *testing.T) {

	err := errors.New("test service error")
	serr := NewNotFoundServiceError(err)
	assert.Equal(t, 404, serr.Code)
}
