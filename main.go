package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"todo/core"
	"todo/global"
	"todo/router"
)

func init() {
	global.VP = core.Viper() // initialize viper and load config
	global.DB = core.Gorm()  // initialize gorm database connection
	global.LOG = core.Zap()  // initialize logger
}

func main() {
	r := router.InitApiRouter()
	r.Run(global.CONFIG.System.Addr())
}
