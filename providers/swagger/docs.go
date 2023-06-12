package swagger

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/contracts"
	"github.com/rogeecn/atom/providers/http"
	"github.com/rogeecn/atom/utils/opt"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag"
)

func Provide(opts ...opt.Option) error {
	o := opt.New(opts...)
	var config Config
	if err := o.UnmarshalConfig(&config); err != nil {
		return err
	}

	return container.Container.Provide(func(http http.Service) contracts.Initial {
		if config.BaseRoute == "" {
			config.BaseRoute = "swagger"
		}

		swaggerInfo := &swag.Spec{
			Version:          config.Version,
			Host:             config.Host,
			BasePath:         config.BasePath,
			Schemes:          []string{},
			Title:            config.Title,
			Description:      config.Description,
			InfoInstanceName: "swagger",
			SwaggerTemplate:  docTemplate,
			LeftDelim:        "{{",
			RightDelim:       "}}",
		}
		swag.Register(swaggerInfo.InstanceName(), swaggerInfo)
		engine := http.GetEngine().(*gin.Engine)

		var handler gin.HandlerFunc
		if config.HandlerConfig != nil {
			handler = ginSwagger.CustomWrapHandler(config.HandlerConfig, swaggerFiles.Handler)
		} else {
			handler = ginSwagger.WrapHandler(swaggerFiles.Handler)
		}
		engine.GET(fmt.Sprintf("/%s/*any", config.BaseRoute), handler)

		return nil
	}, o.DiOptions()...)
}
