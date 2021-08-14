package accounts

import (
	"github.com/theNullP0inter/googly/controller"
	"github.com/theNullP0inter/googly/db/model"
	"github.com/theNullP0inter/googly/logger"
)

type AccountSerializer struct {
	ID       model.BinID `copier:"must" json:"id"`
	Username string      `copier:"must" json:"username"`
}

type AccountCreateRequestSerializer struct {
	Username string `copier:"must" json:"username" bson:"username" `
}

type AccountController struct {
	*controller.GinCrudController
}

func NewAccountController(s AccountServiceInterface, logger logger.GooglyLoggerInterface) *AccountController {
	hydrator := controller.NewGinPaginatedQueryParametersHydrator(logger)

	// Add these to customize your response.
	//
	// by default, model object is sent as response. http response cann be tweaked by using the `json` tag on the model fields
	//
	//
	// var listSerializer []AccountSerializer
	// var detailSerializer AccountSerializer
	var updateRequest AccountCreateRequestSerializer
	var createRequest AccountCreateRequestSerializer

	controller := controller.NewGinCrudController(
		logger, s, hydrator,

		createRequest, updateRequest, nil, nil,
		// nil, nil, nil, nil // This blocks create & update APIs
	)
	return &AccountController{
		controller,
	}
}
