package ingress

import (
	"github.com/spf13/cobra"
)

// Ingress is required to be implemented by all your ingress points
//
// GetEntryCommand should return a *cobra.Command which will be added to root command
type Ingress interface {
	GetEntryCommand() *cobra.Command
}

// BaseIngress is a base implementation of Ingress
type BaseIngress struct {
	EntryCommand *cobra.Command
}

// GetEntryCommand returns command registered with BaseIngress
func (i *BaseIngress) GetEntryCommand() *cobra.Command {
	return i.EntryCommand
}

// NewBaseIngress creates a new BaseIngress
func NewBaseIngress(cmd *cobra.Command) *BaseIngress {
	return &BaseIngress{cmd}
}
