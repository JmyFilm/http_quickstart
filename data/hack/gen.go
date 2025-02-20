package main

import (
	"PROJECTNAME/conf"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func main() {
	conf.Init("")

}

func doGen(sMySQL conf.SMySQL, outPath string, tableName ...string) {
	g := gen.NewGenerator(gen.Config{
		OutPath: outPath,
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	gormDb, _ := gorm.Open(mysql.Open(fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		sMySQL.User, sMySQL.Pwd, sMySQL.Addr, sMySQL.DB,
	)), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	g.UseDB(gormDb)

	var models []any
	for _, name := range tableName {
		models = append(models, g.GenerateModel(name))
	}

	g.ApplyBasic(models...)
	g.Execute()
}
