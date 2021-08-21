package googly

import (
	"testing"

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

// TestRun tests the Run() function
func TestRun(t *testing.T) {

	mock_googly := &MockGoogly{}

	mock_googly.On("Build").Return()
	mock_googly.On("RegisterIngressPoints").Return(mock.AnythingOfType("cobra.Command"), mock.AnythingOfType("di.Container"))

	Run(mock_googly)

	mock_googly.AssertExpectations(t)

}

// TestGoogly Creates a new &Googly{} with a MockApp and test its internal methods
func TestGoogly(t *testing.T) {
	mock_app := new(MockApp)
	mock_runner := new(MockGooglyRunner)
	base_googly := &Googly{
		GooglyRunner: mock_runner,
		InstalledApps: []App{
			mock_app,
		}}

	// Test Googly.Build

	mock_app.On("Build", mock.Anything).Return()
	mock_runner.On("Inject").Return()
	builder := base_googly.Build()
	mock_app.AssertExpectations(t)
	mock_runner.AssertExpectations(t)

	// Test Googly.RegisterIngressPoints
	mock_runner.On("GetIngressPoints", mock.Anything).Return()

	base_googly.RegisterIngressPoints(new(cobra.Command), builder.Build())
	mock_app.AssertExpectations(t)
}
