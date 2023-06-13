package swagger

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rogeecn/atom-addons/providers/http"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/utils/opt"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag"
)

type Swagger struct {
	config *Config
	http   http.Service
}

func (swagger *Swagger) Load(spec string) {
	swaggerInfo := &swag.Spec{
		Version:          swagger.config.Version,
		Host:             swagger.config.Host,
		BasePath:         swagger.config.BasePath,
		Schemes:          []string{},
		Title:            swagger.config.Title,
		Description:      swagger.config.Description,
		InfoInstanceName: "swagger",
		SwaggerTemplate:  spec,
		LeftDelim:        "{{",
		RightDelim:       "}}",
	}
	swag.Register(swaggerInfo.InstanceName(), swaggerInfo)
	engine := swagger.http.GetEngine().(*gin.Engine)

	var handler gin.HandlerFunc
	if swagger.config.HandlerConfig != nil {
		handler = ginSwagger.CustomWrapHandler(swagger.config.HandlerConfig, swaggerFiles.Handler)
	} else {
		handler = ginSwagger.WrapHandler(swaggerFiles.Handler)
	}
	engine.GET(fmt.Sprintf("/%s/*any", swagger.config.BaseRoute), handler)
}

func Provide(opts ...opt.Option) error {
	o := opt.New(opts...)
	var config Config
	if err := o.UnmarshalConfig(&config); err != nil {
		return err
	}

	return container.Container.Provide(func(http http.Service) *Swagger {
		if config.BaseRoute == "" {
			config.BaseRoute = "swagger"
		}
		return &Swagger{
			config: &config,
			http:   http,
		}
	}, o.DiOptions()...)
}
