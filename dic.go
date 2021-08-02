package gogeta

import (
	"github.com/sarulabs/di/v2"
)

var Builder *di.Builder
var Container di.Container

func InitContainer(g *Gogeta) di.Container {
	builder := InitBuilder(g)
	Container = builder.Build()
	return Container
}

func InitBuilder(g *Gogeta) *di.Builder {
	Builder, _ = di.NewBuilder()
	g.Inject(Builder)
	for _, app := range g.InstalledApps {
		app.Build(Builder)
	}
	return Builder
}
