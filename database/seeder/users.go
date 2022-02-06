package seeder

import (
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"todo/global"
	"todo/model"
	"todo/utils"
)

type users struct{}

var Users = new(users)

func (u *users) Seed(db *gorm.DB) error {
	//record := model.User{Username: "user1", Password: "12345"}
	records := []model.User{
		{Username: "user1", Password: utils.CreatePass("12345")},
		{Username: "user2", Password: utils.CreatePass("12345")},
		{Username: "user3", Password: utils.CreatePass("12345")},
		{Username: "user4", Password: utils.CreatePass("12345")},
		{Username: "user5", Password: utils.CreatePass("12345")},
		{Username: "user6", Password: utils.CreatePass("12345")},
		{Username: "user7", Password: utils.CreatePass("12345")},
		{Username: "user8", Password: utils.CreatePass("12345")},
		{Username: "user9", Password: utils.CreatePass("12345")},
	}
	for _, record := range records {
		if err := db.Create(&record).Error; err != nil {
			global.LOG.Error("user seeder failed", zap.Error(err))
			return err
		}
	}
	return nil
}

func (u *users) TableName() string {
	return "users"
}