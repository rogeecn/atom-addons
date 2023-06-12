package single_flight

import (
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/utils/opt"
	"golang.org/x/sync/singleflight"
)

func DefaultProvider() container.ProviderContainer {
	return container.ProviderContainer{
		Provider: Provide,
		Options:  []opt.Option{},
	}
}

func Provide(opts ...opt.Option) error {
	o := opt.New(opts...)
	return container.Container.Provide(func() (*singleflight.Group, error) {
		return &singleflight.Group{}, nil
	}, o.DiOptions()...)
}
