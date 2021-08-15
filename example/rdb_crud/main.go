package main

import (
	"github.com/sarulabs/di/v2"
	"github.com/spf13/viper"
	"github.com/theNullP0inter/googly"
	googly_logrus "github.com/theNullP0inter/googly/contrib/logrus"
	"github.com/theNullP0inter/googly/contrib/rdb"
	"github.com/theNullP0inter/googly/example/rdb_crud/accounts"
	"github.com/theNullP0inter/googly/example/rdb_crud/consts"
	"github.com/theNullP0inter/googly/ingress"
	mysql "gorm.io/driver/mysql"
)

var INSTALLED_APPS = []googly.App{
	&accounts.AccountsApp{},
}

type MainGooglyInterface struct{}

func (a *MainGooglyInterface) Inject(builder *di.Builder) {
	builder.Add(di.Def{
		Name: consts.Logger,
		Build: func(ctn di.Container) (interface{}, error) {
			l := googly_logrus.NewGooglyLogrusLogger()
			return l, nil
		},
	})

	builder.Add(di.Def{
		Name: consts.Rdb,
		Build: func(ctn di.Container) (interface{}, error) {
			dbUrl := viper.GetString("RDB_URL")
			db := rdb.NewRdb(mysql.Open(dbUrl))
			return db, nil
		},
	})
}

func (a *MainGooglyInterface) GetIngressPoints(cnt di.Container) []ingress.Ingress {
	return []ingress.Ingress{
		NewMainGinIngress(cnt, 8080),
		NewMainMigrationIngress(cnt),
	}

}

func main() {
	g := &googly.Googly{
		GooglyInterface: &MainGooglyInterface{},
		InstalledApps:   INSTALLED_APPS,
	}
	googly.Run(g)
}
