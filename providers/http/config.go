package http

import (
	"fmt"
)

const DefaultPrefix = "Http"

type Config struct {
	Static *string
	Host   *string
	Port   uint
	Tls    *Tls
	Cors   *Cors
}

type Tls struct {
	Cert string
	Key  string
}

type Cors struct {
	Mode      string
	Whitelist []Whitelist
}

type Whitelist struct {
	AllowOrigin      string
	AllowHeaders     string
	AllowMethods     string
	ExposeHeaders    string
	AllowCredentials bool
}

func (h *Config) Address() string {
	if h.Host == nil {
		return h.PortString()
	}
	return fmt.Sprintf("%s:%d", *h.Host, h.Port)
}

func (h *Config) PortString() string {
	return fmt.Sprintf(":%d", h.Port)
}
