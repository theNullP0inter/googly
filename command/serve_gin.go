package command

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sarulabs/di/v2"
	"github.com/spf13/cobra"
)

type GinRouterSetup func(*gin.Engine, di.Container) *gin.Engine

func NewGinServerCommand(config *CommandConfig, cnt di.Container, port int, setup GinRouterSetup) *cobra.Command {

	var ginServerCmd = &cobra.Command{
		Use:   config.Name,
		Short: config.Short,
		Run: func(cmd *cobra.Command, args []string) {
			router := gin.New()
			router = setup(router, cnt)
			router.Run(":" + fmt.Sprint(port))
		},
	}

	return ginServerCmd

}
