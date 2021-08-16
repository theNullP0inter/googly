package ingress

import (
	"fmt"

	"github.com/spf13/cobra"
)

// CommandConfig maintains the configuration for a sub command
//
// Name will be used to register a command as sub command to GooglyCmd
type CommandConfig struct {
	Name  string
	Short string
}

// GooglyCmd is the root command for Ingress
var GooglyCmd = &cobra.Command{
	Use:   "googly",
	Short: "googly",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Googly: Running")
		// This is just a root command. Nothing to run
	},
}
