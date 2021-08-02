package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/theNullP0inter/account-management/logger"
	"github.com/theNullP0inter/account-management/resource"
)

type QueryParametersHydratorInterface interface {
	Hydrate(context *gin.Context) (resource.ListParametersInterface, error)
}

type CrudParametersHydrator struct {
	Logger logger.LoggerInterface
	QueryParametersHydratorInterface
}

func (c CrudParametersHydrator) Hydrate(context *gin.Context) (resource.ListParametersInterface, error) {
	var parameters resource.CrudListParameters
	err := context.ShouldBindQuery(&parameters)
	return &parameters, err
}

func NewBaseParametersHydrator(logger logger.LoggerInterface) *CrudParametersHydrator {
	return &CrudParametersHydrator{Logger: logger}
}
