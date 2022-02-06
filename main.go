package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"todo/core"
	"todo/database/seeder"
	"todo/global"
	"todo/router"
)

func init() {
	global.VP = core.Viper()    // initialize viper and load config
	global.DB = core.Gorm()     // initialize gorm database connection
	global.LOG = core.Zap()     // initialize logger
	global.REDIS = core.Redis() // initialize Redis db
	core.InitTimer()            // start cron timer
	if global.CONFIG.Database.Migrate {
		seeder.RunSeeders(global.DB)
	}
}

func main() {
	r := router.InitApiRouter()
	_ = r.Run(global.CONFIG.System.Addr())
}
