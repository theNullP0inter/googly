package googly

import (
	"github.com/sarulabs/di/v2"
)

var Builder *di.Builder
var Container di.Container

func InitContainer(g *Googly) di.Container {
	builder := InitBuilder(g)
	Container = builder.Build()
	return Container
}

func InitBuilder(g *Googly) *di.Builder {
	Builder, _ = di.NewBuilder()
	g.Inject(Builder)
	for _, app := range g.InstalledApps {
		app.Build(Builder)
	}
	return Builder
}
