package resource

import "errors"

var (
	ErrResourceNotFound   = errors.New("resource not found")
	ErrNoModification     = errors.New("no modification")
	ErrInternal           = errors.New("internal error")
	ErrInvalidQuery       = errors.New("invalid query")
	ErrParseBson          = errors.New("error parsing bson request")
	ErrInvalidTransaction = errors.New("invalid transaction")
	ErrUniqueConstraint   = errors.New("similar resource already exist")
	ErrInvalidFormat      = errors.New("invalid format")
)
