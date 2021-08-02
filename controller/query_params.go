package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/theNullP0inter/account-management/errors"
	"github.com/theNullP0inter/account-management/logger"
	"github.com/theNullP0inter/account-management/resource"
)

type QueryParametersHydratorInterface interface {
	Hydrate(context *gin.Context) (resource.ListParametersInterface, *errors.GogetaError)
}

type CrudParametersHydrator struct {
	Logger logger.LoggerInterface
	QueryParametersHydratorInterface
}

func (c CrudParametersHydrator) Hydrate(context *gin.Context) (resource.ListParametersInterface, *errors.GogetaError) {
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
