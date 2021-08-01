package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/theNullP0inter/account-management/resource_manager"
)

type ParametersHydratorInterface interface {
	Hydrate(context *gin.Context) (resource_manager.ListParametersInterface, error)
}

type BaseParametersHydrator struct {
	Logger *logrus.Logger
	ParametersHydratorInterface
}

func (c BaseParametersHydrator) Hydrate(context *gin.Context) (resource_manager.ListParametersInterface, error) {
	var parameters resource_manager.CrudListParameters
	err := context.ShouldBindQuery(&parameters)
	return &parameters, err
}

func NewBaseParametersHydrator(logger *logrus.Logger) *BaseParametersHydrator {
	return &BaseParametersHydrator{Logger: logger}
}
