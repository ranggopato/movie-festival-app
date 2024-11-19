package routes

import (
	"movie-festival-app/controllers"

	"movie-festival-app/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.Engine) {
	auth := router.Group("/auth")
	{
		auth.POST("/register", controllers.RegisterHandler) // Register user
		auth.POST("/login", controllers.LoginHandler)       // Login user
	}

	user := router.Group("/user")
	user.Use(middlewares.AuthMiddleware()) // Memerlukan autentikasi
	{
		user.GET("/", controllers.GetUserHandler) // Get user details
	}
}
