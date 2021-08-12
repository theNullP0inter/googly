package controller

import (
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/theNullP0inter/googly/logger"
	"github.com/theNullP0inter/googly/service"
)

type GinCrudControllerInterface interface {
	GinControllerInterface
	Create(context *gin.Context)
	Get(context *gin.Context)
	List(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type GinCrudController struct {
	*GinController
	Service                 service.CrudInterface
	QueryParametersHydrator GinQueryParametersHydratorInterface

	CreateRequest    SerializerInterface
	UpdateRequest    SerializerInterface
	ListSerializer   SerializerInterface
	DetailSerializer SerializerInterface

	// TODO: add more options like allowed_methods, etc in a map. These can be implemented in AddRoutes()
}

func (s *GinCrudController) AddRoutes(router *gin.RouterGroup) {
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

func (s *GinCrudController) CopyAndSendSuccess(c *gin.Context, data, i SerializerInterface) {
	res := reflect.New(reflect.TypeOf(i)).Interface()
	copier.Copy(res, data)
	s.HttpReplySuccess(c, res)
}

func (s *GinCrudController) Create(c *gin.Context) {
	serializer := reflect.New(reflect.TypeOf(s.CreateRequest)).Interface()

	if err := c.ShouldBindJSON(serializer); err != nil {
		s.Logger.Error(err)
		s.HttpReplyGinBindError(c, err)
		return
	}

	data, err := s.Service.Create(serializer)

	if err != nil {
		s.Logger.Error(err)
		s.HttpReplyServiceError(c, err)
		return
	}
	if s.DetailSerializer != nil {
		s.CopyAndSendSuccess(c, data, s.DetailSerializer)
		return
	}

	s.HttpReplySuccess(c, data)
}

func (s *GinCrudController) Get(c *gin.Context) {
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

func (s *GinCrudController) List(c *gin.Context) {
	query_params, err := s.QueryParametersHydrator.Hydrate(c)
	if err != nil {
		s.HttpReplyGinBindError(c, err)
		return

	}
	data, serr := s.Service.GetList(query_params)

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

func (s *GinCrudController) Update(c *gin.Context) {
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

func (s *GinCrudController) Delete(c *gin.Context) {
	id := c.Param("id")

	serr := s.Service.Delete(id)
	if serr != nil {
		s.HttpReplyServiceError(c, serr)
		return
	}

	s.HttpReplySuccess(c, "deleted")

}

func NewGinCrudController(logger logger.LoggerInterface,
	service service.CrudInterface,
	hydrator GinQueryParametersHydratorInterface,
	create_request SerializerInterface,
	update_request SerializerInterface,
	list_serializer SerializerInterface,
	detail_serializer SerializerInterface,
) *GinCrudController {
	gin_controller := NewGinController(logger)

	return &GinCrudController{
		GinController:           gin_controller,
		Service:                 service,
		QueryParametersHydrator: hydrator,

		CreateRequest:    create_request,
		UpdateRequest:    update_request,
		ListSerializer:   list_serializer,
		DetailSerializer: detail_serializer,
	}
}
