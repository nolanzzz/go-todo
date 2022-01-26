package router

import (
	"github.com/gin-gonic/gin"
	controller2 "todo/controller"
)

type Router struct{}

func InitApiRouter() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/api/v1")
	todoGroup := v1.Group("/todo")
	{
		todoGroup.POST("/", controller2.Todo.Store)
		todoGroup.GET("/", controller2.Todo.All)
		todoGroup.GET("/:id", controller2.Todo.Show)
		todoGroup.PUT("/:id", controller2.Todo.Update)
	}

	usersGroup := v1.Group("/users")
	{
		usersGroup.POST("/", controller2.User.Register)
	}

	return router
}
