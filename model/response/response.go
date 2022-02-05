package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func result(c *gin.Context, httpCode int, msg string, data interface{}) {
	c.JSON(httpCode, gin.H{
		"msg":  msg,
		"data": data,
	})
}

// Ok returns with a default message
func Ok(c *gin.Context) {
	result(c, http.StatusOK, "succeed", nil)
}

// OkWithMessage returns a given message
func OkWithMessage(c *gin.Context, msg string) {
	result(c, http.StatusOK, msg, nil)
}

// OkWithData returns a default message and given data
func OkWithData(c *gin.Context, data interface{}) {
	result(c, http.StatusOK, "succeed", data)
}

// OkWithDetails returns both given message and data
func OkWithDetails(c *gin.Context, msg string, data interface{}) {
	result(c, http.StatusOK, msg, data)
}

// Fail returns with a default message
func Fail(c *gin.Context) {
	result(c, http.StatusBadRequest, "failed", nil)
}

// FailWithMessage returns a given message
func FailWithMessage(c *gin.Context, msg string) {
	result(c, http.StatusBadRequest, msg, nil)
}

func NotFound(c *gin.Context) {
	result(c, http.StatusNotFound, "resource not found", nil)
}

func Unauthorized(c *gin.Context) {
	result(c, http.StatusUnauthorized, "event unauthorized", nil)
}
