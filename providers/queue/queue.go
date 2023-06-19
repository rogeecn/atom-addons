package queue

import (
	"context"

	"github.com/hibiken/asynq"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/utils/opt"
)

func Provide(opts ...opt.Option) error {
	o := opt.New(opts...)
	var config Config
	if err := o.UnmarshalConfig(&config); err != nil {
		return err
	}
	return container.Container.Provide(func() (*Queue, error) {
		srv := asynq.NewServer(
			asynq.RedisClientOpt{Addr: ""},
			asynq.Config{
				// Specify how many concurrent workers to use
				Concurrency: 10,
				// Optionally specify multiple queues with different priority.
				Queues: map[string]int{
					"critical": 6,
					"default":  3,
					"low":      1,
				},
				// See the godoc for other configuration options
			},
		)

		mux := asynq.NewServeMux()
		return &Queue{
			Server: srv,
			Mux:    mux,
		}, nil
	}, o.DiOptions()...)
}

type QueueHandler interface {
	ProcessTask(context.Context, *asynq.Task) error
	GetType() string
}

type Queue struct {
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
