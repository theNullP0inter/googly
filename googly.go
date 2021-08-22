package googly

import (
	"fmt"
	"os"

	"github.com/sarulabs/di/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/theNullP0inter/googly/ingress"
)

// GooglyRunner Connects googly to your application.
//
// Dependencies should be injected using Inject() Method
// Ingress can be registered using GetIngressPoints()
type GooglyRunner interface {
	Inject(builder *di.Builder)
	GetIngressPoints(di.Container) []ingress.Ingress
}

// GooglyInterface is implemented by Googly
type GooglyInterface interface {
	GooglyRunner
	Build() *di.Builder
	RegisterIngressPoints(rootCmd *cobra.Command, cnt di.Container)
}

// App is any thing that Builds. build happens through inject its components into googlys builder
//
// Build(*di.Builder) is where you should inject dependencies for the sub-app
type App interface {
	Build(*di.Builder)
}

// Googly manages your application through dependendency injection
// and runs it using Ingress
//
// GooglyRunner can be used to register Global dependencies
// InstalledApps maintains a list of all your sub apps.
type Googly struct {
	GooglyRunner
	InstalledApps []App
}

// Build builds the application and sub-apps
//
// This is invoked by Run
func (g *Googly) Build() *di.Builder {
	builder, _ := di.NewBuilder()
	g.Inject(builder)

	for _, app := range g.InstalledApps {
		app.Build(builder)
	}
	return builder
}

// RegisterIngressPoints will register all the ingress points returned in GetIngressPoints to GooglyCmd
//
// This is invoked by Run
func (g *Googly) RegisterIngressPoints(rootCmd *cobra.Command, cnt di.Container) {
	for _, ig := range g.GetIngressPoints(cnt) {
		rootCmd.AddCommand(ig.GetEntryCommand())
	}
}

// Run runs your application in the following order.
//
// Reads env
// Injects all dependencties
// Builds the sub-apps
// Registers ingress points
// Runs root command
func Run(g GooglyInterface) {

	// setup env
	viper.AutomaticEnv()

	// Build sub-apps
	builder := g.Build()
	container := builder.Build()

	// Register Ingress
	rootCmd := ingress.GooglyCmd
	g.RegisterIngressPoints(rootCmd, container)

	// Run
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
