package router

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	api "todo/api/v1"
	_ "todo/docs"
	"todo/middleware"
)

func InitApiRouter() *gin.Engine {
	router := gin.Default()
	// Register Swagger handler
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	TodoGroup := router.Group("/api/v1/todo")
	{
		TodoGroup.GET("/:id", api.Todo.Get)
		TodoGroup.GET("/", api.Todo.GetAll)
		TodoGroup.GET("/by/:userID", api.Todo.GetUserAll)

		TodoAuthGroup := TodoGroup.Use(middleware.Auth()) // Only authorized users can make changes
		{
			TodoAuthGroup.POST("/", api.Todo.Create)
			TodoAuthGroup.PUT("/", api.Todo.Update)
			TodoAuthGroup.PUT("/done/:id", api.Todo.Done)
			TodoAuthGroup.PUT("/undone/:id", api.Todo.Undone)
		}
	}

	UserGroup := router.Group("/api/v1/users")
	{
		UserGroup.POST("/register", api.User.Register) // Register
		UserGroup.POST("/login", api.User.Login)       // Login
	}

	RankingGroup := router.Group("/api/v1/ranking")
	{
		RankingGroup.GET("/todos/:limit", api.Ranking.RankingByTodos)
		RankingGroup.GET("/minutes/:limit", api.Ranking.RankingByMinutes)
	}

	return router
}
