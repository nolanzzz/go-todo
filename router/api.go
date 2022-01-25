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
		todoGroup.POST("/", controller.Todo.Store)
		todoGroup.GET("/", controller.Todo.All)
		todoGroup.GET("/:id", controller.Todo.Show)
		todoGroup.PUT("/:id", controller.Todo.Update)
	}
	return router
}
