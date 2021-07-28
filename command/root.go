package command

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Short: "CLI utility",
	Long:  "works along with Keycloak",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		// This is just a root command. Nothing to run
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
