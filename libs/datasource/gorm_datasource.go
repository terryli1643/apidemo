package datasource

import (
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"github.com/terryli1643/apidemo/libs/configure"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// var dbManger datasource.IDatasourceManager
var l sync.Mutex
var models []interface{}
var db *gorm.DB

// auto create/migrate table if you want
func RegisterModels(m ...interface{}) {
	l.Lock()
	models = append(models, m...)
	l.Unlock()
}

func GetDB() *gorm.DB {
	if db == nil {
		l.Lock()
		defer l.Unlock()
		if db == nil {
			db = openConn(configure.ServerConfig.DataSource)
		}
	}
	return db
}

func openConn(config configure.TDataSourceConfig) (db *gorm.DB) {
	sqlLog := logger.New(log, logger.Config{
		SlowThreshold: 100 * time.Millisecond,
		LogLevel:      logger.Info,
		Colorful:      true,
	})

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       config.DSN,
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on used version
	}), &gorm.Config{
		Logger:                                   sqlLog,
		PrepareStmt:                              true,
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_",
			SingularTable: true,
		},
	})

	if err != nil {
		log.WithFields(logrus.Fields{
			"dns":    config.DSN,
			"driver": config.Driver,
		}).Error("DataSourceManager init error")
		panic(err)
	}

	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(config.IdlePoolSize)
	sqlDB.SetMaxOpenConns(config.MaxPoolSize)
	sqlDB.SetConnMaxLifetime(time.Duration(config.MaxLifeTime) * time.Second)

	if config.AutoMigrate && len(models) > 0 {
		db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(models...)
	}

	log.WithField("db", db).Debug("create a new db connetion")
	return db
}
