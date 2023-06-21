package httpgrpc

import (
	"github.com/rogeecn/atom-addons/providers/grpcs"
	"github.com/rogeecn/atom-addons/providers/http/fiber"
	"github.com/rogeecn/atom-addons/providers/log"
	"github.com/rogeecn/atom/container"
)

func Default(providers ...container.ProviderContainer) container.Providers {
	return append(container.Providers{
		log.DefaultProvider(),
		fiber.DefaultProvider(),
		grpcs.DefaultProvider(),
	}, providers...)
}
