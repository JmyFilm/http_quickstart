package data

import (
	"PROJECTNAME/conf"
	"PROJECTNAME/data/dao"
	"PROJECTNAME/xlog"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

var DB *dao.Query

func InitDB() {
	DB = dao.Use(initDB(conf.MySQL))

	for _, fn := range waitDBFn {
		fn()
	}
}

// ====^

var waitDBFn []func()

func WaitDBExec(fn func()) {
	waitDBFn = append(waitDBFn, fn)
}

func initDB(config conf.SMySQL) *gorm.DB {
	var LogLevel logger.LogLevel
	if conf.Log.DebugInfo {
		LogLevel = logger.Info
	} else {
		LogLevel = logger.Silent
	}

	db, err := gorm.Open(mysql.Open(fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User, config.Pwd, config.Addr, config.DB,
	)), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(LogLevel),
	})
	xlog.Fatal(xlog.PS("MySQL", config.Addr, "ERROR"), err)

	if sqlDb, err := db.DB(); err == nil {
		sqlDb.SetMaxOpenConns(config.MaxConns) // 数据库最大连接数

		maxIdleConns := config.MaxConns / 100
		if maxIdleConns < 2 {
			maxIdleConns = 2
		} else if maxIdleConns > 20 {
			maxIdleConns = 20
		}
		sqlDb.SetMaxIdleConns(maxIdleConns) // 数据库最大空闲连接数

		sqlDb.SetConnMaxLifetime(time.Hour)        // 连接最长存活时间
		sqlDb.SetConnMaxIdleTime(time.Minute * 10) // 空闲连接最长存活时间
	}

	var _now = struct {
		Now time.Time `gorm:"now"`
	}{}
	_ = db.Raw("SELECT now() AS now").Scan(&_now).Error
	xlog.Info("MySQL", config.Addr, "Time:", _now.Now)
	return db
}
