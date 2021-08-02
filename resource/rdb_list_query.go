package resource

import (
	"fmt"
	"reflect"

	"github.com/theNullP0inter/googly/logger"
	"gorm.io/gorm"
)

const DefaultPageSize = 30

type ListParametersInterface interface{ ParameterInterface }
type RdbListQueryBuilderInterface interface {
	ListQuery(parameters ListParametersInterface) (*gorm.DB, error)
}
type PaginatedRdbListQueryBuilderInterface interface {
	RdbListQueryBuilderInterface
	PaginationQuery(parameters ListParametersInterface) *gorm.DB
}

type PaginatedRdbListQueryBuilder struct {
	Rdb    *gorm.DB
	Logger logger.LoggerInterface
	PaginatedRdbListQueryBuilderInterface
}

func NewPaginatedRdbListQueryBuilder(db *gorm.DB, logger logger.LoggerInterface) *PaginatedRdbListQueryBuilder {
	return &PaginatedRdbListQueryBuilder{Rdb: db, Logger: logger}
}

// modify for a new type of query builder
func (c PaginatedRdbListQueryBuilder) ListQuery(parameters ListParametersInterface) (*gorm.DB, error) {
	return c.PaginationQuery(parameters), nil
}

func (c PaginatedRdbListQueryBuilder) PaginationQuery(parameters ListParametersInterface) *gorm.DB {
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
