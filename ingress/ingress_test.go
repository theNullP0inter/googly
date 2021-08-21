package ingress

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

// TestBaseIngress creates a new ingress and asserts the response from GetEntryCommand
func TestBaseIngress(t *testing.T) {
	rootCmd := new(cobra.Command)
	baseIngress := NewBaseIngress(rootCmd)

	entryCmd := baseIngress.GetEntryCommand()

	assert.Equal(t, rootCmd, entryCmd)

}
