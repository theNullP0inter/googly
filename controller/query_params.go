package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/theNullP0inter/account-management/resource"
)

type QueryParametersHydratorInterface interface {
	Hydrate(context *gin.Context) (resource.ListParametersInterface, error)
}

type CrudParametersHydrator struct {
	Logger *logrus.Logger
	QueryParametersHydratorInterface
}

func (c CrudParametersHydrator) Hydrate(context *gin.Context) (resource.ListParametersInterface, error) {
	var parameters resource.CrudListParameters
	err := context.ShouldBindQuery(&parameters)
	return &parameters, err
}

func NewBaseParametersHydrator(logger *logrus.Logger) *CrudParametersHydrator {
	return &CrudParametersHydrator{Logger: logger}
}
