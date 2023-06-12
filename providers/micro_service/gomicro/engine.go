package gomicro

import (
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/providers/http"
	"github.com/rogeecn/atom/providers/micro_service"
	"github.com/rogeecn/atom/utils/opt"
	"go-micro.dev/v4"
	goMicro "go-micro.dev/v4"
)

func DefaultProvider() container.ProviderContainer {
	return container.ProviderContainer{
		Provider: Provide,
		Options: []opt.Option{
			opt.Prefix(http.DefaultPrefix),
		},
	}
}

func Provide(opts ...opt.Option) error {
	o := opt.New(opts...)
	var config micro_service.Config
	if err := o.UnmarshalConfig(&config); err != nil {
		return err
	}

	return container.Container.Provide(func() (micro_service.Service, error) {
		service := &Service{
			conf:   &config,
			Engine: goMicro.NewService(),
		}
		return service, nil
	}, o.DiOptions()...)
}

type Service struct {
	conf   *micro_service.Config
	Engine goMicro.Service
}

func (s *Service) Serve() error {
	return s.Engine.Run()
}

func (s *Service) GetEngine() any {
	return s.Engine
}

func (s *Service) Init(f func()) {
	s.Engine.Init(
		micro.Name("abc"),
	)
}
