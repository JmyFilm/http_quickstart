package db

import (
	"edit-your-project-name/conf"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func InitDB() {
	DB = initDB(conf.MySQL)
}

func initDB(config conf.MySQLConfig) *gorm.DB {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User, config.Pwd, config.Addr, config.DB,
	)), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	conf.FatalExt("MySQL ERROR", err)
	return db
}
