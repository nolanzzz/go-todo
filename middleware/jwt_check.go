package middleware

import (
	"github.com/gin-gonic/gin"
	"strings"
	"todo/common/jwt_helper"
	"todo/common/response"
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
		c.Set("user_id", decode.Wid)
		c.Next()
	}
}
