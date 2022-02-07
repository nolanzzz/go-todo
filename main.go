package main

import (
	"todo/core"
	"todo/global"
	"todo/router"
)

// @title go-todo
// @version 1.0
// @description This is a api server for a shared task management tool

// @contact.name Ruizhe Zhang
// @contact.url https://github.com/nolanzzz

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /api/v1
func init() {
	global.VP = core.Viper()    // initialize viper and load config
	global.DB = core.Gorm()     // initialize gorm database connection
	global.LOG = core.Zap()     // initialize logger
	global.REDIS = core.Redis() // initialize Redis db
	core.InitScheduler()        // start cron timer
}

func main() {
	r := router.InitApiRouter()
	_ = r.Run(global.CONFIG.System.Addr())
}
