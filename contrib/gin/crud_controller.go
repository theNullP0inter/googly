package gin

import (
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/theNullP0inter/googly/controller"
	"github.com/theNullP0inter/googly/logger"
	"github.com/theNullP0inter/googly/service"
)

// GinCrudController is a GinController that implements CRUD
type GinCrudController interface {
	GinController
	Create(context *gin.Context)
	Get(context *gin.Context)
	List(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

// BaseGinCrudController is a basic implementation of GinCrudController
type BaseGinCrudController struct {
	*BaseGinController
	Service                 service.CrudService
	QueryParametersHydrator GinQueryParametersHydrator

	CreateRequest    controller.Serializer
	UpdateRequest    controller.Serializer
	ListSerializer   controller.Serializer
	DetailSerializer controller.Serializer

	// TODO: add more options like allowed_methods, etc in a map. These can be implemented in AddRoutes()
}

// AddRoutes is required to be implemented by GinControllerIngress.
//
// This will be used by your ingress to Connect
func (s *BaseGinCrudController) AddRoutes(router *gin.RouterGroup) {
	if s.CreateRequest != nil {
		router.POST("/", s.Create)
	}
	if s.UpdateRequest != nil {
		router.PUT("/:id", s.Update)
	}

	router.GET("/", s.List)
	router.GET("/:id", s.Get)

	router.DELETE("/:id", s.Delete)
}

// CopyAndSendSuccess copies data from services's response to a serializer and sends that to the user
func (s *BaseGinCrudController) CopyAndSendSuccess(c *gin.Context, data service.DataInterface, i controller.Serializer) {
	res := reflect.New(reflect.TypeOf(i)).Interface()
	copier.Copy(res, data)
	s.HttpReplySuccess(c, res)
}

// Create will call the services' Create method.
// This will be added to routes only if  s.CreateRequest !=nil
func (s *BaseGinCrudController) Create(c *gin.Context) {
	serializer := reflect.New(reflect.TypeOf(s.CreateRequest)).Interface()

	if err := c.ShouldBindJSON(serializer); err != nil {
		s.HttpReplyGinBindError(c, err)
		return
	}

	data, err := s.Service.Create(serializer)

	if err != nil {
		s.HttpReplyServiceError(c, err)
		return
	}
	if s.DetailSerializer != nil {
		s.CopyAndSendSuccess(c, data, s.DetailSerializer)
		return
	}

	s.HttpReplySuccess(c, data)
}

// Get will call the services' Get method.
// it needs a path param `id` to be passed in
func (s *BaseGinCrudController) Get(c *gin.Context) {
	id := c.Param("id")

	data, serr := s.Service.GetItem(id)
	if data == nil {
		s.HttpReplyServiceError(c, serr)
		return
	}

	if s.DetailSerializer != nil {
		s.CopyAndSendSuccess(c, data, s.DetailSerializer)
		return
	}

	s.HttpReplySuccess(c, data)

}

// List will call the services' List method
// query params will be binded via s.QueryParametersHydrator is
func (s *BaseGinCrudController) List(c *gin.Context) {
	queryParams, err := s.QueryParametersHydrator.Hydrate(c)
	if err != nil {
		s.HttpReplyGinBindError(c, err)
		return

	}
	data, serr := s.Service.GetList(queryParams)

	if serr != nil {
		s.HttpReplyServiceError(c, serr)
		return
	}

	if s.ListSerializer != nil {
		s.CopyAndSendSuccess(c, data, s.ListSerializer)
		return
	}

	s.HttpReplySuccess(c, data)

}

// Update will call the services' Update method
// This will be added to routes only if  s.UpdateRequest !=nil
// it needs a path param `id` to be passed in
func (s *BaseGinCrudController) Update(c *gin.Context) {
	id := c.Param("id")

	serializer := reflect.New(reflect.TypeOf(s.UpdateRequest)).Interface()
	if err := c.ShouldBindJSON(serializer); err != nil {
		s.HttpReplyGinBindError(c, err)
		return
	}

	serr := s.Service.Update(id, serializer)

	if serr != nil {
		s.HttpReplyServiceError(c, serr)
		return
	}

	s.HttpReplySuccess(c, "updated")

}

// Delete will call the services' Delete method
// it needs a path param `id` to be passed in
func (s *BaseGinCrudController) Delete(c *gin.Context) {
	id := c.Param("id")

	serr := s.Service.Delete(id)
	if serr != nil {
		s.HttpReplyServiceError(c, serr)
		return
	}

	s.HttpReplySuccess(c, "deleted")

}

// NewBaseGinCrudController creates a new BaseGinCrudController
func NewBaseGinCrudController(logger logger.GooglyLoggerInterface,
	service service.CrudService,
	hydrator GinQueryParametersHydrator,
	createRequest controller.Serializer,
	updateRequest controller.Serializer,
	listSerializer controller.Serializer,
	detailSerializer controller.Serializer,
) *BaseGinCrudController {
	ginController := NewBaseGinController(logger)

	return &BaseGinCrudController{
		BaseGinController:       ginController,
		Service:                 service,
		QueryParametersHydrator: hydrator,

		CreateRequest:    createRequest,
		UpdateRequest:    updateRequest,
		ListSerializer:   listSerializer,
		DetailSerializer: detailSerializer,
	}
}
