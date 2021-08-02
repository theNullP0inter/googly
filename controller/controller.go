package controller

import (
	"github.com/theNullP0inter/account-management/logger"
)

type ControllerInterface interface {
}

type Controller struct {
	Logger logger.LoggerInterface
}

func NewController(logger logger.LoggerInterface) ControllerInterface {
	return &Controller{logger}
}
