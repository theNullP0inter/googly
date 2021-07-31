package account

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AccountService struct {
	Logger *logrus.Logger
	Rdb    *gorm.DB
}

func NewAccountService(rdb *gorm.DB, logger *logrus.Logger) *AccountService {
	return &AccountService{
		Rdb:    rdb,
		Logger: logger,
	}
}
