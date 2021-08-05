package service

import (
	"github.com/theNullP0inter/googly/logger"
	"github.com/theNullP0inter/googly/resource"
)

type DbServiceInterface interface {
	ServiceInterface
	resource.DbResourceManagerIntereface
}

type DbService struct {
	*Service
	resource.DbResourceManagerIntereface
}

func NewDbService(logger logger.LoggerInterface, rm resource.DbResourceManagerIntereface) *DbService {
	service := NewService(logger)
	return &DbService{service, rm}
}
