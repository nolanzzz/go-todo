package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"todo/global"
	"todo/initialize"
	"todo/router"
)

func init() {
	global.DB = initialize.Gorm()
	if global.DB != nil {
		initialize.RegisterTables(global.DB)
	}
}

func main() {
	r := router.InitApiRouter()
	r.Run(":8080")
}
