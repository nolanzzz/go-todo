package user_service

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"todo/common/hash"
	"todo/common/jwt_helper"
	"todo/global"
	"todo/model"
)

type UserService struct{}

var UserServiceApp = new(UserService)

func (u *UserService) Register(user model.User) error {
	// Check if username name exists
	checkUser, err := u.GetUserByUsername(user.Username)
	if err == nil || checkUser.ID != 0 {
		return errors.New("username already used")
	}
	// Generate encrypted password
	hashed, err := hash.NewHash().Make([]byte(user.Password))
	if err != nil {
		return err
	}
	user.Password = string(hashed)
	err = global.DB.Create(&user).Error
	return err
}

func (u *UserService) Login(username, password string) (string, error) {
	user, err := u.GetUserByUsername(username)
	if err != nil {
		return "", err
	}
	if user.ID == 0 {
		return "", errors.New("user not found")
	}
	// Check password
	err = hash.NewHash().Check([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("wrong password")
	}

	// JWT
	claims := jwt_helper.Claims{
		Username:       user.Username,
		Wid:            strconv.Itoa(int(user.ID)),
		StandardClaims: jwt.StandardClaims{},
	}
	// Generate token
	token, err := jwt_helper.Encode(claims)
	if err != nil {
		return "", errors.New("generating token failed")
	}
	return token, nil
}

func (u *UserService) GetUserByUsername(username string) (model.User, error) {
	user := model.User{}
	err := global.DB.Where("username = ?", username).Find(&user).Error
	return user, err
}
