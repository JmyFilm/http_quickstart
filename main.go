package main

import (
	"PROJECTNAME/conf"
	"PROJECTNAME/data"
	"PROJECTNAME/handler"
	"PROJECTNAME/utils"
	"PROJECTNAME/xlog"
)

func main() {
	conf.Init("")
	utils.InitNo(conf.App.AppId)
	xlog.Init("v1.0.0") // TODO
	data.InitDB()
	data.InitRDB()
	handler.Init()
}
