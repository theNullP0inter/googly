package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/theNullP0inter/googly/controller"
)

// GinQueryParametersHydrator extracts QueryParameters from gin context
//
// // Hydrate will extract query params from gin context and bind them to any struct
type GinQueryParametersHydrator interface {
	Hydrate(context *gin.Context) (controller.QueryParameters, error)
}

// PaginatedGinQueryParametersHydrator is an implementation of GinQueryParametersHydrator
// that can hydrate to pagination parameters defined in controller.PaginationQueryParameters
type PaginatedGinQueryParametersHydrator struct {
}

// Hydrate will extract query params from gin context and bind them to controller.PaginationQueryParameters
func (c PaginatedGinQueryParametersHydrator) Hydrate(context *gin.Context) (controller.QueryParameters, error) {
	var parameters controller.PaginationQueryParameters
	err := context.ShouldBindQuery(&parameters)
	if err != nil {
		return nil, err
	}
	return &parameters, nil
}

// NewPaginatedGinQueryParametersHydrator will create a new PaginatedGinQueryParametersHydrator
func NewPaginatedGinQueryParametersHydrator() *PaginatedGinQueryParametersHydrator {
	return &PaginatedGinQueryParametersHydrator{}
}
