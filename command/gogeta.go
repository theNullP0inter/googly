package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

type CommandConfig struct {
	Name  string
	Short string
}

var GogetaCmd = &cobra.Command{
	Use:   "gogeta",
	Short: "gogeta",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Gogeta: Initialized")
		// This is just a root command. Nothing to run
	},
}
