package command

import (
	"github.com/spf13/cobra"
	"github.com/theNullP0inter/account-management/dic"
	"github.com/theNullP0inter/account-management/route"
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run server",
	Run: func(cmd *cobra.Command, args []string) {
		// Builder is already generated in main.go. We're using that via pointer
		router := route.Setup(dic.Builder)
		router.Run(":8080")
	},
}
