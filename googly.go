package googly

import (
	"fmt"
	"os"

	"github.com/sarulabs/di/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/theNullP0inter/googly/app"
	"github.com/theNullP0inter/googly/command"
)

type GooglyInjectorInterface interface {
	Inject(builder *di.Builder)
}

type GooglyCommanderInterface interface {
	RegisterCommands(cmd *cobra.Command, cnt di.Container)
}

type GooglyRunnerInterface interface {
	GooglyInjectorInterface
	GooglyCommanderInterface
}

type GooglyRunner struct {
	GooglyRunnerInterface
}

type Googly struct {
	GooglyRunnerInterface
	InstalledApps []app.AppInterface
}

func Run(g *Googly) {

	viper.AutomaticEnv()
	cnt := InitContainer(g)
	root_cmd := command.GooglyCmd
	g.RegisterCommands(root_cmd, cnt)
	if err := root_cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
