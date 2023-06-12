package http

import (
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/contracts"
	"github.com/rogeecn/atom/providers/http"
	"go.uber.org/dig"
)

type Http struct {
	dig.In

	Service  http.Service
	Initials []contracts.Initial `group:"initials"`
	Routes   []http.Route        `group:"routes"`
}

func ServeHttp() error {
	defer container.Close()

	return container.Container.Invoke(func(http Http) error {
		return http.Service.Serve()
	})
}
