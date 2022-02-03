package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"todo/common"
	"todo/global"
	"todo/router"
)

func init() {
	global.DB = common.Gorm()
	global.LOG = common.Zap()
}

func main() {
	r := router.InitApiRouter()
	r.Run(":8080")
}
