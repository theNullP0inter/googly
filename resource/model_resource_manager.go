package resource

import (
	"errors"
	"reflect"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/copier"
	"github.com/sirupsen/logrus"
	"github.com/theNullP0inter/account-management/model"
	"gorm.io/gorm"
)

type ModelResourceManagerInterface interface {
	ModelCrudInterface
	ResourceManagerInterface
	ListQueryBuilderInterface
	GetModel() DataInterface
}

type ModelResourceManager struct {
	*ResourceManager
	ListQueryBuilderInterface
	Rdb   *gorm.DB
	Model model.BaseModelInterface
}

func (s *ModelResourceManager) Create(m DataInterface) (DataInterface, error) {
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

func (s ModelResourceManager) GetResource() ResourceInterface {
	return s.Model
}

func (s ModelResourceManager) GetModel() DataInterface {
	return s.Model
}
func (s ModelResourceManager) Get(id model.BinID) (DataInterface, error) {
	item := reflect.New(reflect.TypeOf(s.GetModel())).Interface()
	b_id, _ := id.MarshalBinary()
	err := s.Rdb.Where("id = ?", b_id).First(item).Error
	return item, err
}

func (s ModelResourceManager) Update(item DataInterface) (DataInterface, error) {
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
func (s ModelResourceManager) Delete(id model.BinID) error {
	item, err := s.Get(id)
	if err != nil {
		return err
	}
	s.Rdb.Delete(item)
	return nil
}

func (s ModelResourceManager) List(parameters DataInterface) (DataInterface, error) {

	items := reflect.New(reflect.SliceOf(reflect.TypeOf(s.GetModel()))).Interface()
	result, err := s.ListQuery(parameters)
	if err != nil {
		return nil, err
	}
	result = result.Find(items)
	if result.Error != nil {
		return nil, result.Error
	}

	return items, nil
}

func (s *ModelResourceManager) ListQuery(parameters ListParametersInterface) (*gorm.DB, error) {
	return s.PaginationQuery(parameters), nil
}

func NewModelResourceManager(
	db *gorm.DB,
	logger *logrus.Logger,
	model model.BaseModelInterface,
	query_builder ListQueryBuilderInterface,
) *ModelResourceManager {
	new_resource := NewResourceManager(logger, model.(ResourceInterface))
	return &ModelResourceManager{
		new_resource,
		query_builder,
		db,
		model,
	}
}
