package gogeta

import (
	"fmt"
	"os"

	"github.com/sarulabs/di/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/theNullP0inter/account-management/app"
	"github.com/theNullP0inter/account-management/command"
)

type GogetaInjectorInterface interface {
	Inject(builder *di.Builder)
}

type GogetaCommanderInterface interface {
	RegisterCommands(cmd *cobra.Command, cnt di.Container)
}

type GogetaRunnerInterface interface {
	GogetaInjectorInterface
	GogetaCommanderInterface
}

type GogetaRunner struct {
	GogetaRunnerInterface
}

type Gogeta struct {
	GogetaRunnerInterface
	InstalledApps []app.AppInterface
}

func Run(g *Gogeta) {

	viper.AutomaticEnv()
	cnt := InitContainer(g)
	root_cmd := command.GogetaCmd
	g.RegisterCommands(root_cmd, cnt)
	if err := root_cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
