package core

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"todo/database/seeder"
	"todo/global"
	"todo/model"
)

// Gorm - initialize global database instance
func Gorm() *gorm.DB {
	// open a DB connection
	db, err := gorm.Open(mysql.Open(Dsn()), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to DB: " + err.Error())
	}
	if global.CONFIG.Database.Migrate {
		err = db.AutoMigrate(
			&model.User{},
			&model.Todo{},
		)
		if err != nil {
			panic("Migration failed: " + err.Error())
		}
	}
	if global.CONFIG.Database.Seed {
		seeder.RunSeeders(db)
	}
	return db
}

func Dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s",
		global.CONFIG.Mysql.Username,
		global.CONFIG.Mysql.Password,
		global.CONFIG.Mysql.Host,
		global.CONFIG.Mysql.Port,
		global.CONFIG.Mysql.Database,
		global.CONFIG.Mysql.Config,
	)
}
