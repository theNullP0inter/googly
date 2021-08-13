package controller

import (
	"github.com/theNullP0inter/googly/logger"
)

type ControllerInterface interface {
}

type Controller struct {
	Logger logger.GooglyLoggerInterface
}

func NewController(logger logger.GooglyLoggerInterface) *Controller {
	return &Controller{logger}
}
