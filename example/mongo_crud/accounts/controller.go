package accounts

import (
	"github.com/gin-gonic/gin"
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

type AccountCrudHttpConnector struct {
}

// Add Custom Actions
func (s *AccountCrudHttpConnector) AddActions(router *gin.RouterGroup) {

}

func NewAccountController(s AccountServiceInterface, logger logger.LoggerInterface) *AccountController {
	hydrator := controller.NewGinPaginatedQueryParametersHydrator(logger)
	crud_http_connector := &AccountCrudHttpConnector{}

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
		crud_http_connector,

		create_request, update_request, nil, nil,
		// nil, nil, nil, nil // This blocks create & update APIs

		// Continution from above => comment the nil in the above line and uncomment below
		// update_request, // This is required for PUT update request to be active
		// list_serializer,
		// detail_serializer,
	)
	return &AccountController{
		controller,
	}
}