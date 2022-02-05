package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strings"
	"todo/core/jwt_helper"
	"todo/core/response"
	"todo/global"
	"todo/model"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" {
			global.LOG.Error("empty auth header")
			response.Unauthorized(c)
			c.Abort()
			return
		}
		split := strings.Split(header, " ")
		if split[0] != "Bearer" && split[1] == "" {
			global.LOG.Error("Bearer check failed", zap.String("split[0]", split[0]), zap.String("split[1]", split[1]))
			response.Unauthorized(c)
			c.Abort()
			return
		}
		decode, err := jwt_helper.Decode(split[1])
		if err != nil {
			global.LOG.Error("decoding failed", zap.Error(err))
			response.Unauthorized(c)
			c.Abort()
			return
		}
		var user model.User
		err = global.DB.Find(&user, "id = ?", decode.Wid).Error
		if err != nil || user.ID == 0 {
			global.LOG.Error("auth id error", zap.Error(err))
			response.NotFound(c)
			c.Abort()
			return
		}
		c.Set("user_id", user.ID)
		c.Next()
	}
}
