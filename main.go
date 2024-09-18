package main

import (
	"edit-your-project-name/conf"
	"edit-your-project-name/db"
	"edit-your-project-name/handler"
)

func main() {
	conf.InitConfig()
	conf.InitLog()
	db.InitRDB()
	db.InitDB()
	handler.InitHandler()
}
