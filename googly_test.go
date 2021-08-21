package googly

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/mock"
)

// TestRun tests the Run() function
func TestRun(t *testing.T) {

	mockGoogly := new(MockGoogly)

	mockGoogly.On("Build").Return()
	mockGoogly.On("RegisterIngressPoints").Return(mock.AnythingOfType("cobra.Command"), mock.AnythingOfType("di.Container"))

	Run(mockGoogly)

	mockGoogly.AssertExpectations(t)

}

// TestGoogly Creates a new &Googly{} with a MockApp and test its internal methods
func TestGoogly(t *testing.T) {
	mockApp := new(MockApp)
	mockRunner := new(MockGooglyRunner)
	baseGoogly := &Googly{
		GooglyRunner: mockRunner,
		InstalledApps: []App{
			mockApp,
		}}

	// Test Googly.Build

	mockApp.On("Build", mock.Anything).Return()
	mockRunner.On("Inject").Return()
	builder := baseGoogly.Build()
	mockApp.AssertExpectations(t)
	mockRunner.AssertExpectations(t)

	// Test Googly.RegisterIngressPoints
	mockRunner.On("GetIngressPoints", mock.Anything).Return()

	baseGoogly.RegisterIngressPoints(new(cobra.Command), builder.Build())
	mockApp.AssertExpectations(t)
}
