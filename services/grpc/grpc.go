package grpc

import (
	"github.com/rogeecn/atom-addons/providers/grpcs"
	"github.com/rogeecn/atom/container"
	"go.uber.org/dig"
)

type GrpcService struct {
	dig.In

	Server   *grpcs.Grpc
	Services []grpcs.ServerService `group:"grpc_server_services"`
}

func ServeGrpc() error {
	defer container.Close()

	return container.Container.Invoke(func(grpc GrpcService) error {
		for _, svc := range grpc.Services {
			grpc.Server.RegisterService(svc.Name(), svc.Register)
		}
		return grpc.Server.Serve()
	})
}
