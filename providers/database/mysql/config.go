package mysql

import (
	"fmt"

	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/utils/opt"
)

const DefaultPrefix = "MySQL"

func DefaultProvider() container.ProviderContainer {
	return container.ProviderContainer{
		Provider: Provide,
		Options: []opt.Option{
			opt.Prefix(DefaultPrefix),
		},
	}
}

// MySQL database config
type Config struct {
	Host         string
	Port         uint
	Database     string
	Username     string
	Password     string
	Prefix       string // 表前缀
	Singular     bool   // 是否开启全局禁用复数，true表示开启
	MaxIdleConns int    // 空闲中的最大连接数
	MaxOpenConns int    // 打开到数据库的最大连接数
	Engine       string // 数据库引擎，默认InnoDB
}

func (m *Config) CreateDatabaseSql() string {
	return fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;", m.Database)
}

func (m *Config) EmptyDsn() string {
	dsnTpl := "%s@tcp(%s:%d)/"

	authString := func() string {
		if len(m.Password) > 0 {
			return m.Username + ":" + m.Password
		}
		return m.Username
	}

	return fmt.Sprintf(dsnTpl, authString(), m.Host, m.Port)
}

// DSN connection dsn
func (m *Config) DSN() string {
	dsnTpl := "%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"

	authString := func() string {
		if len(m.Password) > 0 {
			return m.Username + ":" + m.Password
		}
		return m.Username
	}

	return fmt.Sprintf(dsnTpl, authString(), m.Host, m.Port, m.Database)
}
