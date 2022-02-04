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
	TodoGroup := v1.Group("/todo")
	{
		TodoGroup.GET("/:id", controller.Todo.Get)
		TodoGroup.GET("/", controller.Todo.GetAll)
		TodoGroup.GET("/by/:userID", controller.Todo.GetUserAll)
	}
	TodoAuthGroup := TodoGroup.Use(middleware.Auth()) // Only authorized users can make changes
	{
		TodoAuthGroup.POST("/", controller.Todo.Create)
		TodoAuthGroup.PUT("/", controller.Todo.Update)
		TodoAuthGroup.PUT("/done/:id", controller.Todo.Done)
		TodoAuthGroup.PUT("/undone/:id", controller.Todo.Undone)
	}

	UserGroup := v1.Group("/users")
	{
		UserGroup.POST("/register", controller.User.Register) // Register
		UserGroup.POST("/login", controller.User.Login)       // Login
	}

	return router
}
