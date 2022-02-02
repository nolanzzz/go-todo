package router

import (
	"github.com/gin-gonic/gin"
	"todo/controller"
	"todo/middleware"
)

type Router struct{}

func InitApiRouter() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/api/v1")
	todo := v1.Group("/todo").Use(middleware.Auth())
	{
		todo.GET("/", controller.Todo.All)
		todo.GET("/:id", controller.Todo.Show)

		todoAuth := todo.Use(middleware.Auth())
		todoAuth.POST("/", controller.Todo.Store)
		todoAuth.PUT("/:id", controller.Todo.Update)
		todoAuth.GET("/own", controller.Todo.UserAll)
	}

	users := v1.Group("/users")
	{
		// Register
		users.POST("/register", controller.User.Register)
		// Login
		users.POST("/login", controller.User.Login)
	}

	return router
}
