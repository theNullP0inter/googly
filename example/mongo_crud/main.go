package main

import (
	"github.com/sarulabs/di/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/theNullP0inter/googly"
	"github.com/theNullP0inter/googly/app"
	"github.com/theNullP0inter/googly/command"
	googly_db "github.com/theNullP0inter/googly/db"
	"github.com/theNullP0inter/googly/example/mongo_crud/accounts"
	"github.com/theNullP0inter/googly/example/mongo_crud/consts"
	"github.com/theNullP0inter/googly/ingress"
	"github.com/theNullP0inter/googly/logger"
)

var INSTALLED_APPS = []app.AppInterface{
	&accounts.AccountsApp{},
}

type MainAppRunner struct{}

func (a MainAppRunner) Inject(builder *di.Builder) {
	builder.Add(di.Def{
		Name: consts.Logger,
		Build: func(ctn di.Container) (interface{}, error) {
			l := logger.NewLogger()
			return l, nil
		},
	})

	builder.Add(di.Def{
		Name: consts.MongoDb,
		Build: func(ctn di.Container) (interface{}, error) {
			dbUrl := viper.GetString("MONGO_URL")
			dbName := viper.GetString("MONGO_DB_NAME")
			db := googly_db.NewMongoDatabase(dbUrl, dbName)
			return db, nil
		},
	})
}

func (a MainAppRunner) RegisterCommands(cmd *cobra.Command, cnt di.Container) {
	serve_http := ingress.NewGinServerCommand(
		&command.CommandConfig{
			Name:  "serve_http",
			Short: "serves http",
		},
		cnt,
		8080,
		NewMainIngress(),
	)
	cmd.AddCommand(serve_http)

}

func main() {
	g := &googly.Googly{
		GooglyRunnerInterface: &MainAppRunner{},
		InstalledApps:         INSTALLED_APPS,
	}

	googly.Run(g)

}
