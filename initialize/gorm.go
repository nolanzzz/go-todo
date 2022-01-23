package initialize

import (
	"github.com/jinzhu/gorm"
	"log"
	"os"
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
	err := db.AutoMigrate(
		&model.TodoModel{},
	)
	if err != nil {
		log.Fatal("register table failed")
		os.Exit(0)
	}
}
