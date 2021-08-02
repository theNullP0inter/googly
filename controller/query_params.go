package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/theNullP0inter/googly/errors"
	"github.com/theNullP0inter/googly/logger"
	"github.com/theNullP0inter/googly/resource"
)

type QueryParametersHydratorInterface interface {
	Hydrate(context *gin.Context) (resource.ListParametersInterface, *errors.GooglyError)
}

type CrudParametersHydrator struct {
	Logger logger.LoggerInterface
	QueryParametersHydratorInterface
}

func (c CrudParametersHydrator) Hydrate(context *gin.Context) (resource.ListParametersInterface, *errors.GooglyError) {
	var parameters resource.CrudListParameters
	err := context.ShouldBindQuery(&parameters)
	if err != nil {
		return nil, errors.NewParamsHydrationError(err)
	}
	return &parameters, nil
}

func NewBaseParametersHydrator(logger logger.LoggerInterface) *CrudParametersHydrator {
	return &CrudParametersHydrator{Logger: logger}
}
