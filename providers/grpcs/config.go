package grpcs

import (
	"fmt"
)

const DefaultPrefix = "Grpc"

type Config struct {
	Host *string
	Port uint
	Tls  *Tls
}

type Tls struct {
	CA   string
	Cert string
	Key  string
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
