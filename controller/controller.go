package controller

import (
	"github.com/theNullP0inter/googly/logger"
)

// Controller is an empty interface.
//
// Controller can controll anything ( also nothing )
type Controller interface {
}

// BaseController is a Controller with just a Logger
type BaseController struct {
	Logger *logger.GooglyLogger
}

// NewBaseController creates a new BaseController
func NewBaseController(logger *logger.GooglyLogger) *BaseController {
	return &BaseController{logger}
}
