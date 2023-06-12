package sqlite

import (
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/utils/opt"

	// "gorm.io/driver/sqlite"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func Provide(opts ...opt.Option) error {
	o := opt.New(opts...)
	var conf Config
	if err := o.UnmarshalConfig(&conf); err != nil {
		return err
	}

	return container.Container.Provide(func() (*gorm.DB, error) {
		db, err := gorm.Open(sqlite.Open(conf.File), &gorm.Config{})
		if err != nil {
			return nil, err
		}

		return db, err
	}, o.DiOptions()...)
}
