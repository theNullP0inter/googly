package ingress

import (
	"fmt"
	"net"

	"github.com/sarulabs/di/v2"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

// GrpcIngressConnector acts connector between GrpcIngress and your application
//
// GrpcIngressInterface implements this
type GrpcIngressConnector interface {
	Connect(*grpc.Server)
}

// GrpcIngressInterface is a Grpc binding for IngressInterface
//
// Connects to your application via GrpcIngressConnector
type GrpcIngressInterface interface {
	IngressInterface
	GrpcIngressConnector
}

// GrpcIngress is a basic ingress implementation for Grpc
type GrpcIngress struct {
	*Ingress
	Connector GrpcIngressConnector
}

// NewGrpcServerCommand creates a new cobra command that runs a Grpc server.
//
// This is invoked while creating a new GrpcIngress
func NewGrpcServerCommand(config *CommandConfig, cnt di.Container, port int, connector GrpcIngressConnector) *cobra.Command {

	var grpcServerCmd = &cobra.Command{
		Use:   config.Name,
		Short: config.Short,
		Run: func(cmd *cobra.Command, args []string) {
			grpc_server := grpc.NewServer()

			connector.Connect(grpc_server)

			fmt.Println("GRPC Ingress Connected")

			lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

			if err != nil {
				panic(fmt.Sprintf("failed to listen at %d: %v", port, err))
			}

			if err := grpc_server.Serve((lis)); err != nil {
				panic(fmt.Sprintf("failed to Serve at %d: %v", port, err))
			}

		},
	}

	return grpcServerCmd

}

// NewGrpcIngress creates a new GrpcIngress
func NewGrpcIngress(name string, cnt di.Container, port int, connector GrpcIngressConnector) *GrpcIngress {

	cmd := NewGrpcServerCommand(
		&CommandConfig{
			Name:  name,
			Short: fmt.Sprintf("%s Ingress", name),
		},
		cnt,
		port,
		connector,
	)
	ingress := NewIngress(cmd)
	return &GrpcIngress{ingress, connector}
}
