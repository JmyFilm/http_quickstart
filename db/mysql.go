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
	var err error

	if DB, err = gorm.Open(mysql.Open(fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.MySQL.User, conf.MySQL.Pwd, conf.MySQL.Host, conf.MySQL.Port, conf.MySQL.DB,
	)), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}); err != nil {
		conf.FatalExt("MySQL ERROR", err.Error())
	}
}
