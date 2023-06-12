package sqlite

import (
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/utils/opt"
)

const DefaultPrefix = "SQLite"

func DefaultProvider() container.ProviderContainer {
	return container.ProviderContainer{
		Provider: Provide,
		Options: []opt.Option{
			opt.Prefix(DefaultPrefix),
		},
	}
}

type Config struct {
	File string
}

func (m *Config) CreateDatabaseSql() string {
	return ""
}

func (m *Config) EmptyDsn() string {
	return m.File
}
