package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/rogeecn/atom/container"
	"github.com/spf13/viper"
)

func Load(file, app string) (*viper.Viper, error) {
	v := viper.NewWithOptions(viper.KeyDelimiter("_"))
	v.AutomaticEnv()

	if file == "" {
		v.SetConfigType("toml")
		v.SetConfigName(app + ".toml")

		// execute path
		execPath, err := os.Executable()
		if err == nil {
			v.AddConfigPath(filepath.Dir(execPath))
		}

		// home path
		homePath, err := os.UserHomeDir()
		if err == nil {
			v.AddConfigPath(homePath)
			v.AddConfigPath(homePath + "/" + app)
			v.AddConfigPath(homePath + "/.config")
			v.AddConfigPath(homePath + "/.config/" + app)
		}

		v.AddConfigPath("/etc")
		v.AddConfigPath("/etc/" + app)
		v.AddConfigPath("/usr/local/etc")
		v.AddConfigPath("/usr/local/etc/" + app)
	} else {
		v.SetConfigFile(file)
	}

	err := v.ReadInConfig()
	log.Println("config file:", v.ConfigFileUsed())
	if err != nil {
		return nil, errors.Wrap(err, "config file read error")
	}

	err = container.Container.Provide(func() (*viper.Viper, error) {
		return v, nil
	})
	if err != nil {
		return nil, err
	}

	return v, nil
}
