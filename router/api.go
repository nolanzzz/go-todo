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
		Todo := &controller.TodoController{}
		todoGroup.POST("/", Todo.Store)
		todoGroup.GET("/", Todo.All)
		todoGroup.GET("/:id", Todo.Show)
		todoGroup.PUT("/:id", Todo.Update)
	}
	return router
}
