package resource_manager

import (
	"fmt"
	"reflect"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ListParametersInterface interface{}

type PaginationParameters struct {
	Page      int    `json:"page,default=0"`
	PageSize  int    `json:"page_size,default=30"`
	OrderBy   string `json:"order_by,default=id"`
	OrderDesc bool   `json:"order_desc,default=false"`
}

type CrudListParameters struct {
	*PaginationParameters
}

const DefaultPageSize = 30

type ListQueryBuilderInterface interface {
	ListQuery(parameters ListParametersInterface) (*gorm.DB, error)
	PaginationQuery(parameters ListParametersInterface) *gorm.DB
}

type BaseListQueryBuilder struct {
	Rdb    *gorm.DB
	Logger *logrus.Logger
	ListQueryBuilderInterface
}

func NewBaseListQueryBuilder(db *gorm.DB, logger *logrus.Logger) *BaseListQueryBuilder {
	return &BaseListQueryBuilder{Rdb: db, Logger: logger}
}

func (c BaseListQueryBuilder) PaginationQuery(parameters ListParametersInterface) *gorm.DB {
	query := c.Rdb

	val := reflect.ValueOf(parameters).Elem()
	if val.Kind() != reflect.Struct {
		c.Logger.Error("Unexpected type of parameters for PaginationQuery")
		return query
	}

	paginationParameters := val.FieldByName("PaginationParameters")
	hasPaginationParams := paginationParameters.IsValid() && !paginationParameters.IsNil()

	var page int64
	page = 0
	if hasPaginationParams {
		pageValue := val.FieldByName("Page")
		if !pageValue.IsValid() || pageValue.Kind() != reflect.Int {
			c.Logger.Error("Page is not specified correctly in listQuery")
		} else {
			page = pageValue.Int()
		}
	}

	var pageSize int64
	pageSize = DefaultPageSize
	if hasPaginationParams {
		pageSizeValue := val.FieldByName("PageSize")
		if !pageSizeValue.IsValid() || pageSizeValue.Kind() != reflect.Int {
			c.Logger.Error("PageSize is not specified in listQuery")
		} else {
			pageSize = pageSizeValue.Int()
		}
	}

	limit := pageSize
	offset := page * pageSize
	query = query.Offset(int(offset)).Limit(int(limit))

	var orderBy string
	if hasPaginationParams {
		orderByValue := val.FieldByName("OrderBy")
		if orderByValue.IsValid() && orderByValue.Kind() == reflect.String {
			orderBy = orderByValue.String()
		}
	}

	var orderDesc = false
	if hasPaginationParams {
		orderDescValue := val.FieldByName("OrderDesc")
		if orderDescValue.IsValid() && orderDescValue.Kind() == reflect.Bool {
			orderDesc = orderDescValue.Bool()
		}
	}

	if len(orderBy) > 0 {
		if orderDesc {
			query = query.Order(fmt.Sprintf("%s DESC", orderBy))
		} else {
			query = query.Order(fmt.Sprintf("%s ASC", orderBy))
		}
	}

	return query
}

func (c BaseListQueryBuilder) ListQuery(parameters ListParametersInterface) (*gorm.DB, error) {
	return c.PaginationQuery(parameters), nil
}
