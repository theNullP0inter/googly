package ingress

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sarulabs/di/v2"
	"github.com/spf13/cobra"
)

// GinIngressConnector acts connector between GinIngress and your application
//
// GinIngressInterface implements this
type GinIngressConnector interface {
	Connect(di.Container) *gin.Engine
}

// GinIngressInterface is a Gin binding for IngressInterface
//
// Connects to your application via GinIngressConnector
type GinIngressInterface interface {
	IngressInterface
	GinIngressConnector
}

// GinIngress is a basic ingress implementation for github.com/gin-gonic/gin
type GinIngress struct {
	*Ingress
	Connector GinIngressConnector
}

// NewGinServerCommand creates a new cobra command that runs a gin server.
//
// This is invoked while creating a new GinIngress
func NewGinServerCommand(config *CommandConfig, cnt di.Container, port int, connector GinIngressConnector) *cobra.Command {

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

// NewGinIngress creates a new GinIngress
func NewGinIngress(name string, cnt di.Container, port int, connector GinIngressConnector) *GinIngress {
	cmd := NewGinServerCommand(
		&CommandConfig{
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
