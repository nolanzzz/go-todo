package user_service

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"todo/core/hash"
	"todo/core/jwt_helper"
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
	var hashed []byte
	hashed, err = hash.NewHash().Make([]byte(user.Password))
	if err != nil {
		return err
	}
	user.Password = string(hashed)
	err = global.DB.Create(&user).Error
	return err
}

func (u *UserService) Login(user model.User) (string, error) {
	userDB, _ := u.GetUserByUsername(user.Username)
	if userDB.ID == 0 {
		return "", errors.New("user not found")
	}
	// Check password
	err := hash.NewHash().Check([]byte(userDB.Password), []byte(user.Password))
	if err != nil {
		return "", errors.New("wrong password")
	}
	// Generate token
	claims := jwt_helper.Claims{
		Username:       user.Username,
		Wid:            strconv.Itoa(int(userDB.ID)),
		StandardClaims: jwt.StandardClaims{},
	}
	var token string
	token, err = jwt_helper.Encode(claims)
	return token, err
}

func (u *UserService) GetUserByUsername(username string) (model.User, error) {
	user := model.User{}
	err := global.DB.Where("username = ?", username).Find(&user).Error
	return user, err
}
