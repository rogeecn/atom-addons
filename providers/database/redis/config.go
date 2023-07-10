package redis

import (
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/utils/opt"
)

const DefaultPrefix = "Redis"

func DefaultProvider() container.ProviderContainer {
	return container.ProviderContainer{
		Provider: Provide,
		Options: []opt.Option{
			opt.Prefix(DefaultPrefix),
		},
	}
}

type Config struct {
	ClusterName string
	Host        string
	Port        uint
	Password    string
	DB          uint
}

func (c *Config) ToRedisOptions() *redis.Options {
	return &redis.Options{
		Addr:     fmt.Sprintf("%s:%d", c.Host, c.Port),
		Password: c.Password,
		DB:       int(c.DB),
	}
}

func (c *Config) ToRedisClusterOptions() *redis.ClusterOptions {
	return &redis.ClusterOptions{
		ClientName: c.ClusterName,
		Addrs:      []string{fmt.Sprintf("%s:%d", c.Host, c.Port)},
		Password:   c.Password,
	}
}
