package swagger

import (
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/utils/opt"
)

const DefaultPrefix = "Swagger"

type Config struct {
	BaseRoute   string
	Version     string
	Host        string
	BasePath    string
	Title       string
	Description string
}

func DefaultProvider() container.ProviderContainer {
	return container.ProviderContainer{
		Provider: Provide,
		Options: []opt.Option{
			opt.Prefix(DefaultPrefix),
		},
	}
}
