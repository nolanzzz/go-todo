package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"todo/common/response"
	"todo/model"
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
func (u *UserController) Register(c *gin.Context) {
	var user model.User
	_ = c.ShouldBind(&user)
	if err := user_service.UserServiceApp.Register(user); err != nil {
		log.Println(err.Error())
		response.FailWithMessage(c, "user register failed")
		return
	}
	response.OkWithMessage(c, "user register succeed")
}

// Login
// @Tags Users
// @Summary User login
// @Produce application/json
// @Success 200 {string} string "{"status":200,"data":{"token":string},"msg":"Login succeed"}"
// @Router /api/v1/Users/login [post]
func (u *UserController) Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	token, err := user_service.UserServiceApp.Login(username, password)
	if err != nil {
		response.FailWithMessage(c, "user login failed")
		return
	}
	response.OkWithDetails(c, "user login succeed", gin.H{"token": token})
}
