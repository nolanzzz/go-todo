package user_service

import (
	"errors"
	"todo/common/hash"
	"todo/global"
	"todo/model"
)

type UserService struct{}

var UserServiceApp = new(UserService)

func (u *UserService) Register(username, password string) error {
	// Check if username name exists
	var users []model.User
	if err := global.DB.Where("username = ?", username).Find(&users).Error; err != nil {
		return err
	}
	if len(users) > 0 {
		return errors.New("username already used")
	}
	// Generate encrypted password
	hashed, err := hash.NewHash().Make([]byte(password))
	if err != nil {
		return err
	}
	user := model.User{
		Username: username,
		Password: string(hashed),
	}
	err = global.DB.Create(&user).Error
	return err
}

func (u *UserService) Login(username, password string) error {
	user, err := u.GetUserByUsername(username)
	if err != nil {
		return err
	}
	// Check password
	return hash.NewHash().Check([]byte(user.Password), []byte(password))
}

func (u *UserService) GetUserByUsername(username string) (model.User, error) {
	user := model.User{}
	err := global.DB.Where("username = ?", username).Find(&user).Error
	return user, err
}
