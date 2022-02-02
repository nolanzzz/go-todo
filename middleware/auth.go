package middleware

import (
	"github.com/gin-gonic/gin"
	"strings"
	"todo/common/jwt_helper"
	"todo/common/response"
	"todo/global"
	"todo/model"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" {
			response.Unauthorized(c, "unauthorized", nil)
			c.Abort()
			return
		}
		split := strings.Split(header, " ")
		if split[0] != "Bearer" && split[1] == "" {
			response.Unauthorized(c, "invalid authorization", nil)
			c.Abort()
			return
		}
		decode, err := jwt_helper.Decode(split[1])
		if err != nil {
			response.Unauthorized(c, "decoding token failed: "+err.Error(), nil)
			c.Abort()
			return
		}
		user := &model.User{}
		err = global.DB.Find(user, "id = ?", decode.Wid).Error
		if err != nil || user.ID == 0 {
			response.NotFound(c, "user not found", nil)
			c.Abort()
			return
		}
		c.Set("user_id", user.ID)
		c.Next()
	}
}
