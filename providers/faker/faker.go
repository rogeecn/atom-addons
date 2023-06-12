package faker

import (
	"time"

	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/utils/opt"

	"github.com/brianvoe/gofakeit/v6"
)

func DefaultProvider() container.ProviderContainer {
	return container.ProviderContainer{
		Provider: Provide,
		Options:  []opt.Option{},
	}
}

func Provide(opts ...opt.Option) error {
	o := opt.New(opts...)
	return container.Container.Provide(func() (*gofakeit.Faker, error) {
		faker := gofakeit.New(time.Now().UnixNano())
		gofakeit.SetGlobalFaker(faker)

		return faker, nil
	}, o.DiOptions()...)
}
