package rdb

import (
	"gorm.io/gorm"
)

func NewDb(dialect gorm.Dialector) *gorm.DB {
	db, err := gorm.Open(dialect)
	if err != nil {
		panic("failed to connect to the database")
	}

	return db
}
