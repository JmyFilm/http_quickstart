package main

import (
	"edit-your-project-name/config"
	"edit-your-project-name/data"
	"edit-your-project-name/handler"
	"edit-your-project-name/slog"
	"edit-your-project-name/utils"
)

func main() {
	config.InitConfig()
	utils.InitNo()
	slog.InitLog()
	data.InitRDB()
	data.InitDB()
	handler.InitHandler()
}
