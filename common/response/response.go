package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func respJson(context *gin.Context, httpCode int, msg string, data interface{}) {
	context.JSON(httpCode, gin.H{
		"status": httpCode,
		"data":   data,
		"msg":    msg,
	})
}

func Success(context *gin.Context, msg string, data interface{}) {
	respJson(context, http.StatusOK, msg, data)
}

func Fail(context *gin.Context, msg string, data interface{}) {
	respJson(context, http.StatusBadRequest, msg, data)
}

func NotFound(context *gin.Context, msg string, data interface{}) {
	respJson(context, http.StatusNotFound, msg, data)
}

func Unauthorized(context *gin.Context, msg string, data interface{}) {
	respJson(context, http.StatusUnauthorized, msg, data)
}
