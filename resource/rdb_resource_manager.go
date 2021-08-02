package resource

import (
	go_errors "errors"
	"reflect"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/copier"
	"github.com/theNullP0inter/googly/errors"
	"github.com/theNullP0inter/googly/logger"
	"github.com/theNullP0inter/googly/model"
	"gorm.io/gorm"
)

type RdbResourceManager struct {
	*ResourceManager
	Rdb          *gorm.DB
	Model        model.BaseModelInterface
	QueryBuilder RdbListQueryBuilderInterface
}

func (s RdbResourceManager) Create(m DataInterface) (DataInterface, *errors.GooglyError) {
	item := reflect.New(reflect.TypeOf(s.GetModel())).Interface()
	copier.Copy(item, m)
	result := s.Rdb.Create(item)

	if result.Error != nil {
		var mysqlErr *mysql.MySQLError
		if go_errors.As(result.Error, &mysqlErr) && mysqlErr.Number == 1062 {
			return nil, errors.NewUniqueConstraintError("", result.Error)
		}
		return nil, errors.NewInternalError(result.Error)
	}
	return item, nil
}

func (s RdbResourceManager) GetResource() Resource {
	return s.Model
}

func (s RdbResourceManager) GetModel() DataInterface {
	return s.Model
}
func (s RdbResourceManager) Get(id DataInterface) (DataInterface, *errors.GooglyError) {
	item := reflect.New(reflect.TypeOf(s.GetModel())).Interface()
	bin_id, ok := id.(model.BinID)
	if !ok {
		return nil, errors.NewBinIdAssertionError(id)
	}
	b_id, _ := bin_id.MarshalBinary()
	err := s.Rdb.Where("id = ?", b_id).First(item).Error
	if err != nil {
		return nil, errors.NewResourceNotFoundError("", err)
	}
	return item, nil
}

func (s RdbResourceManager) Update(data DataInterface) (DataInterface, *errors.GooglyError) {
	item := reflect.New(reflect.TypeOf(s.GetModel())).Interface()
	copier.Copy(item, data)
	result := s.Rdb.Save(item)
	if result.Error != nil {
		var mysqlErr *mysql.MySQLError
		if go_errors.As(result.Error, &mysqlErr) && mysqlErr.Number == 1062 {
			return nil, errors.NewUniqueConstraintError("", result.Error)
		}
		return nil, errors.NewInternalError(result.Error)
	}
	return item, nil
}
func (s RdbResourceManager) Delete(id DataInterface) *errors.GooglyError {
	item, err := s.Get(id)
	if err != nil {
		return err
	}
	s.Rdb.Delete(item)
	return nil
}

func (s RdbResourceManager) List(parameters DataInterface) (DataInterface, *errors.GooglyError) {

	items := reflect.New(reflect.SliceOf(reflect.TypeOf(s.GetModel()))).Interface()
	result, err := s.QueryBuilder.ListQuery(parameters)
	if err != nil {
		return nil, errors.NewInternalError(err)
	}
	result = result.Find(items)
	if result.Error != nil {
		return nil, errors.NewResourceNotFoundError("", result.Error)
	}

	return items, nil
}

func NewRdbResourceManager(
	db *gorm.DB,
	logger logger.LoggerInterface,
	model model.BaseModelInterface,
	query_builder PaginatedRdbListQueryBuilderInterface,
) DbResourceManagerIntereface {
	resource_manager := NewResourceManager(logger, model)
	return &RdbResourceManager{
		ResourceManager: resource_manager.(*ResourceManager),
		Rdb:             db,
		Model:           model,
		QueryBuilder:    query_builder,
	}

}
