package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/mdhenriques/api-go/controllers"
	"github.com/mdhenriques/api-go/middlewares"

	// Swagger:
	_ "github.com/mdhenriques/api-go/docs"     // <-- Importa os docs gerados
	swaggerFiles "github.com/swaggo/files"     // <-- Import Swagger Files
	ginSwagger "github.com/swaggo/gin-swagger" // <-- Import Swagger Handler
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.POST("/users", controllers.CreateUser)
	r.POST("/login", controllers.Login)

	auth := r.Group("/")
	auth.Use(middlewares.AuthMiddleware())
	{
		auth.GET("/me", controllers.GetMe)
		auth.POST("/tasks", controllers.CreateTask)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
