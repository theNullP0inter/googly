package accounts

import (
	"github.com/theNullP0inter/googly/controller"
	"github.com/theNullP0inter/googly/logger"
	"github.com/theNullP0inter/googly/model"
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

func NewAccountController(s AccountServiceInterface, logger logger.LoggerInterface) *AccountController {
	hydrator := controller.NewGinPaginatedQueryParametersHydrator(logger)

	// Add these to customize your response.
	//
	// by default, model object is sent as response. http response cann be tweaked by using the `json` tag on the model fields
	//
	//
	// var list_serializer []AccountSerializer
	// var detail_serializer AccountSerializer
	var update_request AccountCreateRequestSerializer
	var create_request AccountCreateRequestSerializer

	controller := controller.NewGinCrudController(
		logger, s, hydrator,

		create_request, update_request, nil, nil,
		// nil, nil, nil, nil // This blocks create & update APIs
	)
	return &AccountController{
		controller,
	}
}
