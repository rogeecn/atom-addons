package fiber

import (
	"time"

	"github.com/rogeecn/atom-addons/providers/http"
	"github.com/rogeecn/atom-addons/providers/log"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/utils/opt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func DefaultProvider() container.ProviderContainer {
	return container.ProviderContainer{
		Provider: Provide,
		Options: []opt.Option{
			opt.Prefix(http.DefaultPrefix),
		},
	}
}

type Service struct {
	conf   *http.Config
	Engine *fiber.App
}

func (e *Service) Use(middleware ...interface{}) fiber.Router {
	return e.Engine.Use(middleware...)
}

func (e *Service) GetEngine() interface{} {
	return e.Engine
}

func (e *Service) Serve() error {
	if e.conf.Tls != nil {
		return e.Engine.ListenTLS(e.conf.PortString(), e.conf.Tls.Cert, e.conf.Tls.Key)
	}
	return e.Engine.Listen(e.conf.PortString())
}

func Provide(opts ...opt.Option) error {
	o := opt.New(opts...)
	var config http.Config
	if err := o.UnmarshalConfig(&config); err != nil {
		return err
	}

	return container.Container.Provide(func(l *log.Logger) (http.Service, error) {
		engine := fiber.New(fiber.Config{
			EnablePrintRoutes: true,
			StrictRouting:     true,
		})
		engine.Use(recover.New())

		if config.StaticRoute != nil && config.StaticPath != nil {
			engine.Use(config.StaticRoute, config.StaticPath)
		}

		engine.Use(logger.New(logger.Config{
			Format:     "[${ip}:${port}] - [${time}] - ${method} - ${status} - ${path} ${latency} ${ua} \n",
			TimeFormat: time.RFC1123,
			TimeZone:   "Asia/Shanghai",
		}))

		return &Service{Engine: engine, conf: &config}, nil
	}, o.DiOptions()...)
}
