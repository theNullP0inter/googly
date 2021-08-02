package ingress

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sarulabs/di/v2"
	"github.com/spf13/cobra"
	"github.com/theNullP0inter/googly/command"
)

type GinIngressInterface interface {
	Connect(di.Container) *gin.Engine
}

func NewGinServerCommand(config *command.CommandConfig, cnt di.Container, port int, ingress GinIngressInterface) *cobra.Command {

	var ginServerCmd = &cobra.Command{
		Use:   config.Name,
		Short: config.Short,
		Run: func(cmd *cobra.Command, args []string) {
			router := ingress.Connect(cnt)
			router.Run(":" + fmt.Sprint(port))
		},
	}

	return ginServerCmd

}
