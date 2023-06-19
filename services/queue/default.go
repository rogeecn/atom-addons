package queue

import (
	"github.com/rogeecn/atom-addons/providers/log"
	"github.com/rogeecn/atom-addons/providers/queue"
	"github.com/rogeecn/atom/container"
)

func Default(providers ...container.ProviderContainer) container.Providers {
	return append(container.Providers{
		log.DefaultProvider(),
		queue.DefaultProvider(),
	}, providers...)
}
