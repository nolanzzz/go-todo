package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"todo/common"
	"todo/global"
	"todo/router"
)

func init() {
	global.VP = common.Viper() // initialize viper and load config
	global.DB = common.Gorm()  // initialize gorm database connection
	global.LOG = common.Zap()  // initialize logger
}

func main() {
	r := router.InitApiRouter()
	r.Run(":8080")
}
