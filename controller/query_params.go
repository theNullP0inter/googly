package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/theNullP0inter/googly/logger"
	"github.com/theNullP0inter/googly/resource"
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
	if err != nil {
		return nil, err
	}
	return &parameters, nil
}

func NewBaseParametersHydrator(logger logger.LoggerInterface) *CrudParametersHydrator {
	return &CrudParametersHydrator{Logger: logger}
}
