package router

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"todo/controller"
	_ "todo/docs"
	"todo/middleware"
)

type Router struct{}

func InitApiRouter() *gin.Engine {
	router := gin.Default()
	// Register Swagger handler
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := router.Group("/api/v1")
	todo := v1.Group("/todo").Use(middleware.Auth())
	{
		todo.GET("/:id", controller.Todo.Get)
		todo.GET("/", controller.Todo.GetAll)
		todo.GET("/by/:userID", controller.Todo.GetUserAll)
	}
	todoAuth := todo.Use(middleware.Auth()) // Only authorized users can make changes
	{
		todoAuth.POST("/", controller.Todo.Create)
		todoAuth.PUT("/:id", controller.Todo.Update)
	}

	users := v1.Group("/users")
	{
		users.POST("/register", controller.User.Register) // Register
		users.POST("/login", controller.User.Login)       // Login
	}

	return router
}
