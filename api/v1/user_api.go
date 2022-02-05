package v1

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"todo/global"
	"todo/model"
	"todo/model/response"
	"todo/service"
)

type UserController struct{}

var User *UserController

// Register
// @Tags Users
// @Summary Register new account
// @Accept application/json
// @Produce application/json
// @Success 200 {string} string "{"status":200,"data":{},"msg":"user register succeed"}"
// @Router /api/v1/Users/register [post]
func (u *UserController) Register(c *gin.Context) {
	var user model.User
	_ = c.ShouldBind(&user)
	if err := service.UserServiceApp.Register(user); err != nil {
		global.LOG.Error("user register failed", zap.Error(err))
		response.FailWithMessage(c, err.Error())
		return
	}
	response.OkWithMessage(c, "user register succeed")
}

// Login
// @Tags Users
// @Summary User login
// @Accept application/json
// @Produce application/json
// @Success 200 {string} string "{"status":200,"data":{"token":string},"msg":"user login succeed"}"
// @Router /api/v1/Users/login [post]
func (u *UserController) Login(c *gin.Context) {
	var user model.User
	_ = c.ShouldBind(&user)
	token, err := service.UserServiceApp.Login(user)
	if err != nil {
		global.LOG.Error("user login failed", zap.Error(err))
		response.FailWithMessage(c, err.Error())
		return
	}
	response.OkWithDetails(c, "user login succeed", gin.H{"token": token})
}
