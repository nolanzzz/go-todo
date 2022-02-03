package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"todo/common/zap"
	"todo/global"
	"todo/initialize"
	"todo/router"
)

func init() {
	global.DB = initialize.Gorm()
	global.LOG = zap.Zap()
}

func main() {
	r := router.InitApiRouter()
	r.Run(":8080")
}
