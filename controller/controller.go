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
	Logger logger.GooglyLoggerInterface
}

// NewBaseController creates a new BaseController
func NewBaseController(logger logger.GooglyLoggerInterface) *BaseController {
	return &BaseController{logger}
}
