package resource

import (
	"errors"
	"reflect"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/copier"
	"github.com/theNullP0inter/account-management/logger"
	"github.com/theNullP0inter/account-management/model"
	"gorm.io/gorm"
)

type RdbCrudResourceManagerIntereface interface {
	CrudResourceManagerInterface
	GetModel() DataInterface
}

type RdbCrudResourceManager struct {
	*ResourceManager
	RdbCrudResourceManagerIntereface
	Rdb          *gorm.DB
	Model        model.BaseModelInterface
	QueryBuilder RdbListQueryBuilderInterface
}

func (s RdbCrudResourceManager) Create(m DataInterface) (DataInterface, error) {
	item := reflect.New(reflect.TypeOf(s.GetModel())).Interface()
	copier.Copy(item, m)
	result := s.Rdb.Create(item)

	if result.Error != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(result.Error, &mysqlErr) && mysqlErr.Number == 1062 {
			return nil, errors.New(UNIQUE_CONSTRAINT_ERROR)
		}
		return nil, errors.New(INTERNAL_ERROR)
	}
	return item, nil
}

func (s RdbCrudResourceManager) GetResource() Resource {
	return s.Model
}

func (s RdbCrudResourceManager) GetModel() DataInterface {
	return s.Model
}
func (s RdbCrudResourceManager) Get(id DataInterface) (DataInterface, error) {
	item := reflect.New(reflect.TypeOf(s.GetModel())).Interface()
	bin_id, ok := id.(model.BinID)
	if !ok {
		return nil, errors.New(BIN_ID_ASSERTION_FAILED)
	}
	b_id, _ := bin_id.MarshalBinary()
	err := s.Rdb.Where("id = ?", b_id).First(item).Error
	return item, err
}

func (s RdbCrudResourceManager) Update(item DataInterface) (DataInterface, error) {
	result := s.Rdb.Save(item)
	if result.Error != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(result.Error, &mysqlErr) && mysqlErr.Number == 1062 {
			return nil, errors.New(UNIQUE_CONSTRAINT_ERROR)
		}
		return nil, errors.New(INTERNAL_ERROR)
	}
	return item, nil
}
func (s RdbCrudResourceManager) Delete(id DataInterface) error {
	item, err := s.Get(id)
	if err != nil {
		return err
	}
	s.Rdb.Delete(item)
	return nil
}

func (s RdbCrudResourceManager) List(parameters DataInterface) (DataInterface, error) {

	items := reflect.New(reflect.SliceOf(reflect.TypeOf(s.GetModel()))).Interface()
	result, err := s.QueryBuilder.ListQuery(parameters)
	if err != nil {
		return nil, err
	}
	result = result.Find(items)
	if result.Error != nil {
		return nil, result.Error
	}

	return items, nil
}

func NewRdbCrudResourceManager(
	db *gorm.DB,
	logger logger.LoggerInterface,
	model model.BaseModelInterface,
	query_builder PaginatedRdbListQueryBuilderInterface,
) *RdbCrudResourceManager {
	resource_manager := NewResourceManager(logger, model)
	return &RdbCrudResourceManager{
		ResourceManager: resource_manager,
		Rdb:             db,
		Model:           model,
		QueryBuilder:    query_builder,
	}

}
