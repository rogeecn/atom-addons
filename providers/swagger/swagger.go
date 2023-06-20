package swagger

import (
	"fmt"
	"strings"

	jsonpatch "github.com/evanphx/json-patch"
	"github.com/gofiber/fiber/v2"
	"github.com/rogeecn/atom-addons/providers/http"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/utils/opt"
	fiberSwagger "github.com/swaggo/fiber-swagger"
	"github.com/swaggo/swag"
)

type Swagger struct {
	config *Config
	http   http.Service
}

const infoTpl = `{"schemes": "__ marshal .Schemes __",
"swagger": "2.0",
"info": {
	"description": "{{escape .Description}}",
	"title": "{{.Title}}",
	"contact": {},
	"version": "{{.Version}}"
},
"host": "{{.Host}}",
"basePath": "{{.BasePath}}"}`

func (swagger *Swagger) Load(spec string) error {
	original := []byte(spec)
	target := []byte(infoTpl)
	patch, err := jsonpatch.MergeMergePatches(original, target)
	if err != nil {
		return err
	}

	merged, err := jsonpatch.MergePatch(original, patch)
	if err != nil {
		return err
	}

	spec = strings.NewReplacer(`"__`, "{{", `__"`, "}}").Replace(string(merged))

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
	engine := swagger.http.GetEngine().(*fiber.App)

	engine.Get(fmt.Sprintf("/%s/*", swagger.config.BaseRoute), fiberSwagger.WrapHandler)
	return nil
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
