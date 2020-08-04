package datasource

import (
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"github.com/terryli1643/apidemo/libs/configure"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// var dbManger datasource.IDatasourceManager
var l sync.Mutex
var models []interface{}
var db *gorm.DB

// auto create/migrate table if you want
func RegisterModels(models ...interface{}) {
	models = append(models, models...)
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
	// logMode := logger.Default.LogMode(logger.Silent)
	// if config.SqlDebug == 1 {
	// 	logMode = logger.Default.LogMode(logger.Info)
	// }

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       config.DNS,
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on used version
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		// Logger: logMode,
	})

	if err != nil {
		log.WithFields(log.Fields{
			"dns":    config.DNS,
			"driver": config.Driver,
		}).Error("DataSourceManager init error")
		panic(err)
	}

	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(config.IdlePoolSize)
	sqlDB.SetMaxOpenConns(config.MaxPoolSize)
	sqlDB.SetConnMaxLifetime(time.Duration(config.MaxLifeTime) * time.Second)

	if config.AutoCreate {
		if len(models) > 0 {
			db.AutoMigrate(models...)
		}
	}

	log.WithField("db", db).Debug("create a new db connetion")
	return db
}
