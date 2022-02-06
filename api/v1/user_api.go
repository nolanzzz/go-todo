package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"todo/global"
	"todo/model"
	"todo/model/response"
	"todo/service"
)

type UserApi struct{}

var User *UserApi

// Register
// @Tags	Users
// @Summary Register new account
// @Accept 	application/json
// @Produce application/json
// @Success 200 {string} string "{"status":200,"data":{},"msg":"user register succeed"}"
// @Router 	/api/v1/Users/register [post]
func (u *UserApi) Register(c *gin.Context) {
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
// @Tags 	Users
// @Summary User login
// @Accept 	application/json
// @Produce application/json
// @Success 200 {string} string "{"status":200,"data":{"token":string},"msg":"user login succeed"}"
// @Router 	/api/v1/Users/login [post]
func (u *UserApi) Login(c *gin.Context) {
	var user model.User
	_ = c.ShouldBind(&user)
	token, err := service.UserServiceApp.Login(user)
	if err != nil {
		global.LOG.Error("user login failed", zap.Error(err))
		response.FailWithMessage(c, err.Error())
		return
	}
	if global.CONFIG.System.UseRedisJWT {
		// Store jwt token to redis
		_, err = service.JwtServiceApp.GetRedisJWT(user.Username)
		if err == redis.Nil || err == nil {
			if err = service.JwtServiceApp.SetRedisJWT(user.Username, token); err != nil {
				global.LOG.Error("set redis jwt failed", zap.Error(err))
				response.FailWithMessage(c, err.Error())
				return
			}
		} else if err != nil {
			global.LOG.Error("access redis jwt failed", zap.Error(err))
			response.FailWithMessage(c, err.Error())
			return
		}
	}
	response.OkWithDetails(c, "user login succeed", gin.H{"token": token})
}
