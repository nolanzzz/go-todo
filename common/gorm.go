package common

import (
	"github.com/jinzhu/gorm"
	"log"
	"todo/global"
	"todo/model"
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

// RegisterTables - Migrate database tables
func RegisterTables(db *gorm.DB) {
	db.AutoMigrate(
		&model.Todo{},
		&model.User{},
	)
	if db.Error != nil {
		log.Fatal("register table failed: ", db.Error)
	}
}
