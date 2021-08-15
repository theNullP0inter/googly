package resource

import "errors"

// List of  errors that can be raised by a resource manager
var (
	ErrResourceNotFound   = errors.New("resource not found")
	ErrNoModification     = errors.New("no modification")
	ErrInternal           = errors.New("internal error")
	ErrInvalidQuery       = errors.New("invalid query")
	ErrInvalidTransaction = errors.New("invalid transaction")
	ErrUniqueConstraint   = errors.New("similar resource already exist")
	ErrInvalidFormat      = errors.New("invalid format")
)
