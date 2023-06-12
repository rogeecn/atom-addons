package mysql

import (
	"database/sql"

	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/providers/log"
	"github.com/rogeecn/atom/utils/opt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func Provide(opts ...opt.Option) error {
	o := opt.New(opts...)
	var conf Config
	if err := o.UnmarshalConfig(&conf); err != nil {
		return err
	}

	return container.Container.Provide(func() (*gorm.DB, error) {
		if err := createMySQLDatabase(conf.EmptyDsn(), "mysql", conf.CreateDatabaseSql()); err != nil {
			return nil, err
		}

		mysqlConfig := mysql.Config{
			DSN:                       conf.DSN(), // DSN data source name
			DefaultStringSize:         191,        // string 类型字段的默认长度
			SkipInitializeWithVersion: false,      // 根据版本自动配置
		}

		gormConfig := gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:   conf.Prefix,
				SingularTable: conf.Singular,
			},
			DisableForeignKeyConstraintWhenMigrating: true,
		}

		// TODO: config logger
		// _default := logger.New(NewWriter(log.New(os.Stdout, "\r\n", log.LstdFlags)), logger.Config{
		// 	SlowThreshold: 200 * time.Millisecond,
		// 	LogLevel:      logger.Warn,
		// 	Colorful:      true,
		// })
		// conf.Logger = _default.LogMode(logger.Warn)

		db, err := gorm.Open(mysql.New(mysqlConfig), &gormConfig)
		if err != nil {
			return nil, err
		}

		// config instance
		db.InstanceSet("gorm:table_options", "ENGINE="+conf.Engine)

		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(conf.MaxIdleConns)
		sqlDB.SetMaxOpenConns(conf.MaxOpenConns)

		return db, err
	}, o.DiOptions()...)
}

// createDatabase 创建数据库
func createMySQLDatabase(dsn, driver, createSql string) error {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {
			log.Error(err)
		}
	}(db)
	err = db.Ping()
	if err != nil {
		return err
	}
	_, err = db.Exec(createSql)
	return err
}
