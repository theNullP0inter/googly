package ingress

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sarulabs/di/v2"
	"github.com/spf13/cobra"
	"github.com/theNullP0inter/googly/command"
)

type GinIngressConnector interface {
	Connect(di.Container) *gin.Engine
}

type GinIngress struct {
	*Ingress
	Connector GinIngressConnector
}

func NewGinServerCommand(config *command.CommandConfig, cnt di.Container, port int, connector GinIngressConnector) *cobra.Command {

	var ginServerCmd = &cobra.Command{
		Use:   config.Name,
		Short: config.Short,
		Run: func(cmd *cobra.Command, args []string) {
			router := connector.Connect(cnt)
			router.Run(":" + fmt.Sprint(port))
		},
	}

	return ginServerCmd

}

func NewGinIngress(name string, cnt di.Container, port int, connector GinIngressConnector) *GinIngress {
	cmd := NewGinServerCommand(
		&command.CommandConfig{
			Name:  name,
			Short: fmt.Sprintf("%s Ingress", name),
		},
		cnt,
		port,
		connector,
	)
	ingress := NewIngress(cmd)
	return &GinIngress{ingress, connector}
}
