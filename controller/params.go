package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/theNullP0inter/account-management/resource"
)

type ParametersHydratorInterface interface {
	Hydrate(context *gin.Context) (resource.ListParametersInterface, error)
}

type BaseParametersHydrator struct {
	Logger *logrus.Logger
	ParametersHydratorInterface
}

func (c BaseParametersHydrator) Hydrate(context *gin.Context) (resource.ListParametersInterface, error) {
	var parameters resource.CrudListParameters
	err := context.ShouldBindQuery(&parameters)
	return &parameters, err
}

func NewBaseParametersHydrator(logger *logrus.Logger) *BaseParametersHydrator {
	return &BaseParametersHydrator{Logger: logger}
}
