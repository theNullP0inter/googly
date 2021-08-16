package accounts

import (
	gin_contrib "github.com/theNullP0inter/googly/contrib/gin"
	"github.com/theNullP0inter/googly/contrib/rdb"
	"github.com/theNullP0inter/googly/logger"
)

type AccountSerializer struct {
	ID       rdb.BinID `copier:"must" json:"id"`
	Username string    `copier:"must" json:"username"`
}

type AccountCreateRequestSerializer struct {
	Username string `copier:"must" json:"username" bson:"username" `
}

type AccountController struct {
	*gin_contrib.BaseGinCrudController
}

func NewAccountController(s AccountServiceInterface, logger logger.GooglyLoggerInterface) *AccountController {
	hydrator := gin_contrib.NewPaginatedGinQueryParametersHydrator()

	// Add these to customize your response.
	//
	// by default, model object is sent as response. http response cann be tweaked by using the `json` tag on the model fields
	//
	//
	// var listSerializer []AccountSerializer
	// var detailSerializer AccountSerializer
	var updateRequest AccountCreateRequestSerializer
	var createRequest AccountCreateRequestSerializer

	controller := gin_contrib.NewBaseGinCrudController(
		logger, s, hydrator,

		createRequest, updateRequest, nil, nil,
		// nil, nil, nil, nil // This blocks create & update APIs
	)
	return &AccountController{
		controller,
	}
}
