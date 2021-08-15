package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/theNullP0inter/googly/controller"
	"github.com/theNullP0inter/googly/logger"
	"github.com/theNullP0inter/googly/resource"
)

type GinQueryParametersHydratorInterface interface {
	Hydrate(context *gin.Context) (resource.QueryParameters, error)
}

type GinPaginatedQueryParametersHydrator struct {
	Logger logger.GooglyLoggerInterface
	GinQueryParametersHydratorInterface
}

func (c GinPaginatedQueryParametersHydrator) Hydrate(context *gin.Context) (resource.QueryParameters, error) {
	var parameters controller.ListPaginationQueryParameters
	err := context.ShouldBindQuery(&parameters)
	if err != nil {
		return nil, err
	}
	return &parameters, nil
}

func NewGinPaginatedQueryParametersHydrator(logger logger.GooglyLoggerInterface) *GinPaginatedQueryParametersHydrator {
	return &GinPaginatedQueryParametersHydrator{Logger: logger}
}
