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
		response.FailWithMessage(context, "user register failed")
		return
	}
	response.OkWithMessage(context, "user register succeed")
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
		response.FailWithMessage(context, "user login failed")
		return
	}
	response.OkWithDetails(context, "user login succeed", gin.H{"token": token})
}
