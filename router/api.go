package router

import (
	"github.com/gin-gonic/gin"
	"todo/http/controller"
)

type Router struct{}

func InitApiRouter() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/api/v1")
	todoGroup := v1.Group("/todo")
	{
		todoGroup.POST("/", controller.Todo.Store)
		todoGroup.GET("/", controller.Todo.All)
		todoGroup.GET("/:id", controller.Todo.Show)
		todoGroup.PUT("/:id", controller.Todo.Update)
	}

	usersGroup := v1.Group("/users")
	{
		usersGroup.POST("/", controller.User.Register)
	}

	return router
}
