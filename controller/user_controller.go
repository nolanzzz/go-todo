package controller

import (
	"github.com/gin-gonic/gin"
	"todo/common/response"
	"todo/service/user_service"
)

type UserController struct{}

var User *UserController

// Register
// @Tags Users
// @Summary Register new account
// @Produce application/json
// @Success 200 {string} string "{"status":200,"data":{},"msg":"Register new user succeed"}"
// @Router /api/v1/Users/register [post]
func (u *UserController) Register(context *gin.Context) {
	username := context.PostForm("username")
	password := context.PostForm("password")
	if err := user_service.UserServiceApp.Register(username, password); err != nil {
		response.Fail(context, "Register new user failed: "+err.Error(), nil)
		return
	}
	response.Success(context, "Register new user succeed", nil)
}

// Login
// @Tags Users
// @Summary User login
// @Produce application/json
// @Success 200 {string} string "{"status":200,"data":{"token":string},"msg":"Login succeed"}"
// @Router /api/v1/Users/login [post]
func (u *UserController) Login(context *gin.Context) {
	username := context.PostForm("username")
	password := context.PostForm("password")
	token, err := user_service.UserServiceApp.Login(username, password)
	if err != nil {
		response.Fail(context, "Login failed, please confirm your username and password", nil)
		return
	}
	response.Success(context, "Login succeed", gin.H{"token": token})
}
