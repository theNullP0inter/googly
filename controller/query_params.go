package controller

import (
	"github.com/theNullP0inter/googly/resource"
)

// QueryParameters is any thing that controller receives in GET request or on similar lines
type QueryParameters interface {
}

// PaginationQueryParameters is the set of query params that should be used to get pagination
//
// To make it easier, it is set to pagination params accepted by DbResourceManager for a PaginatedListQuery
type PaginationQueryParameters struct {
	*resource.PaginatedListQuery
}
