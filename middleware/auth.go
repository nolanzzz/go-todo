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
			response.Unauthorized(c)
			c.Abort()
			return
		}
		split := strings.Split(header, " ")
		if split[0] != "Bearer" && split[1] == "" {
			response.Unauthorized(c)
			c.Abort()
			return
		}
		decode, err := jwt_helper.Decode(split[1])
		if err != nil {
			response.Unauthorized(c)
			c.Abort()
			return
		}
		user := &model.User{}
		err = global.DB.Find(user, "id = ?", decode.Wid).Error
		if err != nil || user.ID == 0 {
			response.NotFound(c)
			c.Abort()
			return
		}
		c.Set("user_id", user.ID)
		c.Next()
	}
}
