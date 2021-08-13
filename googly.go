package googly

import (
	"fmt"
	"os"

	"github.com/sarulabs/di/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/theNullP0inter/googly/command"
	"github.com/theNullP0inter/googly/ingress"
)

type GooglyInterface interface {
	Inject(builder *di.Builder)
	GetIngressPoints(di.Container) []ingress.IngressInterface
}

type Googly struct {
	GooglyInterface
	InstalledApps []AppInterface
}

func (g *Googly) RegisterIngressPoints(root_cmd *cobra.Command, cnt di.Container) {
	for _, ig := range g.GetIngressPoints(cnt) {
		root_cmd.AddCommand(ig.GetEntryCommand())
	}
}

func Run(g *Googly) {

	viper.AutomaticEnv()
	cnt := InitContainer(g)
	root_cmd := command.GooglyCmd
	g.RegisterIngressPoints(root_cmd, cnt)
	if err := root_cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
