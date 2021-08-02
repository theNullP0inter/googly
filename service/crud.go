package service

import (
	"github.com/theNullP0inter/googly/errors"
)

type CrudInterface interface {
	GetItem(id DataInterface) (DataInterface, *errors.GooglyError)
	GetList(req DataInterface) (DataInterface, *errors.GooglyError)
	Create(req DataInterface) (DataInterface, *errors.GooglyError)
	Update(item DataInterface) (DataInterface, *errors.GooglyError)
	Delete(id DataInterface) *errors.GooglyError
}
