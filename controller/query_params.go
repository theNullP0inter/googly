package controller

import (
	"github.com/theNullP0inter/googly/resource"
)

// ListPaginationQueryParameters is the set of query params that should be used to get pagination
//
// To make it easier, it is set to pagination params accepted by DbResourceManager
type ListPaginationQueryParameters struct {
	*resource.PaginationQueryParameters
}
