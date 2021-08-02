package controller

import (
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/theNullP0inter/account-management/errors"
	"github.com/theNullP0inter/account-management/logger"
	"github.com/theNullP0inter/account-management/model"
	"github.com/theNullP0inter/account-management/service"
)

type CrudHttpControllerConnectorInterface interface {
	AddActions(*gin.RouterGroup)
}

type CrudHttpControllerInterface interface {
	HttpControllerInterface
	Create(context *gin.Context)
	Get(context *gin.Context)
	List(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type CrudHttpController struct {
	*HttpController
	CrudHttpControllerConnectorInterface
	Service            service.CrudServiceInterface
	ParametersHydrator QueryParametersHydratorInterface

	CreateRequest    SerializerInterface
	UpdateRequest    SerializerInterface
	ListSerializer   SerializerInterface
	DetailSerializer SerializerInterface

	// TODO: add more options like allowed_methods, etc in a map. These can be implemented in AddRoutes()
}

func (s *CrudHttpController) AddRoutes(router *gin.RouterGroup) {
	if s.CreateRequest != nil {
		router.POST("/", s.Create)
	}
	if s.UpdateRequest != nil {
		router.PUT("/:id", s.Update)
	}

	router.GET("/", s.List)
	router.GET("/:id", s.Get)

	router.DELETE("/:id", s.Delete)
	s.AddActions(router)
}

func (s *CrudHttpController) CopyAndSendSuccess(c *gin.Context, data, i SerializerInterface) {
	res := reflect.New(reflect.TypeOf(i)).Interface()
	copier.Copy(res, data)
	s.HttpReplySuccess(c, res)
}

func (s *CrudHttpController) Create(c *gin.Context) {
	serializer := reflect.New(reflect.TypeOf(s.CreateRequest)).Interface()

	if err := c.ShouldBindJSON(serializer); err != nil {
		s.Logger.Error(err)
		s.HttpReplyBindError(c, err)
		return
	}

	// copier.Copy(res, data)

	data, err := s.Service.Create(serializer)

	if err != nil {
		s.Logger.Error(err)
		s.HttpReplyGenericBadRequest(c)
		return
	}
	if s.DetailSerializer != nil {
		s.CopyAndSendSuccess(c, data, s.DetailSerializer)
		return
	}

	s.HttpReplySuccess(c, data)
}

func (s *CrudHttpController) Get(c *gin.Context) {
	id, err := model.StringToBinID(c.Param("id"))
	if err != nil {
		s.Logger.Error(err)
		s.HttpReplyBadRequestFromError(c, err)
		return
	}

	data, err := s.Service.GetItem(id)
	if data == nil {
		s.HttpReplyNotFound(c)
		return
	}

	if s.DetailSerializer != nil {
		s.CopyAndSendSuccess(c, data, s.DetailSerializer)
		return
	}

	s.HttpReplySuccess(c, data)

}

func (s *CrudHttpController) List(c *gin.Context) {
	params, err := s.ParametersHydrator.Hydrate(c)
	if err != nil {
		s.HttpReplyBadRequestFromError(c, err)
		return

	}
	data, err := s.Service.GetList(params)

	if err != nil {
		s.HttpReplyNotFound(c)
		return
	}

	if s.ListSerializer != nil {
		s.CopyAndSendSuccess(c, data, s.ListSerializer)
		return
	}

	s.HttpReplySuccess(c, data)

}

func (s *CrudHttpController) Update(c *gin.Context) {
	id, err := model.StringToBinID(c.Param("id"))
	if err != nil {
		s.Logger.Error(err)
		s.HttpReplyBadRequestFromError(c, err)
		return
	}

	serializer := reflect.New(reflect.TypeOf(s.UpdateRequest)).Interface()
	if err := c.ShouldBindJSON(serializer); err != nil {
		s.HttpReplyBadRequestFromError(c, err)
		return
	}

	item, _ := s.Service.GetItem(id)

	if item == nil {
		s.HttpReplyNotFound(c)
		return
	}

	copier.Copy(item, serializer)
	item, err = s.Service.Update(item)

	if item == nil {
		s.HttpReplyInternalError(c, err)
		return
	}

	if s.DetailSerializer != nil {
		s.CopyAndSendSuccess(c, item, s.DetailSerializer)
		return
	}

	s.HttpReplySuccess(c, item)

}

func (s *CrudHttpController) Delete(c *gin.Context) {
	id, err := model.StringToBinID(c.Param("id"))
	if err != nil {
		s.Logger.Error(err)
		s.HttpReplyBadRequestFromError(c, err)
		return
	}
	err = s.Service.Delete(id)
	if _, is_err := err.(*errors.GogetaError); is_err {
		s.HttpReplyInternalError(c, err)
		return
	}

	s.HttpReplySuccess(c, "deleted")

}

func NewCrudHttpController(logger logger.LoggerInterface,
	service service.CrudServiceInterface,
	hydrator *CrudParametersHydrator,
	crud_http_connector CrudHttpControllerConnectorInterface,
	create_request SerializerInterface,
	update_request SerializerInterface,
	list_serializer SerializerInterface,
	detail_serializer SerializerInterface,
) *CrudHttpController {
	http_controller := NewHttpController(logger)

	return &CrudHttpController{
		HttpController:                       http_controller.(*HttpController),
		CrudHttpControllerConnectorInterface: crud_http_connector,
		Service:                              service,
		ParametersHydrator:                   hydrator,

		CreateRequest:    create_request,
		UpdateRequest:    update_request,
		ListSerializer:   list_serializer,
		DetailSerializer: detail_serializer,
	}
}
