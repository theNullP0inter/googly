package ingress

import (
	"fmt"
	"net"

	"github.com/sarulabs/di/v2"
	"github.com/spf13/cobra"
	"github.com/theNullP0inter/googly/command"
	"google.golang.org/grpc"
)

type GrpcProtoSetup func(*grpc.Server)

type GrpcIngressInterface interface {
	Connect(*grpc.Server)
}

func NewGrpcServerCommand(config *command.CommandConfig, cnt di.Container, port int, ingress GrpcIngressInterface) *cobra.Command {

	var ginServerCmd = &cobra.Command{
		Use:   config.Name,
		Short: config.Short,
		Run: func(cmd *cobra.Command, args []string) {
			grpc_server := grpc.NewServer()

			ingress.Connect(grpc_server)

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

	return ginServerCmd

}
