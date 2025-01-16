package main

import (
	"PROJECTNAME/conf"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var g *gen.Generator

func main() {
	conf.Init()
	initGen()

	g.ApplyBasic(
		g.GenerateModel("xxx"),
		g.GenerateModel("xxx"),
		g.GenerateModel("xxx"),
	)

	g.Execute()
}

func initGen() {
	g = gen.NewGenerator(gen.Config{
		OutPath: "./data/dao",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	gormDb, _ := gorm.Open(mysql.Open(fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.MySQL.User, conf.MySQL.Pwd, conf.MySQL.Addr, conf.MySQL.DB,
	)), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	g.UseDB(gormDb)
}
