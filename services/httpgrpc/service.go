package httpgrpc

import (
	"github.com/rogeecn/atom-addons/providers/grpcs"
	"github.com/rogeecn/atom-addons/providers/http"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/contracts"
	"github.com/spf13/cobra"
	"go.uber.org/dig"
	"golang.org/x/sync/errgroup"
)

type Service struct {
	dig.In

	Http     http.Service
	Grpc     *grpcs.Grpc
	Initials []contracts.Initial   `group:"initials"`
	Handlers []grpcs.ServerService `group:"grpc_server_services"`
	Routes   []http.Route          `group:"routes"`
}

func Serve() error {
	defer container.Close()

	return container.Container.Invoke(func(svc Service) error {
		for _, hdl := range svc.Handlers {
			svc.Grpc.RegisterService(hdl.Name(), hdl.Register)
		}

		var eg errgroup.Group
		eg.Go(svc.Http.Serve)
		eg.Go(svc.Grpc.Serve)
		return eg.Wait()
	})
}

func ServeRunE(cmd *cobra.Command, args []string) error {
	return Serve()
}
