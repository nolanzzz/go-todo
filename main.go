package main

import (
	"todo/core"
	"todo/global"
	"todo/router"
)

// @title go-todo
// @version 1.0
// @description This is a api server for a shared task management tool

// @contact.name Nolan Zhang
// @contact.url https://github.com/nolanzzz

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /api/v1
func main() {
	global.VP = core.Viper()    // initialize viper and load config
	global.LOG = core.Zap()     // initialize logger
	global.DB = core.Gorm()     // initialize gorm database connection
	global.REDIS = core.Redis() // initialize Redis db
	core.InitScheduler()        // start cron scheduler

	r := router.InitApiRouter()
	_ = r.Run(global.CONFIG.System.Addr())
}
