package ingress

import (
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/mock"
)

// MockIngress is a mock implementation for Ingress
type MockIngress struct {
	mock.Mock
}

func (i *MockIngress) GetEntryCommand() *cobra.Command {
	i.Called()
	return new(cobra.Command)
}
