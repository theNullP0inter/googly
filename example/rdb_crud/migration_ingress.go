package main

import (
	"database/sql"

	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/sarulabs/di/v2"
	"github.com/spf13/viper"
	"github.com/theNullP0inter/googly/ingress"
)

func NewMainMigrationIngress(cnt di.Container) *ingress.MigrationIngress {
	// Migrations
	db, err := sql.Open("mysql", viper.GetString("RDB_URL"))
	if err != nil {
		panic(err)
	}
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		panic(err)
	}
	return ingress.NewMigrationIngress(
		"migrate",
		"/migrations",
		driver,
	)
}
