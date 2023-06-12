package http

import (
	"github.com/rogeecn/atom-addons/providers/http/gin"
	"github.com/rogeecn/atom-addons/providers/log"
	"github.com/rogeecn/atom/container"
)

func DefaultHTTP(providers ...container.ProviderContainer) container.Providers {
	return append(container.Providers{
		log.DefaultProvider(),
		gin.DefaultProvider(),
	}, providers...)
}
