package mongo_db

import "errors"

var (
	ErrParseBson = errors.New("error parsing bson request")
)
