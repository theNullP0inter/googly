package mongo_db

import "errors"

// Errors given by this module
var (
	ErrParseBson = errors.New("error parsing bson request")
)
