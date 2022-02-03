package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
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
			log.Println("empty header")
			response.Unauthorized(c)
			c.Abort()
			return
		}
		split := strings.Split(header, " ")
		if split[0] != "Bearer" && split[1] == "" {
			log.Println("split[0]: ", split[0])
			log.Println("split[1]: ", split[1])
			response.Unauthorized(c)
			c.Abort()
			return
		}
		decode, err := jwt_helper.Decode(split[1])
		if err != nil {
			log.Println(err.Error())
			response.Unauthorized(c)
			c.Abort()
			return
		}
		var user model.User
		err = global.DB.Find(&user, "id = ?", decode.Wid).Error
		if err != nil || user.ID == 0 {
			log.Println("id: ", decode.Wid, " ", err.Error())
			response.NotFound(c)
			c.Abort()
			return
		}
		c.Set("user_id", user.ID)
		c.Next()
	}
}
