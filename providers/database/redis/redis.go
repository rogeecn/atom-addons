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

	return container.Container.Provide(func() (*RedisClientWrapper, redis.Cmdable, error) {
		var client redis.Cmdable
		if conf.IsClusterMode() {
			client = redis.NewClusterClient(conf.ToRedisClusterOptions())
		} else {
			client = redis.NewClient(conf.ToRedisOptions())
		}

		_, err := client.Ping(context.Background()).Result()
		if err != nil {
			return nil, nil, errors.Wrap(err, "failed to ping")
		}

		redisWrapper := &RedisClientWrapper{Client: client}
		container.AddCloseAble(func() {
			redisWrapper.Close()
			switch client.(type) {
			case *redis.Client:
				client.(*redis.Client).Close()
			case *redis.ClusterClient:
				client.(*redis.ClusterClient).Close()
			}
		})
		return redisWrapper, client, nil
	}, o.DiOptions()...)
}

type RedisClientWrapper struct {
	Client redis.Cmdable
}

func (r *RedisClientWrapper) Close() error {
	if r.Client == nil {
		return nil
	}
	switch r.Client.(type) {
	case *redis.Client:
		r.Client.(*redis.Client).Close()
	case *redis.ClusterClient:
		r.Client.(*redis.ClusterClient).Close()
	}
	return nil
}

func (r *RedisClientWrapper) MakeRedisClient() interface{} {
	return r.Client
}
