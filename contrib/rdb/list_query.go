package rdb

import (
	"fmt"
	"reflect"

	"github.com/theNullP0inter/googly/logger"
	"github.com/theNullP0inter/googly/resource"
	"gorm.io/gorm"
)

const DefaultPageSize = 30

// RdbListQueryBuilderInterface
type RdbListQueryBuilderInterface interface {
	ListQuery(resource.ListQuery) (*gorm.DB, error)
}

type PaginatedRdbListQueryBuilderInterface interface {
	RdbListQueryBuilderInterface
	PaginationQuery(resource.ListQuery) *gorm.DB
}

type PaginatedRdbListQueryBuilder struct {
	Rdb    *gorm.DB
	Logger logger.GooglyLoggerInterface
}

func NewPaginatedRdbListQueryBuilder(db *gorm.DB, logger logger.GooglyLoggerInterface) *PaginatedRdbListQueryBuilder {
	return &PaginatedRdbListQueryBuilder{Rdb: db, Logger: logger}
}

func (qb *PaginatedRdbListQueryBuilder) ListQuery(parameters resource.ListQuery) (*gorm.DB, error) {
	return qb.PaginationQuery(parameters), nil
}

func (qb *PaginatedRdbListQueryBuilder) PaginationQuery(parameters resource.ListQuery) *gorm.DB {
	query := qb.Rdb

	val := reflect.ValueOf(parameters).Elem()
	if val.Kind() != reflect.Struct {
		qb.Logger.Errorf("Unexpected type of parameters for PaginationQuery")
		return query
	}

	paginationParameters := val.FieldByName("PaginatedListQuery")
	hasPaginationParams := paginationParameters.IsValid() && !paginationParameters.IsNil()

	var page int64
	page = 0
	if hasPaginationParams {
		pageValue := val.FieldByName("Page")
		if !pageValue.IsValid() || pageValue.Kind() != reflect.Int {
			qb.Logger.Errorf("Page is not specified correctly in listQuery")
		} else {
			page = pageValue.Int()
		}
	}

	var pageSize int64
	pageSize = DefaultPageSize
	if hasPaginationParams {
		pageSizeValue := val.FieldByName("PageSize")
		if !pageSizeValue.IsValid() || pageSizeValue.Kind() != reflect.Int {
			qb.Logger.Errorf("PageSize is not specified in listQuery")
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
