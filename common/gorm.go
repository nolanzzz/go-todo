package common

import (
	"github.com/jinzhu/gorm"
	"log"
	"todo/model"
)

// Gorm - initialize global database instance
func Gorm() *gorm.DB {
	// open a DB connection
	db, err := gorm.Open("mysql", "root:@/demo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Failed to connect to DB")
	}
	return db
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
