package ingress

import (
	"github.com/spf13/cobra"
)

type IngressInterface interface {
	GetEntryCommand() *cobra.Command
}

type Ingress struct {
	IngressInterface
	EntryCommand *cobra.Command
}

func (i *Ingress) GetEntryCommand() *cobra.Command {
	return i.EntryCommand
}

func NewIngress(cmd *cobra.Command) *Ingress {
	return &Ingress{
		EntryCommand: cmd,
	}
}
