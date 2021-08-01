package controller

import (
	"github.com/sirupsen/logrus"
)

type ControllerInterface interface {
}

type Controller struct {
	Logger *logrus.Logger
}

func NewController(logger *logrus.Logger) *Controller {
	return &Controller{logger}
}
