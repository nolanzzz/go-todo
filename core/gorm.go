package core

import (
	"github.com/jinzhu/gorm"
	"todo/global"
)

// Gorm - initialize global database instance
func Gorm() *gorm.DB {
	// open a DB connection
	db, err := gorm.Open("mysql", getConfig())
	if err != nil {
		panic("Failed to connect to DB")
	}
	return db
}

func getConfig() string {
	return global.CONFIG.Mysql.Username + ":" + global.CONFIG.Mysql.Password +
		"@tcp(" + global.CONFIG.Mysql.Host + ":" + global.CONFIG.Mysql.Port + ")/" +
		global.CONFIG.Mysql.Database + "?" + global.CONFIG.Mysql.Config
}
