package account

import (
	"github.com/gin-gonic/gin"
	"github.com/theNullP0inter/account-management/model"
	"github.com/theNullP0inter/account-management/service"
)

func (s *AccountService) AddRoutes(router *gin.RouterGroup) {
	router.POST("/", s.HttpCreate)
	router.GET("/", s.HttpQuery)
	router.DELETE("/:id", s.HttpDelete)

}

func (s *AccountService) HttpCreate(c *gin.Context) {
	var account_create_request AccountServiceCreateRequest
	c.BindJSON(&account_create_request)

	account_resource, serr := s.Create(&account_create_request)

	if serr != nil {
		serr.RespondToHttp(c)
		return
	}
	c.JSON(200, account_resource)
}

func (s *AccountService) HttpQuery(c *gin.Context) {

	account_resources, serr := s.Query()

	if serr != nil {
		c.JSON(serr.HttpStatus, serr)
		return
	}
	c.JSON(200, account_resources)
}

func (s *AccountService) HttpDelete(c *gin.Context) {
	id, err := model.StringToBinID(c.Param("id"))
	if err != nil {
		service.NewBinIdError().RespondToHttp(c)
		return
	}
	serr := s.Delete(id)

	if serr != nil {
		c.JSON(serr.HttpStatus, serr)
		return
	}
	c.JSON(200, gin.H{})
}
