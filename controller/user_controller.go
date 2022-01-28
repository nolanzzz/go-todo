package controller

import (
	"github.com/gin-gonic/gin"
	"todo/common/response"
	"todo/service/user_service"
)

type UserController struct{}

var User *UserController

func (u *UserController) Register(context *gin.Context) {
	username := context.PostForm("username")
	password := context.PostForm("password")
	if err := user_service.UserServiceApp.Register(username, password); err != nil {
		response.Fail(context, "Register new user failed: "+err.Error(), nil)
		return
	}
	response.Success(context, "Register new user succeed", nil)
}

func (u *UserController) Login(context *gin.Context) {
	username := context.PostForm("username")
	password := context.PostForm("password")
	if err := user_service.UserServiceApp.Login(username, password); err != nil {
		response.Fail(context, "Login failed, please confirm your username and password", nil)
		return
	}
	response.Success(context, "Login succeed", nil)
}
