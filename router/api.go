package router

import (
	"github.com/gin-gonic/gin"
	v1 "todo/api/v1"
)

type Router struct{}

func InitApiRouter() *gin.Engine {
	router := gin.Default()
	todoGroup := router.Group("/api/v1/todos")
	{
		TodoApi := &v1.TodoApi{}
		todoGroup.POST("/", TodoApi.CreateTodo)
		todoGroup.GET("/", TodoApi.FetchAllTodos)
		//todoGroup.GET("/:id", fetchSingleTodo)
		//todoGroup.PUT("/:id", updateTodo)
		//todoGroup.DELETE("/:id", deleteTodo)
	}
	return router
}
