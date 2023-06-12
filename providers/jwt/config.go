package jwt

import (
	"time"

	"github.com/rogeecn/atom-addons/providers/log"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/utils/opt"
)

const DefaultPrefix = "JWT"

func DefaultProvider() container.ProviderContainer {
	return container.ProviderContainer{
		Provider: Provide,
		Options: []opt.Option{
			opt.Prefix(DefaultPrefix),
		},
	}
}

type Config struct {
	SigningKey  string // jwt签名
	ExpiresTime string // 过期时间
	Issuer      string // 签发者
}

func (c *Config) ExpiresTimeDuration() time.Duration {
	d, err := time.ParseDuration(c.ExpiresTime)
	if err != nil {
		log.Fatal(err)
	}
	return d
}
