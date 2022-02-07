package seeder

import (
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"todo/global"
	"todo/model"
	"todo/utils"
)

type users struct{}

var Users = new(users)

func (u *users) Seed(db *gorm.DB) error {
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
	if err := db.Create(&records).Error; err != nil {
		global.LOG.Error("user seeder failed", zap.Error(err))
		return err
	}
	return nil
}

func (u *users) TableName() string {
	return "users"
}

func (u *users) CheckDataExist(db *gorm.DB) bool {
	if errors.Is(db.First(&model.User{}, "username = ?", "user9").Error, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}
