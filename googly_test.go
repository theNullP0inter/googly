package googly

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/mock"
)

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
