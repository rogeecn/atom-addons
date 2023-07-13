package grpcs

import (
	"net"

	"github.com/pkg/errors"
	"github.com/rogeecn/atom-addons/providers/log"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/utils/opt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

func DefaultProvider() container.ProviderContainer {
	return container.ProviderContainer{
		Provider: Provide,
		Options: []opt.Option{
			opt.Prefix(DefaultPrefix),
		},
	}
}

// 给注入启动调用使用
type ServerService interface {
	Name() string
	Register(*grpc.Server)
}

type Grpc struct {
	server *grpc.Server
	config *Config
}

func Provide(opts ...opt.Option) error {
	o := opt.New(opts...)
	var config Config
	if err := o.UnmarshalConfig(&config); err != nil {
		return err
	}

	return container.Container.Provide(func() (*Grpc, error) {
		serverOptions := []grpc.ServerOption{}

		// tls
		if config.Tls != nil {
			tlsConfig, err := credentials.NewServerTLSFromFile(config.Tls.Cert, config.Tls.Key)
			if err != nil {
				return nil, errors.Wrap(err, "Failed to create tls credential")
			}
			serverOptions = append(serverOptions, grpc.Creds(tlsConfig))
		}

		svc := &Grpc{
			server: grpc.NewServer(serverOptions...),
			config: &config,
		}
		container.AddCloseAble(svc.Close)
		return svc, nil
	}, o.DiOptions()...)
}

func (g *Grpc) Serve() error {
	ld, err := net.Listen("tcp", g.config.Address())
	if err != nil {
		return errors.Wrapf(err, "bind address failed: %s", g.config.Address())
	}

	log.Infof("grpc server listen on %s", g.config.Address())
	reflection.Register(g.server)

	return g.server.Serve(ld)
}

func (g *Grpc) RegisterService(name string, f func(*grpc.Server)) {
	log.Debug("register service:", name)
	f(g.server)
}

func (g *Grpc) GracefulStop() {
	g.server.GracefulStop()
}

func (g *Grpc) Stop() {
	g.server.Stop()
}

func (g *Grpc) Close() {
	g.GracefulStop()
}
