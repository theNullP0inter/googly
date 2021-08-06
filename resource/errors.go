package resource

import "errors"

var (
	ErrResourceNotFound   = errors.New("resource not found")
	ErrInternal           = errors.New("internal error")
	ErrInvalidQuery       = errors.New("invalid query")
	ErrInvalidTransaction = errors.New("invalid transaction")
	ErrUniqueConstraint   = errors.New("similar resource already exist")
	ErrInvalidFormat      = errors.New("invalid format")
)
