package queue

import (
	"context"
	"encoding/json"

	"github.com/hibiken/asynq"
	"github.com/rogeecn/atom-addons/providers/database/redis"
	"github.com/rogeecn/atom-addons/providers/log"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/utils/opt"
)

func Provide(opts ...opt.Option) error {
	o := opt.New(opts...)
	var config Config
	if err := o.UnmarshalConfig(&config); err != nil {
		return err
	}
	return container.Container.Provide(func(logger *log.Logger, redis *redis.RedisClientWrapper) (*Queue, error) {
		return &Queue{
			Server: asynq.NewServer(
				redis,
				asynq.Config{
					// Specify how many concurrent workers to use
					Concurrency: config.Concurrency,
					// Optionally specify multiple queues with different priority.
					Queues: map[string]int{
						"critical": 6,
						"default":  3,
						"low":      1,
					},
					Logger:   logger.Logger,
					LogLevel: asynq.DebugLevel,
					// See the godoc for other configuration options
				},
			),
			Mux:    asynq.NewServeMux(),
			Client: asynq.NewClient(redis),
		}, nil
	}, o.DiOptions()...)
}

type QueueHandler interface {
	Type() string
	Options() []asynq.Option
	EnqueueContext(ctx context.Context, payload interface{}) (*asynq.TaskInfo, error)
	ProcessTask(context.Context, *asynq.Task) error
}

type Queue struct {
	Client *asynq.Client
	Server *asynq.Server
	Mux    *asynq.ServeMux
}

func (q *Queue) Serve() error {
	return q.Server.Run(q.Mux)
}

func (q *Queue) Handle(handlers ...QueueHandler) {
	for _, hdl := range handlers {
		q.Mux.Handle(hdl.Type(), hdl)
	}
}

func (q *Queue) Enqueue(hdl QueueHandler, payload interface{}) (*asynq.TaskInfo, error) {
	return q.EnqueueContext(context.Background(), hdl, payload)
}

func (q *Queue) EnqueueContext(ctx context.Context, hdl QueueHandler, payload interface{}) (*asynq.TaskInfo, error) {
	d, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	t := asynq.NewTask(hdl.Type(), d)
	return q.Client.EnqueueContext(ctx, t, hdl.Options()...)
}
