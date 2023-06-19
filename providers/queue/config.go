package queue

import (
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/utils/opt"
)

const DefaultPrefix = "Queue"

func DefaultProvider() container.ProviderContainer {
	return container.ProviderContainer{
		Provider: Provide,
		Options: []opt.Option{
			opt.Prefix(DefaultPrefix),
		},
	}
}

type Config struct {
	Concurrency int
}
