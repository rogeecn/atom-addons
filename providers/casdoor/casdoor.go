package casdoor

import (
	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/rogeecn/atom-addons/providers/cert"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/utils/opt"
)

type Casdoor struct{}

func Provide(opts ...opt.Option) error {
	o := opt.New(opts...)
	var config Config
	if err := o.UnmarshalConfig(&config); err != nil {
		return err
	}
	return container.Container.Provide(func(cert *cert.Cert) *Casdoor {
		certificate := config.Certificate
		if certificate == "" {
			certificate = cert.Cert
		}
		casdoorsdk.InitConfig(
			config.Endpoint,
			config.ClientId,
			config.ClientSecret,
			certificate,
			config.OrganizationName,
			config.ApplicationName,
		)
		return &Casdoor{}
	}, o.DiOptions()...)
}
