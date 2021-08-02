package main

import (
	"database/sql"

	mysqldb "github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/sarulabs/di/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	gogeta "github.com/theNullP0inter/account-management"
	"github.com/theNullP0inter/account-management/app"
	"github.com/theNullP0inter/account-management/command"
	"github.com/theNullP0inter/account-management/example/rdb_crud/accounts"
	"github.com/theNullP0inter/account-management/example/rdb_crud/consts"
	"github.com/theNullP0inter/account-management/logger"
	"github.com/theNullP0inter/account-management/rdb"
	mysql "gorm.io/driver/mysql"
)

var INSTALLED_APPS = []app.AppInterface{
	&accounts.AccountsApp{},
}

type MainAppRunner struct{}

func (a MainAppRunner) Inject(builder *di.Builder) {
	// builder.Add(di.Def{
	// 	Name: SentryClient,
	// 	Build: func(ctn di.Container) (interface{}, error) {
	// 		return logger.NewSentryClient(viper.GetString("SENTRY_DSN")), nil
	// 	},
	// })

	builder.Add(di.Def{
		Name: consts.Logger,
		Build: func(ctn di.Container) (interface{}, error) {
			l := logger.NewLogger()
			// logger.AddSentryHookToLogrus(l.(*logrus.Logger), viper.GetString("SENTRY_DSN"), viper.GetInt("SENTRY_TIMEOUT"))
			return l, nil
		},
	})

	builder.Add(di.Def{
		Name: consts.Rdb,
		Build: func(ctn di.Container) (interface{}, error) {
			dbUrl := viper.GetString("RDB_URL")
			db := rdb.NewDb(mysql.Open(dbUrl))

			sqlDB, err := db.DB()

			if err != nil {
				panic("failed to configure the database")
			}

			sqlDB.SetMaxIdleConns(viper.GetInt("RDB_MAX_CONNECTIONS"))
			sqlDB.SetMaxOpenConns(viper.GetInt("RDB_MAX_CONNECTIONS"))

			return db, nil
		},
	})
}

func (a MainAppRunner) RegisterCommands(cmd *cobra.Command, cnt di.Container) {
	serve_http := command.NewGinServerCommand(
		&command.CommandConfig{
			Name:  "serve_http",
			Short: "serves http",
		},
		cnt,
		8080,
		RouteSetup,
	)
	cmd.AddCommand(serve_http)

	// Migrations
	db, err := sql.Open("mysql", viper.GetString("RDB_URL"))
	if err != nil {
		panic(err)
	}
	driver, err := mysqldb.WithInstance(db, &mysqldb.Config{})
	if err != nil {
		panic(err)
	}

	migrate_cmd := command.NewMigrateCommand(
		&command.CommandConfig{
			Name:  "migrate",
			Short: "DB Migrator",
		},
		"file:///migrations",
		"mysql", driver,
	)

	cmd.AddCommand(migrate_cmd)
}

func main() {
	g := &gogeta.Gogeta{
		GogetaRunnerInterface: &MainAppRunner{},
		InstalledApps:         INSTALLED_APPS,
	}
	// client := gogeta.Container.Get(SentryClient).(*sentry.Client)
	// if client != nil {
	// 	func() {
	// 		gogeta.Run(g)
	// 	}()
	// } else {
	gogeta.Run(g)
	// }
}
