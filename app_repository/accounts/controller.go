package accounts

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/theNullP0inter/account-management/controller"
	"github.com/theNullP0inter/account-management/model"
)

type AccountSerializer struct {
	ID       model.BinID `copier:"must" json:"id"`
	Username string      `copier:"must" json:"username"`
}

type AccountCreateRequestSerializer struct {
	Username string `copier:"must" json:"username"`
}

type AccountController struct {
	*controller.CrudHttpController
}

type AccountCrudHttpConnector struct {
}

// Add Custom Actions
func (s *AccountCrudHttpConnector) AddActions(router *gin.RouterGroup) {

}

func NewAccountController(s AccountServiceInterface, logger *logrus.Logger) *AccountController {
	hydrator := controller.NewBaseParametersHydrator(logger)
	crud_http_connector := &AccountCrudHttpConnector{}

	// Add these to customize your response.
	//
	// by default, model object is sent as response. http response cann be tweaked by using the `json` tag on the model fields
	//
	//
	// var list_serializer []AccountSerializer
	// var detail_serializer AccountSerializer
	// var update_request AccountCreateRequestSerializer
	var create_request AccountCreateRequestSerializer

	controller := controller.NewCrudHttpController(
		logger, s, hydrator,
		crud_http_connector,

		create_request, nil, nil, nil,

		// Continution from above => comment the nil in the above line and uncomment below
		// update_request, // This is required for PUT update request to be active
		// list_serializer,
		// detail_serializer,
	)
	return &AccountController{
		controller,
	}
}
