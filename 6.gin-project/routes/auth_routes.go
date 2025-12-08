package routes

import (
	"gin-project/controllers"
	"gin-project/middlewares"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	router.POST("/register", controllers.Register) // Added register route
	router.POST("/login", controllers.Login)
	protected := router.Group("/").Use(middlewares.AuthMiddleware())
	{
		protected.GET("/profile", controllers.GetProfile)
	}
}
