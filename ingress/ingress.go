package ingress

import (
	"github.com/spf13/cobra"
)

// IngressInterface is required to be implemented by all your ingress points
//
// GetEntryCommand should return a *cobra.Command which will be added to root command
type IngressInterface interface {
	GetEntryCommand() *cobra.Command
}

// Ingress is a base implementation of IngressInterface
type Ingress struct {
	IngressInterface
	EntryCommand *cobra.Command
}

// GetEntryCommand returns command registered with Ingress
func (i *Ingress) GetEntryCommand() *cobra.Command {
	return i.EntryCommand
}

// NewIngress creates a new instance of Ingress
func NewIngress(cmd *cobra.Command) *Ingress {
	return &Ingress{
		EntryCommand: cmd,
	}
}
