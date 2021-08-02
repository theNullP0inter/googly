package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

type CommandConfig struct {
	Name  string
	Short string
}

var GooglyCmd = &cobra.Command{
	Use:   "googly",
	Short: "googly",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Googly: Running")
		// This is just a root command. Nothing to run
	},
}
