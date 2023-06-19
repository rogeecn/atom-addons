package queue

import (
	"context"

	"github.com/hibiken/asynq"
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
	return container.Container.Provide(func(logger *log.Logger) (*Queue, error) {
		redisClientOptions := asynq.RedisClientOpt{}
		return &Queue{
			Server: asynq.NewServer(
				redisClientOptions,
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
			Client: asynq.NewClient(redisClientOptions),
		}, nil
	}, o.DiOptions()...)
}

type QueueHandler interface {
	ProcessTask(context.Context, *asynq.Task) error
	GetType() string
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
		q.Mux.Handle(hdl.GetType(), hdl)
	}
}

func (q *Queue) Enqueue(task *asynq.Task, opts ...asynq.Option) (*asynq.TaskInfo, error) {
	return q.Client.Enqueue(task, opts...)
}

func (q *Queue) EnqueueContext(ctx context.Context, task *asynq.Task, opts ...asynq.Option) (*asynq.TaskInfo, error) {
	return q.Client.EnqueueContext(ctx, task, opts...)
}
