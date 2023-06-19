package redis

import (
	"context"

	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/utils/opt"
)

func Provide(opts ...opt.Option) error {
	o := opt.New(opts...)
	var conf Config
	if err := o.UnmarshalConfig(&conf); err != nil {
		return err
	}

	return container.Container.Provide(func() (*RedisClientWrapper, *redis.Client, error) {
		client := redis.NewClient(conf.ToRedisOptions())

		_, err := client.Ping(context.Background()).Result()
		if err != nil {
			return nil, nil, errors.Wrap(err, "failed to ping")
		}

		redis := &RedisClientWrapper{Client: client}
		container.AddCloseAble(func() {
			redis.Close()
		})
		return redis, redis.Client, nil
	}, o.DiOptions()...)
}

type RedisClientWrapper struct {
	Client *redis.Client
}

func (r *RedisClientWrapper) Close() error {
	if r.Client == nil {
		return nil
	}
	return r.Client.Close()
}

func (r *RedisClientWrapper) MakeRedisClient() interface{} {
	return r.Client
}
