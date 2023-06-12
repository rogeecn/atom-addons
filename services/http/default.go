package http

import (
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/providers/http/gin"
	"github.com/rogeecn/atom/providers/log"
)

func DefaultHTTP(providers ...container.ProviderContainer) container.Providers {
	return append(container.Providers{
		log.DefaultProvider(),
		gin.DefaultProvider(),
	}, providers...)
}
