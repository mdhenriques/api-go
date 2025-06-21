package routes

import (
	"github.com/gin-gonic/gin"

	// Swagger:
	"github.com/mdhenriques/api-go/controllers"
	_ "github.com/mdhenriques/api-go/docs" // Import ne
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.POST("/users", controllers.CreateUser)
	r.POST("/login", controllers.Login)

	r.GET("swagger/*any", ginSwagger.WrapHandler((swaggerFiles.Handler)))

	return r
}