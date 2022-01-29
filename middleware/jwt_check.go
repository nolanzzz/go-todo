package middleware

import (
	"github.com/gin-gonic/gin"
	"strings"
	"todo/common/jwt_helper"
	"todo/common/response"
)

func JwtCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" {
			response.Unauthorized(c, "no authorization", nil)
			c.Abort()
			return
		}
		split := strings.Split(header, " ")
		if split[0] != "Bearer" && split[1] == "" {
			response.Unauthorized(c, "cannot access", nil)
			c.Abort()
			return
		}
		decode, err := jwt_helper.Decode(split[1])
		if err != nil {
			response.Unauthorized(c, "decoding token failed: "+err.Error(), nil)
			c.Abort()
			return
		}
		c.Set("X-ID", decode.Wid)
		c.Set("X-USERNAME", decode.Username)
		c.Next()
	}
}
