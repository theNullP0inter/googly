package rdb

import (
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDb() *gorm.DB {
	dbUrl := viper.GetString("RDB_URL")
	db, err := gorm.Open(mysql.Open(dbUrl))
	if err != nil {
		panic("failed to connect to the database")
	}

	sqlDB, err := db.DB()

	if err != nil {
		panic("failed to configure the database")
	}

	sqlDB.SetMaxIdleConns(viper.GetInt("RDB_MAX_CONNECTIONS"))
	sqlDB.SetMaxOpenConns(viper.GetInt("RDB_MAX_CONNECTIONS"))

	return db
}
