package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func respJson(c *gin.Context, httpCode int, msg string, data interface{}) {
	c.JSON(httpCode, gin.H{
		"status": httpCode,
		"data":   data,
		"msg":    msg,
	})
}

func Success(c *gin.Context, msg string, data interface{}) {
	respJson(c, http.StatusOK, msg, data)
}

func Fail(c *gin.Context, msg string, data interface{}) {
	respJson(c, http.StatusBadRequest, msg, data)
}

func NotFound(c *gin.Context, msg string, data interface{}) {
	respJson(c, http.StatusNotFound, msg, data)
}
