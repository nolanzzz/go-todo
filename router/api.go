package router

import (
	"github.com/gin-gonic/gin"
	"todo/http/controller"
)

type Router struct{}

func InitApiRouter() *gin.Engine {
	router := gin.Default()
	todoGroup := router.Group("/app/controller/todos")
	{
		TodoApi := &controller.TodoApi{}
		todoGroup.POST("/", TodoApi.Store)
		todoGroup.GET("/", TodoApi.All)
		todoGroup.GET("/:id", TodoApi.Show)
		todoGroup.PUT("/:id", TodoApi.Update)
	}
	return router
}
