package data

import (
	"edit-your-project-name/config"
	"edit-your-project-name/slog"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func InitDB() {
	DB = initDB(config.MySQL)

	for _, fn := range waitDBFn {
		fn()
	}
}

// ====^

var waitDBFn []func()

func WaitDBExec(fn func()) {
	waitDBFn = append(waitDBFn, fn)
}

func initDB(config config.MySQLConfig) *gorm.DB {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User, config.Pwd, config.Addr, config.DB,
	)), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Error),
	})
	slog.Fatal(slog.PS("MySQL", config.Addr, "ERROR"), err)
	return db
}
