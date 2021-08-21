package googly

import (
	"github.com/sarulabs/di/v2"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/mock"
	"github.com/theNullP0inter/googly/ingress"
)

// MockApp is a mock for App
type MockApp struct {
	mock.Mock
}

func (a *MockApp) Build(builder *di.Builder) {
	a.Called()
}

// MockGoogly is a mock for Googly
type MockGoogly struct {
	mock.Mock
	GooglyRunner
	InstalledApps []App
}

func (g *MockGoogly) Build() *di.Builder {
	g.Called()
	builder, _ := di.NewBuilder()

	return builder

}

func (g *MockGoogly) RegisterIngressPoints(rootCmd *cobra.Command, cnt di.Container) {
	g.Called()
}

// MockGooglyRunner is a mock  implementation for GooglyRunner
type MockGooglyRunner struct {
	mock.Mock
}

func (r *MockGooglyRunner) Inject(builder *di.Builder) {
	r.Called()
}

func (r *MockGooglyRunner) GetIngressPoints(di.Container) []ingress.Ingress {
	r.Called()
	return []ingress.Ingress{}
}
